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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateSpoeAgentCreatedCode is the HTTP code returned for type CreateSpoeAgentCreated
const CreateSpoeAgentCreatedCode int = 201

/*CreateSpoeAgentCreated Spoe agent created

swagger:response createSpoeAgentCreated
*/
type CreateSpoeAgentCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SpoeAgent `json:"body,omitempty"`
}

// NewCreateSpoeAgentCreated creates CreateSpoeAgentCreated with default headers values
func NewCreateSpoeAgentCreated() *CreateSpoeAgentCreated {

	return &CreateSpoeAgentCreated{}
}

// WithPayload adds the payload to the create spoe agent created response
func (o *CreateSpoeAgentCreated) WithPayload(payload *models.SpoeAgent) *CreateSpoeAgentCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe agent created response
func (o *CreateSpoeAgentCreated) SetPayload(payload *models.SpoeAgent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeAgentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSpoeAgentBadRequestCode is the HTTP code returned for type CreateSpoeAgentBadRequest
const CreateSpoeAgentBadRequestCode int = 400

/*CreateSpoeAgentBadRequest Bad request

swagger:response createSpoeAgentBadRequest
*/
type CreateSpoeAgentBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSpoeAgentBadRequest creates CreateSpoeAgentBadRequest with default headers values
func NewCreateSpoeAgentBadRequest() *CreateSpoeAgentBadRequest {

	return &CreateSpoeAgentBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create spoe agent bad request response
func (o *CreateSpoeAgentBadRequest) WithConfigurationVersion(configurationVersion string) *CreateSpoeAgentBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create spoe agent bad request response
func (o *CreateSpoeAgentBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create spoe agent bad request response
func (o *CreateSpoeAgentBadRequest) WithPayload(payload *models.Error) *CreateSpoeAgentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe agent bad request response
func (o *CreateSpoeAgentBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeAgentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateSpoeAgentConflictCode is the HTTP code returned for type CreateSpoeAgentConflict
const CreateSpoeAgentConflictCode int = 409

/*CreateSpoeAgentConflict The specified resource already exists

swagger:response createSpoeAgentConflict
*/
type CreateSpoeAgentConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSpoeAgentConflict creates CreateSpoeAgentConflict with default headers values
func NewCreateSpoeAgentConflict() *CreateSpoeAgentConflict {

	return &CreateSpoeAgentConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create spoe agent conflict response
func (o *CreateSpoeAgentConflict) WithConfigurationVersion(configurationVersion string) *CreateSpoeAgentConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create spoe agent conflict response
func (o *CreateSpoeAgentConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create spoe agent conflict response
func (o *CreateSpoeAgentConflict) WithPayload(payload *models.Error) *CreateSpoeAgentConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe agent conflict response
func (o *CreateSpoeAgentConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeAgentConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateSpoeAgentDefault General Error

swagger:response createSpoeAgentDefault
*/
type CreateSpoeAgentDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSpoeAgentDefault creates CreateSpoeAgentDefault with default headers values
func NewCreateSpoeAgentDefault(code int) *CreateSpoeAgentDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateSpoeAgentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create spoe agent default response
func (o *CreateSpoeAgentDefault) WithStatusCode(code int) *CreateSpoeAgentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create spoe agent default response
func (o *CreateSpoeAgentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create spoe agent default response
func (o *CreateSpoeAgentDefault) WithConfigurationVersion(configurationVersion string) *CreateSpoeAgentDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create spoe agent default response
func (o *CreateSpoeAgentDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create spoe agent default response
func (o *CreateSpoeAgentDefault) WithPayload(payload *models.Error) *CreateSpoeAgentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe agent default response
func (o *CreateSpoeAgentDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeAgentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
