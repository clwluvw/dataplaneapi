// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package haproxy

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/renameio"
	"github.com/haproxytech/client-native/v3/models"
	"github.com/haproxytech/dataplaneapi/log"
)

const (
	logFieldReloadID = "reload_id"
)

type IReloadAgent interface {
	Reload() string
	ReloadWithCallback(func()) string
	Restart() error
	ForceReload() error
	ForceReloadWithCallback(func()) error
	GetReloads() models.Reloads
	GetReload(id string) *models.Reload
}

type reloadCache struct {
	failedReloads map[string]*models.Reload
	lastSuccess   *models.Reload
	next          string
	current       string
	index         int64
	retention     int
	mu            sync.RWMutex
	callbacks     map[string]func()
}

type ReloadAgentParams struct {
	Delay      int
	ReloadCmd  string
	RestartCmd string
	ConfigFile string
	BackupDir  string
	Retention  int
	Ctx        context.Context
}

// ReloadAgent handles all reloads, scheduled or forced
type ReloadAgent struct {
	delay         int
	reloadCmd     string
	restartCmd    string
	configFile    string
	lkgConfigFile string
	done          <-chan struct{}
	cache         reloadCache
}

func NewReloadAgent(params ReloadAgentParams) (*ReloadAgent, error) {
	ra := &ReloadAgent{}

	ra.reloadCmd = params.ReloadCmd
	ra.restartCmd = params.RestartCmd
	ra.configFile = params.ConfigFile

	if params.Ctx == nil {
		params.Ctx = context.Background()
	}
	ra.done = params.Ctx.Done()

	params.Delay *= 1000 // delay is defined in seconds - internally in miliseconds
	d := os.Getenv("CI_DATAPLANE_RELOAD_DELAY_OVERRIDE")
	if d != "" {
		params.Delay, _ = strconv.Atoi(d) // in case of err in conversion 0 is returned
	}
	if params.Delay == 0 {
		params.Delay = 5000
	}
	ra.delay = params.Delay

	ra.setLkgPath(params.ConfigFile, params.BackupDir)

	// create last known good file, assume it is valid when starting
	if err := copyFile(ra.configFile, ra.lkgConfigFile); err != nil {
		return nil, err
	}
	ra.cache.Init(params.Retention)
	go ra.handleReloads()

	return ra, nil
}

func (ra *ReloadAgent) setLkgPath(configFile, path string) {
	if path != "" {
		ra.lkgConfigFile = fmt.Sprintf("%s/%s.lkg", path, filepath.Base(configFile))
		return
	}
	ra.lkgConfigFile = configFile + ".lkg"
}

func (ra *ReloadAgent) handleReload(id string) {
	logFields := map[string]interface{}{logFieldReloadID: id}
	ra.cache.mu.Lock()
	ra.cache.current = id

	defer func() {
		ra.cache.next = ""
		ra.cache.mu.Unlock()
	}()

	response, err := ra.reloadHAProxy(id)
	if err != nil {
		ra.cache.failReload(response)
		log.WithFieldsf(logFields, log.WarnLevel, "Reload failed: %s", err)
	} else {
		ra.cache.succeedReload(response)
		callback, ok := ra.cache.callbacks[id]
		if ok {
			callback()
		}
		log.WithFields(logFields, log.DebugLevel, "Handling reload completed, waiting for new requests")
	}
	delete(ra.cache.callbacks, id)
}

func (ra *ReloadAgent) handleReloads() {
	ticker := time.NewTicker(time.Duration(ra.delay) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			if next := ra.cache.getNext(); next != "" {
				ra.handleReload(next)
			}
		case <-ra.done:
			ticker.Stop()
			return
		}
	}
}

