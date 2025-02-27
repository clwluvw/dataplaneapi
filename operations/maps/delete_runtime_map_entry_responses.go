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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteRuntimeMapEntryNoContentCode is the HTTP code returned for type DeleteRuntimeMapEntryNoContent
const DeleteRuntimeMapEntryNoContentCode int = 204

/*DeleteRuntimeMapEntryNoContent Map key/value deleted

swagger:response deleteRuntimeMapEntryNoContent
*/
type DeleteRuntimeMapEntryNoContent struct {
}

// NewDeleteRuntimeMapEntryNoContent creates DeleteRuntimeMapEntryNoContent with default headers values
func NewDeleteRuntimeMapEntryNoContent() *DeleteRuntimeMapEntryNoContent {

	return &DeleteRuntimeMapEntryNoContent{}
}

// WriteResponse to the client
func (o *DeleteRuntimeMapEntryNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteRuntimeMapEntryNotFoundCode is the HTTP code returned for type DeleteRuntimeMapEntryNotFound
const DeleteRuntimeMapEntryNotFoundCode int = 404

/*DeleteRuntimeMapEntryNotFound The specified resource was not found

swagger:response deleteRuntimeMapEntryNotFound
*/
type DeleteRuntimeMapEntryNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteRuntimeMapEntryNotFound creates DeleteRuntimeMapEntryNotFound with default headers values
func NewDeleteRuntimeMapEntryNotFound() *DeleteRuntimeMapEntryNotFound {

	return &DeleteRuntimeMapEntryNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete runtime map entry not found response
func (o *DeleteRuntimeMapEntryNotFound) WithConfigurationVersion(configurationVersion string) *DeleteRuntimeMapEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete runtime map entry not found response
func (o *DeleteRuntimeMapEntryNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete runtime map entry not found response
func (o *DeleteRuntimeMapEntryNotFound) WithPayload(payload *models.Error) *DeleteRuntimeMapEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete runtime map entry not found response
func (o *DeleteRuntimeMapEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRuntimeMapEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteRuntimeMapEntryDefault General Error

swagger:response deleteRuntimeMapEntryDefault
*/
type DeleteRuntimeMapEntryDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteRuntimeMapEntryDefault creates DeleteRuntimeMapEntryDefault with default headers values
func NewDeleteRuntimeMapEntryDefault(code int) *DeleteRuntimeMapEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteRuntimeMapEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete runtime map entry default response
func (o *DeleteRuntimeMapEntryDefault) WithStatusCode(code int) *DeleteRuntimeMapEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete runtime map entry default response
func (o *DeleteRuntimeMapEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete runtime map entry default response
func (o *DeleteRuntimeMapEntryDefault) WithConfigurationVersion(configurationVersion string) *DeleteRuntimeMapEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete runtime map entry default response
func (o *DeleteRuntimeMapEntryDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete runtime map entry default response
func (o *DeleteRuntimeMapEntryDefault) WithPayload(payload *models.Error) *DeleteRuntimeMapEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete runtime map entry default response
func (o *DeleteRuntimeMapEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRuntimeMapEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
