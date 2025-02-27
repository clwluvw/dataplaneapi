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

package bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateBindCreatedCode is the HTTP code returned for type CreateBindCreated
const CreateBindCreatedCode int = 201

/*CreateBindCreated Bind created

swagger:response createBindCreated
*/
type CreateBindCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Bind `json:"body,omitempty"`
}

// NewCreateBindCreated creates CreateBindCreated with default headers values
func NewCreateBindCreated() *CreateBindCreated {

	return &CreateBindCreated{}
}

// WithPayload adds the payload to the create bind created response
func (o *CreateBindCreated) WithPayload(payload *models.Bind) *CreateBindCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bind created response
func (o *CreateBindCreated) SetPayload(payload *models.Bind) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBindCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBindAcceptedCode is the HTTP code returned for type CreateBindAccepted
const CreateBindAcceptedCode int = 202

/*CreateBindAccepted Configuration change accepted and reload requested

swagger:response createBindAccepted
*/
type CreateBindAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Bind `json:"body,omitempty"`
}

// NewCreateBindAccepted creates CreateBindAccepted with default headers values
func NewCreateBindAccepted() *CreateBindAccepted {

	return &CreateBindAccepted{}
}

// WithReloadID adds the reloadId to the create bind accepted response
func (o *CreateBindAccepted) WithReloadID(reloadID string) *CreateBindAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create bind accepted response
func (o *CreateBindAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create bind accepted response
func (o *CreateBindAccepted) WithPayload(payload *models.Bind) *CreateBindAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bind accepted response
func (o *CreateBindAccepted) SetPayload(payload *models.Bind) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBindAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBindBadRequestCode is the HTTP code returned for type CreateBindBadRequest
const CreateBindBadRequestCode int = 400

/*CreateBindBadRequest Bad request

swagger:response createBindBadRequest
*/
type CreateBindBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBindBadRequest creates CreateBindBadRequest with default headers values
func NewCreateBindBadRequest() *CreateBindBadRequest {

	return &CreateBindBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create bind bad request response
func (o *CreateBindBadRequest) WithConfigurationVersion(configurationVersion string) *CreateBindBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create bind bad request response
func (o *CreateBindBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create bind bad request response
func (o *CreateBindBadRequest) WithPayload(payload *models.Error) *CreateBindBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bind bad request response
func (o *CreateBindBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBindBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBindConflictCode is the HTTP code returned for type CreateBindConflict
const CreateBindConflictCode int = 409

/*CreateBindConflict The specified resource already exists

swagger:response createBindConflict
*/
type CreateBindConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBindConflict creates CreateBindConflict with default headers values
func NewCreateBindConflict() *CreateBindConflict {

	return &CreateBindConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create bind conflict response
func (o *CreateBindConflict) WithConfigurationVersion(configurationVersion string) *CreateBindConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create bind conflict response
func (o *CreateBindConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create bind conflict response
func (o *CreateBindConflict) WithPayload(payload *models.Error) *CreateBindConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bind conflict response
func (o *CreateBindConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBindConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateBindDefault General Error

swagger:response createBindDefault
*/
type CreateBindDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBindDefault creates CreateBindDefault with default headers values
func NewCreateBindDefault(code int) *CreateBindDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateBindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create bind default response
func (o *CreateBindDefault) WithStatusCode(code int) *CreateBindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create bind default response
func (o *CreateBindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create bind default response
func (o *CreateBindDefault) WithConfigurationVersion(configurationVersion string) *CreateBindDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create bind default response
func (o *CreateBindDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create bind default response
func (o *CreateBindDefault) WithPayload(payload *models.Error) *CreateBindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bind default response
func (o *CreateBindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