func (ra *ReloadAgent) reloadHAProxy(id string) (string, error) {
	logFields := map[string]interface{}{logFieldReloadID: id}
	// try the reload
	log.WithFields(logFields, log.DebugLevel, "Reload started")
	t := time.Now()
	output, err := execCmd(ra.reloadCmd)
	log.WithFieldsf(logFields, log.DebugLevel, "Reload finished in %s", time.Since(t))
	if err != nil {
		reloadFailedError := err
		// if failed, return to last known good file and restart and return the original file
		log.WithFields(logFields, log.InfoLevel, "Reload failed, restarting with last known good config...")
		if err := copyFile(ra.configFile, ra.configFile+".bck"); err != nil {
			return fmt.Sprintf("Reload failed: %s, failed to backup original config file for restart.", output), err
		}
		defer func() {
			// nolint:errcheck
			copyFile(ra.configFile+".bck", ra.configFile)
			os.Remove(ra.configFile + ".bck")
		}()
		if err := copyFile(ra.lkgConfigFile, ra.configFile); err != nil {
			return fmt.Sprintf("Reload failed: %s, failed to revert to last known good config file", output), err
		}
		if err := ra.restartHAProxy(); err != nil {
			log.WithFieldsf(logFields, log.WarnLevel, "Restart failed, please check the reason and restart manually: %s", err)
			return fmt.Sprintf("Reload failed: %s, failed to restart HAProxy, please check and start manually", output), err
		}
		log.WithFields(logFields, log.DebugLevel, "HAProxy restarted with last known good config")
		return output, reloadFailedError
	}
	log.WithFields(logFields, log.DebugLevel, "Reload successful")
	// if success, replace last known good file
	// nolint:errcheck
	copyFile(ra.configFile, ra.lkgConfigFile)
	return output, nil
}

func (ra *ReloadAgent) restartHAProxy() error {
	_, err := execCmd(ra.restartCmd)
	if err != nil {
		return err
	}
	return nil
}

func execCmd(cmd string) (string, error) {
	strArr := strings.Split(cmd, " ")
	var c *exec.Cmd
	if len(strArr) == 1 {
		//nolint:gosec
		c = exec.Command(strArr[0])
	} else {
		//nolint:gosec
		c = exec.Command(strArr[0], strArr[1:]...)
	}
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run()
	if err != nil {
		return stderr.String(), fmt.Errorf("executing %s failed: %s", cmd, err)
	}
	return stdout.String(), nil
}

// Reload schedules a reload
func (ra *ReloadAgent) Reload() string {
	next := ra.cache.getNext()
	if next == "" {
		next = ra.cache.newReload()
		log.WithFields(map[string]interface{}{logFieldReloadID: next}, log.DebugLevel, "Scheduling a new reload...")
	}

	return next
}

// ForceReload calls reload directly
func (ra *ReloadAgent) ForceReload() error {
	r, err := ra.reloadHAProxy("force")
	if err != nil {
		return NewReloadError(fmt.Sprintf("Reload failed: %v, %v", err, r))
	}
	return nil
}

// Reload schedules a reload, callback is called only if reload is successfull
func (ra *ReloadAgent) ReloadWithCallback(callback func()) string {
	next := ra.cache.getNext()
	if next == "" {
		next = ra.cache.newReloadWithCallback(callback)
		log.WithFields(map[string]interface{}{logFieldReloadID: next}, log.DebugLevel, "Scheduling a new reload...")
	}

	return next
}

// ForceReload calls reload directly, callback is called only if reload is successfull
func (ra *ReloadAgent) ForceReloadWithCallback(callback func()) error {
	r, err := ra.reloadHAProxy("force")
	if err != nil {
		return NewReloadError(fmt.Sprintf("Reload failed: %v, %v", err, r))
	}
	callback()
	return nil
}

func (rc *reloadCache) Init(retention int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.failedReloads = make(map[string]*models.Reload)
	rc.current = ""
	rc.next = ""
	rc.lastSuccess = nil
	rc.index = 0
	rc.retention = retention
	rc.callbacks = make(map[string]func())
}

func (rc *reloadCache) newReload() string {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	id := fmt.Sprintf("%s-%v", time.Now().Format("2006-01-02"), rc.index)
	rc.index++
	rc.next = id
	return rc.next
}

func (rc *reloadCache) newReloadWithCallback(callback func()) string {
	next := rc.newReload()
	rc.callbacks[next] = callback
	return next
}

func (rc *reloadCache) getNext() string {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.next
}

