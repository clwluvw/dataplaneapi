// Code generated by go-swagger; DO NOT EDIT.

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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteNameserverAcceptedCode is the HTTP code returned for type DeleteNameserverAccepted
const DeleteNameserverAcceptedCode int = 202

/*DeleteNameserverAccepted Configuration change accepted and reload requested

swagger:response deleteNameserverAccepted
*/
type DeleteNameserverAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteNameserverAccepted creates DeleteNameserverAccepted with default headers values
func NewDeleteNameserverAccepted() *DeleteNameserverAccepted {

	return &DeleteNameserverAccepted{}
}

// WithReloadID adds the reloadId to the delete nameserver accepted response
func (o *DeleteNameserverAccepted) WithReloadID(reloadID string) *DeleteNameserverAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete nameserver accepted response
func (o *DeleteNameserverAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteNameserverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteNameserverNoContentCode is the HTTP code returned for type DeleteNameserverNoContent
const DeleteNameserverNoContentCode int = 204

/*DeleteNameserverNoContent Nameserver deleted

swagger:response deleteNameserverNoContent
*/
type DeleteNameserverNoContent struct {
}

// NewDeleteNameserverNoContent creates DeleteNameserverNoContent with default headers values
func NewDeleteNameserverNoContent() *DeleteNameserverNoContent {

	return &DeleteNameserverNoContent{}
}

// WriteResponse to the client
func (o *DeleteNameserverNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteNameserverNotFoundCode is the HTTP code returned for type DeleteNameserverNotFound
const DeleteNameserverNotFoundCode int = 404

/*DeleteNameserverNotFound The specified resource was not found

swagger:response deleteNameserverNotFound
*/
type DeleteNameserverNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNameserverNotFound creates DeleteNameserverNotFound with default headers values
func NewDeleteNameserverNotFound() *DeleteNameserverNotFound {

	return &DeleteNameserverNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete nameserver not found response
func (o *DeleteNameserverNotFound) WithConfigurationVersion(configurationVersion string) *DeleteNameserverNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete nameserver not found response
func (o *DeleteNameserverNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete nameserver not found response
func (o *DeleteNameserverNotFound) WithPayload(payload *models.Error) *DeleteNameserverNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete nameserver not found response
func (o *DeleteNameserverNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNameserverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteNameserverDefault General Error

swagger:response deleteNameserverDefault
*/
type DeleteNameserverDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNameserverDefault creates DeleteNameserverDefault with default headers values
func NewDeleteNameserverDefault(code int) *DeleteNameserverDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteNameserverDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete nameserver default response
func (o *DeleteNameserverDefault) WithStatusCode(code int) *DeleteNameserverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete nameserver default response
func (o *DeleteNameserverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete nameserver default response
func (o *DeleteNameserverDefault) WithConfigurationVersion(configurationVersion string) *DeleteNameserverDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete nameserver default response
func (o *DeleteNameserverDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete nameserver default response
func (o *DeleteNameserverDefault) WithPayload(payload *models.Error) *DeleteNameserverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete nameserver default response
func (o *DeleteNameserverDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNameserverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
