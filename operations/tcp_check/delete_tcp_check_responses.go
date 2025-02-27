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

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteTCPCheckAcceptedCode is the HTTP code returned for type DeleteTCPCheckAccepted
const DeleteTCPCheckAcceptedCode int = 202

/*DeleteTCPCheckAccepted Configuration change accepted and reload requested

swagger:response deleteTcpCheckAccepted
*/
type DeleteTCPCheckAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteTCPCheckAccepted creates DeleteTCPCheckAccepted with default headers values
func NewDeleteTCPCheckAccepted() *DeleteTCPCheckAccepted {

	return &DeleteTCPCheckAccepted{}
}

// WithReloadID adds the reloadId to the delete Tcp check accepted response
func (o *DeleteTCPCheckAccepted) WithReloadID(reloadID string) *DeleteTCPCheckAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Tcp check accepted response
func (o *DeleteTCPCheckAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteTCPCheckAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteTCPCheckNoContentCode is the HTTP code returned for type DeleteTCPCheckNoContent
const DeleteTCPCheckNoContentCode int = 204

/*DeleteTCPCheckNoContent TCP check deleted

swagger:response deleteTcpCheckNoContent
*/
type DeleteTCPCheckNoContent struct {
}

// NewDeleteTCPCheckNoContent creates DeleteTCPCheckNoContent with default headers values
func NewDeleteTCPCheckNoContent() *DeleteTCPCheckNoContent {

	return &DeleteTCPCheckNoContent{}
}

// WriteResponse to the client
func (o *DeleteTCPCheckNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTCPCheckNotFoundCode is the HTTP code returned for type DeleteTCPCheckNotFound
const DeleteTCPCheckNotFoundCode int = 404

/*DeleteTCPCheckNotFound The specified resource was not found

swagger:response deleteTcpCheckNotFound
*/
type DeleteTCPCheckNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTCPCheckNotFound creates DeleteTCPCheckNotFound with default headers values
func NewDeleteTCPCheckNotFound() *DeleteTCPCheckNotFound {

	return &DeleteTCPCheckNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Tcp check not found response
func (o *DeleteTCPCheckNotFound) WithConfigurationVersion(configurationVersion string) *DeleteTCPCheckNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Tcp check not found response
func (o *DeleteTCPCheckNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Tcp check not found response
func (o *DeleteTCPCheckNotFound) WithPayload(payload *models.Error) *DeleteTCPCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Tcp check not found response
func (o *DeleteTCPCheckNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTCPCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteTCPCheckDefault General Error

swagger:response deleteTcpCheckDefault
*/
type DeleteTCPCheckDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTCPCheckDefault creates DeleteTCPCheckDefault with default headers values
func NewDeleteTCPCheckDefault(code int) *DeleteTCPCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTCPCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete TCP check default response
func (o *DeleteTCPCheckDefault) WithStatusCode(code int) *DeleteTCPCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete TCP check default response
func (o *DeleteTCPCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete TCP check default response
func (o *DeleteTCPCheckDefault) WithConfigurationVersion(configurationVersion string) *DeleteTCPCheckDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete TCP check default response
func (o *DeleteTCPCheckDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete TCP check default response
func (o *DeleteTCPCheckDefault) WithPayload(payload *models.Error) *DeleteTCPCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete TCP check default response
func (o *DeleteTCPCheckDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTCPCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
