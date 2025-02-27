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

package configuration

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/haproxytech/client-native/v3/models"
	"github.com/haproxytech/dataplaneapi/log"
	jsoniter "github.com/json-iterator/go"
)

var cfg *Configuration

type HAProxyConfiguration struct {
	ConfigFile           string `short:"c" long:"config-file" description:"Path to the haproxy configuration file" default:"/etc/haproxy/haproxy.cfg" group:"haproxy"`
	Userlist             string `short:"u" long:"userlist" description:"Userlist in HAProxy configuration to use for API Basic Authentication" default:"controller" group:"userlist"`
	HAProxy              string `short:"b" long:"haproxy-bin" description:"Path to the haproxy binary file" default:"haproxy" group:"haproxy"`
	ReloadDelay          int    `short:"d" long:"reload-delay" description:"Minimum delay between two reloads (in s)" default:"5" group:"reload"`
	ReloadCmd            string `short:"r" long:"reload-cmd" description:"Reload command" group:"reload"`
	RestartCmd           string `short:"s" long:"restart-cmd" description:"Restart command" group:"reload"`
	ReloadRetention      int    `long:"reload-retention" description:"Reload retention in days, every older reload id will be deleted" default:"1" group:"reload"`
	TransactionDir       string `short:"t" long:"transaction-dir" description:"Path to the transaction directory" default:"/tmp/haproxy" group:"transaction"`
	BackupsNumber        int    `short:"n" long:"backups-number" description:"Number of backup configuration files you want to keep, stored in the config dir with version number suffix" default:"0" group:"transaction"`
	BackupsDir           string `long:"backups-dir" description:"Path to directory in which to place backup files" group:"transaction"`
	MasterRuntime        string `short:"m" long:"master-runtime" description:"Path to the master Runtime API socket" group:"haproxy"`
	ShowSystemInfo       bool   `short:"i" long:"show-system-info" description:"Show system info on info endpoint" group:"dataplaneapi"`
	DataplaneConfig      string `short:"f" description:"Path to the dataplane configuration file" default:"/etc/haproxy/dataplaneapi.hcl" yaml:"-"`
	UserListFile         string `long:"userlist-file" description:"Path to the dataplaneapi userlist file. By default userlist is read from HAProxy conf. When specified userlist would be read from this file" group:"userlist"`
	NodeIDFile           string `long:"fid" description:"Path to file that will dataplaneapi use to write its id (not a pid) that was given to him after joining a cluster" group:"haproxy"`
	MapsDir              string `short:"p" long:"maps-dir" description:"Path to directory of map files managed by dataplane" default:"/etc/haproxy/maps" group:"resources"`
	SSLCertsDir          string `long:"ssl-certs-dir" description:"Path to SSL certificates directory" default:"/etc/haproxy/ssl" group:"resources"`
	GeneralStorageDir    string `long:"general-storage-dir" description:"Path to general storage directory" default:"/etc/haproxy/general" group:"resources"`
	UpdateMapFiles       bool   `long:"update-map-files" description:"Flag used for syncing map files with runtime maps values" group:"resources"`
	UpdateMapFilesPeriod int64  `long:"update-map-files-period" description:"Elapsed time in seconds between two maps syncing operations" default:"10" group:"resources"`
	ClusterTLSCertDir    string `long:"cluster-tls-dir" description:"Path where cluster tls certificates will be stored. Defaults to same directory as dataplane configuration file" group:"cluster"`
	SpoeDir              string `long:"spoe-dir" description:"Path to SPOE directory." default:"/etc/haproxy/spoe" group:"resources"`
	SpoeTransactionDir   string `long:"spoe-transaction-dir" description:"Path to the SPOE transaction directory" default:"/tmp/spoe-haproxy" group:"resources"`
	MasterWorkerMode     bool   `long:"master-worker-mode" description:"Flag to enable helpers when running within HAProxy" group:"haproxy"`
	MaxOpenTransactions  int64  `long:"max-open-transactions" description:"Limit for active transaction in pending state" default:"20" group:"transaction"`
	ValidateCmd          string `long:"validate-cmd" description:"Executes a custom command to perform the HAProxy configuration check" group:"reload"`
	DisableInotify       bool   `long:"disable-inotify" description:"Disables inotify watcher watcher for the configuration file" group:"dataplaneapi"`
	PIDFile              string `long:"pid-file" description:"Path to file that will dataplaneapi use to write its pid" group:"dataplaneapi" example:"/tmp/dataplane.pid"`
	UID                  int    `long:"uid" description:"User id value to set on start" group:"dataplaneapi" example:"1000"`
	GID                  int    `long:"gid" description:"Group id value to set on start" group:"dataplaneapi" example:"1000"`
}

type User struct {
	Name     string `long:"name" description:"User name" group:"user" hcl:"name,key" example:"admin"`
	Insecure bool   `long:"insecure" description:"insecure password" group:"user" example:"true"`
	Password string `long:"password" description:"password" group:"user" example:"adminpwd"`
}

