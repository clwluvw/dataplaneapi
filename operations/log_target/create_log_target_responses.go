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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateLogTargetCreatedCode is the HTTP code returned for type CreateLogTargetCreated
const CreateLogTargetCreatedCode int = 201

/*CreateLogTargetCreated Log Target created

swagger:response createLogTargetCreated
*/
type CreateLogTargetCreated struct {

	/*
	  In: Body
	*/
	Payload *models.LogTarget `json:"body,omitempty"`
}

// NewCreateLogTargetCreated creates CreateLogTargetCreated with default headers values
func NewCreateLogTargetCreated() *CreateLogTargetCreated {

	return &CreateLogTargetCreated{}
}

// WithPayload adds the payload to the create log target created response
func (o *CreateLogTargetCreated) WithPayload(payload *models.LogTarget) *CreateLogTargetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log target created response
func (o *CreateLogTargetCreated) SetPayload(payload *models.LogTarget) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogTargetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateLogTargetAcceptedCode is the HTTP code returned for type CreateLogTargetAccepted
const CreateLogTargetAcceptedCode int = 202

/*CreateLogTargetAccepted Configuration change accepted and reload requested

swagger:response createLogTargetAccepted
*/
type CreateLogTargetAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.LogTarget `json:"body,omitempty"`
}

// NewCreateLogTargetAccepted creates CreateLogTargetAccepted with default headers values
func NewCreateLogTargetAccepted() *CreateLogTargetAccepted {

	return &CreateLogTargetAccepted{}
}

// WithReloadID adds the reloadId to the create log target accepted response
func (o *CreateLogTargetAccepted) WithReloadID(reloadID string) *CreateLogTargetAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create log target accepted response
func (o *CreateLogTargetAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create log target accepted response
func (o *CreateLogTargetAccepted) WithPayload(payload *models.LogTarget) *CreateLogTargetAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log target accepted response
func (o *CreateLogTargetAccepted) SetPayload(payload *models.LogTarget) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogTargetAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateLogTargetBadRequestCode is the HTTP code returned for type CreateLogTargetBadRequest
const CreateLogTargetBadRequestCode int = 400

/*CreateLogTargetBadRequest Bad request

swagger:response createLogTargetBadRequest
*/
type CreateLogTargetBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogTargetBadRequest creates CreateLogTargetBadRequest with default headers values
func NewCreateLogTargetBadRequest() *CreateLogTargetBadRequest {

	return &CreateLogTargetBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create log target bad request response
func (o *CreateLogTargetBadRequest) WithConfigurationVersion(configurationVersion string) *CreateLogTargetBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log target bad request response
func (o *CreateLogTargetBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log target bad request response
func (o *CreateLogTargetBadRequest) WithPayload(payload *models.Error) *CreateLogTargetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log target bad request response
func (o *CreateLogTargetBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogTargetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateLogTargetConflictCode is the HTTP code returned for type CreateLogTargetConflict
const CreateLogTargetConflictCode int = 409

/*CreateLogTargetConflict The specified resource already exists

swagger:response createLogTargetConflict
*/
type CreateLogTargetConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogTargetConflict creates CreateLogTargetConflict with default headers values
func NewCreateLogTargetConflict() *CreateLogTargetConflict {

	return &CreateLogTargetConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create log target conflict response
func (o *CreateLogTargetConflict) WithConfigurationVersion(configurationVersion string) *CreateLogTargetConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log target conflict response
func (o *CreateLogTargetConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log target conflict response
func (o *CreateLogTargetConflict) WithPayload(payload *models.Error) *CreateLogTargetConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log target conflict response
func (o *CreateLogTargetConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogTargetConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateLogTargetDefault General Error

swagger:response createLogTargetDefault
*/
type CreateLogTargetDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateLogTargetDefault creates CreateLogTargetDefault with default headers values
func NewCreateLogTargetDefault(code int) *CreateLogTargetDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateLogTargetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create log target default response
func (o *CreateLogTargetDefault) WithStatusCode(code int) *CreateLogTargetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create log target default response
func (o *CreateLogTargetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create log target default response
func (o *CreateLogTargetDefault) WithConfigurationVersion(configurationVersion string) *CreateLogTargetDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create log target default response
func (o *CreateLogTargetDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create log target default response
func (o *CreateLogTargetDefault) WithPayload(payload *models.Error) *CreateLogTargetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create log target default response
func (o *CreateLogTargetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateLogTargetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
