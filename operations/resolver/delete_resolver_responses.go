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

package resolver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteResolverAcceptedCode is the HTTP code returned for type DeleteResolverAccepted
const DeleteResolverAcceptedCode int = 202

/*DeleteResolverAccepted Configuration change accepted and reload requested

swagger:response deleteResolverAccepted
*/
type DeleteResolverAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteResolverAccepted creates DeleteResolverAccepted with default headers values
func NewDeleteResolverAccepted() *DeleteResolverAccepted {

	return &DeleteResolverAccepted{}
}

// WithReloadID adds the reloadId to the delete resolver accepted response
func (o *DeleteResolverAccepted) WithReloadID(reloadID string) *DeleteResolverAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete resolver accepted response
func (o *DeleteResolverAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteResolverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteResolverNoContentCode is the HTTP code returned for type DeleteResolverNoContent
const DeleteResolverNoContentCode int = 204

/*DeleteResolverNoContent Resolver deleted

swagger:response deleteResolverNoContent
*/
type DeleteResolverNoContent struct {
}

// NewDeleteResolverNoContent creates DeleteResolverNoContent with default headers values
func NewDeleteResolverNoContent() *DeleteResolverNoContent {

	return &DeleteResolverNoContent{}
}

// WriteResponse to the client
func (o *DeleteResolverNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteResolverNotFoundCode is the HTTP code returned for type DeleteResolverNotFound
const DeleteResolverNotFoundCode int = 404

/*DeleteResolverNotFound The specified resource was not found

swagger:response deleteResolverNotFound
*/
type DeleteResolverNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteResolverNotFound creates DeleteResolverNotFound with default headers values
func NewDeleteResolverNotFound() *DeleteResolverNotFound {

	return &DeleteResolverNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete resolver not found response
func (o *DeleteResolverNotFound) WithConfigurationVersion(configurationVersion string) *DeleteResolverNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete resolver not found response
func (o *DeleteResolverNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete resolver not found response
func (o *DeleteResolverNotFound) WithPayload(payload *models.Error) *DeleteResolverNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete resolver not found response
func (o *DeleteResolverNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteResolverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteResolverDefault General Error

swagger:response deleteResolverDefault
*/
type DeleteResolverDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteResolverDefault creates DeleteResolverDefault with default headers values
func NewDeleteResolverDefault(code int) *DeleteResolverDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteResolverDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete resolver default response
func (o *DeleteResolverDefault) WithStatusCode(code int) *DeleteResolverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete resolver default response
func (o *DeleteResolverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete resolver default response
func (o *DeleteResolverDefault) WithConfigurationVersion(configurationVersion string) *DeleteResolverDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete resolver default response
func (o *DeleteResolverDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete resolver default response
func (o *DeleteResolverDefault) WithPayload(payload *models.Error) *DeleteResolverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete resolver default response
func (o *DeleteResolverDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteResolverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