type APIConfiguration struct {
	APIAddress string `long:"api-address" description:"Advertised API address" group:"advertised" hcl:"address" example:"10.2.3.4" save:"true"`
	APIPort    int64  `long:"api-port" description:"Advertised API port" group:"advertised" hcl:"port" example:"80" save:"true"`
}

type ClusterConfiguration struct {
	ID                 AtomicString               `yaml:"id,omitempty" group:"cluster" save:"true"`
	BootstrapKey       AtomicString               `yaml:"bootstrap_key,omitempty" group:"cluster" save:"true"`
	ActiveBootstrapKey AtomicString               `yaml:"active_bootstrap_key,omitempty" group:"cluster" save:"true"`
	Token              AtomicString               `yaml:"token,omitempty" group:"cluster" save:"true"`
	URL                AtomicString               `yaml:"url,omitempty" group:"cluster" save:"true"`
	Port               AtomicInt                  `yaml:"port,omitempty" group:"cluster" save:"true"`
	APIBasePath        AtomicString               `yaml:"api_base_path,omitempty" group:"cluster" save:"true"`
	APINodesPath       AtomicString               `yaml:"api_nodes_path,omitempty" group:"cluster" save:"true"`
	APIRegisterPath    AtomicString               `yaml:"api_register_path,omitempty" group:"cluster" save:"true"`
	StorageDir         AtomicString               `yaml:"storage_dir,omitempty" group:"cluster" save:"true"`
	CertificateDir     AtomicString               `yaml:"cert_path,omitempty" group:"cluster" save:"true"`
	CertificateFetched AtomicBool                 `yaml:"cert_fetched,omitempty" group:"cluster" save:"true" example:"false"`
	Name               AtomicString               `yaml:"name,omitempty" group:"cluster" save:"true"`
	Description        AtomicString               `yaml:"description,omitempty" group:"cluster" save:"true"`
	ClusterLogTargets  []*models.ClusterLogTarget `yaml:"cluster_log_targets,omitempty" group:"cluster" save:"true"`
	ClusterID          AtomicString               `yaml:"cluster_id,omitempty" group:"cluster" save:"true"`
}

func (c *ClusterConfiguration) Clear() {
	c.ID.Store("")
	c.ActiveBootstrapKey.Store("")
	c.Token.Store("")
	c.Port.Store(0)
	c.APIBasePath.Store("")
	c.APINodesPath.Store("")
	c.APIRegisterPath.Store("")
	c.CertificateFetched.Store(false)
	c.Name.Store("")
	c.Description.Store("")
	c.ClusterID.Store("")
	c.ClusterLogTargets = nil
}

type RuntimeData struct {
	Host        string
	Port        int
	APIBasePath string
}

type NotifyConfiguration struct {
	BootstrapKeyChanged *ChanNotify `yaml:"-"`
	ServerStarted       *ChanNotify `yaml:"-"`
	CertificateRefresh  *ChanNotify `yaml:"-"`
	Reload              *ChanNotify `yaml:"-"`
	Shutdown            *ChanNotify `yaml:"-"`
}

type ServiceDiscovery struct {
	consulMu   sync.Mutex
	Consuls    []*models.Consul `yaml:"consuls" group:"service_discovery" save:"true"`
	awsMu      sync.Mutex
	AWSRegions []*models.AwsRegion `yaml:"aws-regions" group:"service_discovery" save:"true"`
}

//nolint:staticcheck
type Configuration struct {
	Name                   AtomicString         `yaml:"name" example:"famous_condor"`
	storage                Storage              `yaml:"-"`
	HAProxy                HAProxyConfiguration `yaml:"-"`
	Logging                log.LoggingOptions   `yaml:"-"`
	LogTargets             log.Targets          `yaml:"log_targets,omitempty" hcl:"log_targets" group:"log"`
	Syslog                 log.SyslogOptions    `yaml:"-"`
	APIOptions             APIConfiguration     `yaml:"-"`
	Cluster                ClusterConfiguration `yaml:"-"`
	RuntimeData            RuntimeData          `yaml:"-"`
	Notify                 NotifyConfiguration  `yaml:"-"`
	ServiceDiscovery       ServiceDiscovery     `yaml:"-"`
	Users                  []User               `yaml:"-"`
	Mode                   AtomicString         `yaml:"mode" default:"single"`
	DeprecatedBootstrapKey AtomicString         `yaml:"bootstrap_key,omitempty" deprecated:"true"` // deprecated - use Cluster.BootstrapKey
	Status                 AtomicString         `yaml:"status,omitempty"`
	Cmdline                AtomicString         `yaml:"-"`
	MapSync                *MapSync             `yaml:"-"`
	mutex                  sync.Mutex
	shutdownSignal         chan os.Signal
	reloadSignal           chan os.Signal
}

var (
	cfgInitOnce sync.Once
)

