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

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateFilterCreatedCode is the HTTP code returned for type CreateFilterCreated
const CreateFilterCreatedCode int = 201

/*CreateFilterCreated Filter created

swagger:response createFilterCreated
*/
type CreateFilterCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Filter `json:"body,omitempty"`
}

// NewCreateFilterCreated creates CreateFilterCreated with default headers values
func NewCreateFilterCreated() *CreateFilterCreated {

	return &CreateFilterCreated{}
}

// WithPayload adds the payload to the create filter created response
func (o *CreateFilterCreated) WithPayload(payload *models.Filter) *CreateFilterCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create filter created response
func (o *CreateFilterCreated) SetPayload(payload *models.Filter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFilterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFilterAcceptedCode is the HTTP code returned for type CreateFilterAccepted
const CreateFilterAcceptedCode int = 202

/*CreateFilterAccepted Configuration change accepted and reload requested

swagger:response createFilterAccepted
*/
type CreateFilterAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Filter `json:"body,omitempty"`
}

// NewCreateFilterAccepted creates CreateFilterAccepted with default headers values
func NewCreateFilterAccepted() *CreateFilterAccepted {

	return &CreateFilterAccepted{}
}

// WithReloadID adds the reloadId to the create filter accepted response
func (o *CreateFilterAccepted) WithReloadID(reloadID string) *CreateFilterAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create filter accepted response
func (o *CreateFilterAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create filter accepted response
func (o *CreateFilterAccepted) WithPayload(payload *models.Filter) *CreateFilterAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create filter accepted response
func (o *CreateFilterAccepted) SetPayload(payload *models.Filter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFilterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateFilterBadRequestCode is the HTTP code returned for type CreateFilterBadRequest
const CreateFilterBadRequestCode int = 400

/*CreateFilterBadRequest Bad request

swagger:response createFilterBadRequest
*/
type CreateFilterBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFilterBadRequest creates CreateFilterBadRequest with default headers values
func NewCreateFilterBadRequest() *CreateFilterBadRequest {

	return &CreateFilterBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create filter bad request response
func (o *CreateFilterBadRequest) WithConfigurationVersion(configurationVersion string) *CreateFilterBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create filter bad request response
func (o *CreateFilterBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create filter bad request response
func (o *CreateFilterBadRequest) WithPayload(payload *models.Error) *CreateFilterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create filter bad request response
func (o *CreateFilterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFilterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateFilterConflictCode is the HTTP code returned for type CreateFilterConflict
const CreateFilterConflictCode int = 409

/*CreateFilterConflict The specified resource already exists

swagger:response createFilterConflict
*/
type CreateFilterConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFilterConflict creates CreateFilterConflict with default headers values
func NewCreateFilterConflict() *CreateFilterConflict {

	return &CreateFilterConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create filter conflict response
func (o *CreateFilterConflict) WithConfigurationVersion(configurationVersion string) *CreateFilterConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create filter conflict response
func (o *CreateFilterConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create filter conflict response
func (o *CreateFilterConflict) WithPayload(payload *models.Error) *CreateFilterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create filter conflict response
func (o *CreateFilterConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFilterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateFilterDefault General Error

swagger:response createFilterDefault
*/
type CreateFilterDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFilterDefault creates CreateFilterDefault with default headers values
func NewCreateFilterDefault(code int) *CreateFilterDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateFilterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create filter default response
func (o *CreateFilterDefault) WithStatusCode(code int) *CreateFilterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create filter default response
func (o *CreateFilterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create filter default response
func (o *CreateFilterDefault) WithConfigurationVersion(configurationVersion string) *CreateFilterDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create filter default response
func (o *CreateFilterDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create filter default response
func (o *CreateFilterDefault) WithPayload(payload *models.Error) *CreateFilterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create filter default response
func (o *CreateFilterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFilterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
