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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteStorageSSLCertificateAcceptedCode is the HTTP code returned for type DeleteStorageSSLCertificateAccepted
const DeleteStorageSSLCertificateAcceptedCode int = 202

/*DeleteStorageSSLCertificateAccepted SSL certificate deleted and reload requested

swagger:response deleteStorageSSLCertificateAccepted
*/
type DeleteStorageSSLCertificateAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteStorageSSLCertificateAccepted creates DeleteStorageSSLCertificateAccepted with default headers values
func NewDeleteStorageSSLCertificateAccepted() *DeleteStorageSSLCertificateAccepted {

	return &DeleteStorageSSLCertificateAccepted{}
}

// WithReloadID adds the reloadId to the delete storage s s l certificate accepted response
func (o *DeleteStorageSSLCertificateAccepted) WithReloadID(reloadID string) *DeleteStorageSSLCertificateAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete storage s s l certificate accepted response
func (o *DeleteStorageSSLCertificateAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteStorageSSLCertificateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteStorageSSLCertificateNoContentCode is the HTTP code returned for type DeleteStorageSSLCertificateNoContent
const DeleteStorageSSLCertificateNoContentCode int = 204

/*DeleteStorageSSLCertificateNoContent SSL certificate deleted

swagger:response deleteStorageSSLCertificateNoContent
*/
type DeleteStorageSSLCertificateNoContent struct {
}

// NewDeleteStorageSSLCertificateNoContent creates DeleteStorageSSLCertificateNoContent with default headers values
func NewDeleteStorageSSLCertificateNoContent() *DeleteStorageSSLCertificateNoContent {

	return &DeleteStorageSSLCertificateNoContent{}
}

// WriteResponse to the client
func (o *DeleteStorageSSLCertificateNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteStorageSSLCertificateNotFoundCode is the HTTP code returned for type DeleteStorageSSLCertificateNotFound
const DeleteStorageSSLCertificateNotFoundCode int = 404

/*DeleteStorageSSLCertificateNotFound The specified resource was not found

swagger:response deleteStorageSSLCertificateNotFound
*/
type DeleteStorageSSLCertificateNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStorageSSLCertificateNotFound creates DeleteStorageSSLCertificateNotFound with default headers values
func NewDeleteStorageSSLCertificateNotFound() *DeleteStorageSSLCertificateNotFound {

	return &DeleteStorageSSLCertificateNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete storage s s l certificate not found response
func (o *DeleteStorageSSLCertificateNotFound) WithConfigurationVersion(configurationVersion string) *DeleteStorageSSLCertificateNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete storage s s l certificate not found response
func (o *DeleteStorageSSLCertificateNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete storage s s l certificate not found response
func (o *DeleteStorageSSLCertificateNotFound) WithPayload(payload *models.Error) *DeleteStorageSSLCertificateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage s s l certificate not found response
func (o *DeleteStorageSSLCertificateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageSSLCertificateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteStorageSSLCertificateDefault General Error

swagger:response deleteStorageSSLCertificateDefault
*/
type DeleteStorageSSLCertificateDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStorageSSLCertificateDefault creates DeleteStorageSSLCertificateDefault with default headers values
func NewDeleteStorageSSLCertificateDefault(code int) *DeleteStorageSSLCertificateDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteStorageSSLCertificateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) WithStatusCode(code int) *DeleteStorageSSLCertificateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) WithConfigurationVersion(configurationVersion string) *DeleteStorageSSLCertificateDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) WithPayload(payload *models.Error) *DeleteStorageSSLCertificateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageSSLCertificateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