// Get returns pointer to configuration
func Get() *Configuration {
	cfgInitOnce.Do(func() {
		cfg = &Configuration{}
		cfg.Notify.BootstrapKeyChanged = NewChanNotify()
		cfg.Notify.ServerStarted = NewChanNotify()
		cfg.Notify.CertificateRefresh = NewChanNotify()
		cfg.Notify.Reload = NewChanNotify()
		cfg.Notify.Shutdown = NewChanNotify()
		cfg.MapSync = NewMapSync()

		var sb strings.Builder
		for _, v := range os.Args {
			if !strings.HasPrefix(v, "-") && !strings.Contains(v, `\ `) && strings.ContainsAny(v, " ") {
				fmt.Fprintf(&sb, "\"%s\" ", v)
			} else {
				fmt.Fprintf(&sb, "%s ", v)
			}
		}

		cfg.Cmdline.Store(sb.String())
	})
	return cfg
}

func (c *Configuration) GetStorageData() *StorageDataplaneAPIConfiguration {
	return c.storage.Get()
}

func (c *Configuration) UnSubscribeAll() {
	c.Notify.BootstrapKeyChanged.UnSubscribeAll()
	c.Notify.ServerStarted.UnSubscribeAll()
	c.Notify.CertificateRefresh.UnSubscribeAll()
	c.Notify.Reload.UnSubscribeAll()
	c.Notify.Shutdown.UnSubscribeAll()
}

func (c *Configuration) Load() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var err error
	if c.HAProxy.DataplaneConfig == "" {
		c.storage = &StorageDummy{}
		_ = c.storage.Load("")
	} else {
		ext := strings.ToLower(filepath.Ext(c.HAProxy.DataplaneConfig))

		switch ext {
		case ".yml", ".yaml":
			c.storage = &StorageYML{}
		case ".hcl":
			c.storage = &StorageHCL{}
		default:
			if err = (&StorageYML{}).Load(c.HAProxy.DataplaneConfig); err == nil {
				c.storage = &StorageYML{}

				break
			}
			c.storage = &StorageHCL{}
		}
		if err = c.storage.Load(c.HAProxy.DataplaneConfig); err != nil {
			if os.IsNotExist(err) {
				log.Warningf("configuration file %s does not exists, creating one", c.HAProxy.DataplaneConfig)
			} else {
				return fmt.Errorf("configuration file %s not valid: %w", c.HAProxy.DataplaneConfig, err)
			}
		}
	}
	copyToConfiguration(c)

	if c.DeprecatedBootstrapKey.Load() != "" {
		c.Cluster.BootstrapKey.Store(c.DeprecatedBootstrapKey.Load())
	}

	if c.Mode.Load() == "" {
		c.Mode.Store(ModeSingle)
	}

	if c.Name.Load() == "" {
		rand.Seed(time.Now().UnixNano())
		c.Name.Store(petname.Generate(2, "_"))
	}

	return nil
}

func (c *Configuration) LoadRuntimeVars(swaggerJSON json.RawMessage, host string, port int) error {
	var m map[string]interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(swaggerJSON, &m)
	if err != nil {
		return err
	}

	cfg.RuntimeData.APIBasePath = m["basePath"].(string)
	if host == "localhost" {
		host = "127.0.0.1"
	}
	cfg.RuntimeData.Host = host
	cfg.RuntimeData.Port = port

	return nil
}

func (c *Configuration) Save() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	copyConfigurationToStorage(c)
	if len(c.ServiceDiscovery.Consuls) == 0 {
		cfg := c.storage.Get()
		cfg.ServiceDiscovery.Consuls = nil
	}
	if len(c.ServiceDiscovery.AWSRegions) == 0 {
		cfg := c.storage.Get()
		cfg.ServiceDiscovery.AWSRegions = nil
	}
	if cfg.ServiceDiscovery.Consuls == nil && cfg.ServiceDiscovery.AWSRegions == nil {
		cfg := c.storage.Get()
		cfg.ServiceDiscovery = nil
	}
	if len(c.LogTargets) == 0 {
		cfg := c.storage.Get()
		cfg.LogTargets = nil
	}
	if len(c.Cluster.ClusterLogTargets) == 0 {
		cfg := c.storage.Get()
		cfg.Cluster.ClusterLogTargets = nil
	}
	// clean storage data if we are not in cluster mode or preparing to go into that mode
	if cfg.Mode.Load() != ModeCluster && cfg.Cluster.BootstrapKey.Load() == "" {
		storage := cfg.storage.Get()
		storage.Cluster = nil
	}
	return c.storage.Save()
}

func (c *Configuration) GetClusterCertDir() string {
	dir := c.Cluster.CertificateDir.Load()
	if dir == "" {
		dir = c.HAProxy.ClusterTLSCertDir
	}
	if dir == "" {
		// use same dir as dataplane config file
		url := c.HAProxy.DataplaneConfig
		dir = filepath.Dir(url)
	}
	return dir
}

func (c *Configuration) SaveConsuls(consuls []*models.Consul) error {
	c.ServiceDiscovery.consulMu.Lock()
	c.ServiceDiscovery.Consuls = consuls
	c.ServiceDiscovery.consulMu.Unlock()
	return c.Save()
}

func (c *Configuration) SaveAWS(aws []*models.AwsRegion) error {
	c.ServiceDiscovery.awsMu.Lock()
	defer c.ServiceDiscovery.awsMu.Unlock()

	c.ServiceDiscovery.AWSRegions = aws
	return c.Save()
}