func (rc *reloadCache) failReload(response string) {
	r := &models.Reload{
		ID:              rc.current,
		Status:          models.ReloadStatusFailed,
		Response:        response,
		ReloadTimestamp: time.Now().Unix(),
	}

	rc.failedReloads[rc.current] = r
	rc.current = ""
	rc.clearReloads()
}

func (rc *reloadCache) succeedReload(response string) {
	r := &models.Reload{
		ID:              rc.current,
		Status:          models.ReloadStatusSucceeded,
		Response:        response,
		ReloadTimestamp: time.Now().Unix(),
	}

	rc.lastSuccess = r
	rc.current = ""
}

func (rc *reloadCache) clearReloads() {
	now := time.Now().Unix()

	for k, v := range rc.failedReloads {
		if (now - v.ReloadTimestamp) > int64((rc.retention * 86400)) {
			delete(rc.failedReloads, k)
		}
	}
}

func (ra *ReloadAgent) GetReloads() models.Reloads {
	ra.cache.mu.RLock()
	defer ra.cache.mu.RUnlock()

	v := make([]*models.Reload, 0, len(ra.cache.failedReloads))
	for _, value := range ra.cache.failedReloads {
		v = append(v, value)
	}

	if ra.cache.lastSuccess != nil {
		v = append(v, ra.cache.lastSuccess)
	}

	if ra.cache.current != "" {
		r := &models.Reload{
			ID:     ra.cache.current,
			Status: models.ReloadStatusInProgress,
		}
		v = append(v, r)
	}

	if ra.cache.next != "" {
		r := &models.Reload{
			ID:     ra.cache.next,
			Status: models.ReloadStatusInProgress,
		}
		v = append(v, r)
	}
	return v
}

func (ra *ReloadAgent) GetReload(id string) *models.Reload {
	ra.cache.mu.RLock()
	defer ra.cache.mu.RUnlock()

	if ra.cache.current == id {
		return &models.Reload{
			ID:     ra.cache.current,
			Status: models.ReloadStatusInProgress,
		}
	}
	if ra.cache.next == id {
		return &models.Reload{
			ID:     ra.cache.current,
			Status: models.ReloadStatusInProgress,
		}
	}

	v, ok := ra.cache.failedReloads[id]
	if ok {
		return v
	}
	if ra.cache.lastSuccess != nil {
		if ra.cache.lastSuccess.ID == id {
			return ra.cache.lastSuccess
		}

		// if it is older than last success return success
		sDate, sIndex, err := getTimeIndexFromID(ra.cache.lastSuccess.ID)
		if err != nil {
			return nil
		}
		gDate, gIndex, err := getTimeIndexFromID(id)
		if err != nil {
			return nil
		}

		if gDate.Before(sDate) {
			return &models.Reload{
				ID:     id,
				Status: models.ReloadStatusSucceeded,
			}
		}

		if sIndex > gIndex {
			return &models.Reload{
				ID:     id,
				Status: models.ReloadStatusSucceeded,
			}
		}
	}
	return nil
}

func (ra *ReloadAgent) Restart() error {
	return ra.restartHAProxy()
}

func getTimeIndexFromID(id string) (time.Time, int64, error) {
	data := strings.Split(id, "-")
	index, err := strconv.ParseInt(data[len(data)-1], 10, 64)
	if err != nil {
		return time.Now(), 0, err
	}
	date, err := time.Parse("2006-01-02", strings.Join(data[:len(data)-1], "-"))
	if err != nil {
		return date, 0, err
	}

	return date, index, nil
}

// ReloadError general configuration client error
type ReloadError struct {
	msg string
}

// Error implementation for ConfError
func (e *ReloadError) Error() string {
	return fmt.Sprintf(e.msg)
}

// NewReloadError constructor for ReloadError
func NewReloadError(msg string) *ReloadError {
	return &ReloadError{msg: msg}
}

func copyFile(src, dest string) error {
	srcContent, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcContent.Close()

	data, err := ioutil.ReadAll(srcContent)
	if err != nil {
		return err
	}
	return renameio.WriteFile(dest, data, 0o644)
}
