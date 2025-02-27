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

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateDeclareCaptureCreatedCode is the HTTP code returned for type CreateDeclareCaptureCreated
const CreateDeclareCaptureCreatedCode int = 201

/*CreateDeclareCaptureCreated Declare capture created

swagger:response createDeclareCaptureCreated
*/
type CreateDeclareCaptureCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Capture `json:"body,omitempty"`
}

// NewCreateDeclareCaptureCreated creates CreateDeclareCaptureCreated with default headers values
func NewCreateDeclareCaptureCreated() *CreateDeclareCaptureCreated {

	return &CreateDeclareCaptureCreated{}
}

// WithPayload adds the payload to the create declare capture created response
func (o *CreateDeclareCaptureCreated) WithPayload(payload *models.Capture) *CreateDeclareCaptureCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create declare capture created response
func (o *CreateDeclareCaptureCreated) SetPayload(payload *models.Capture) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeclareCaptureCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDeclareCaptureAcceptedCode is the HTTP code returned for type CreateDeclareCaptureAccepted
const CreateDeclareCaptureAcceptedCode int = 202

/*CreateDeclareCaptureAccepted Configuration change accepted and reload requested

swagger:response createDeclareCaptureAccepted
*/
type CreateDeclareCaptureAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Capture `json:"body,omitempty"`
}

// NewCreateDeclareCaptureAccepted creates CreateDeclareCaptureAccepted with default headers values
func NewCreateDeclareCaptureAccepted() *CreateDeclareCaptureAccepted {

	return &CreateDeclareCaptureAccepted{}
}

// WithReloadID adds the reloadId to the create declare capture accepted response
func (o *CreateDeclareCaptureAccepted) WithReloadID(reloadID string) *CreateDeclareCaptureAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create declare capture accepted response
func (o *CreateDeclareCaptureAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create declare capture accepted response
func (o *CreateDeclareCaptureAccepted) WithPayload(payload *models.Capture) *CreateDeclareCaptureAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create declare capture accepted response
func (o *CreateDeclareCaptureAccepted) SetPayload(payload *models.Capture) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeclareCaptureAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateDeclareCaptureBadRequestCode is the HTTP code returned for type CreateDeclareCaptureBadRequest
const CreateDeclareCaptureBadRequestCode int = 400

/*CreateDeclareCaptureBadRequest Bad request

swagger:response createDeclareCaptureBadRequest
*/
type CreateDeclareCaptureBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDeclareCaptureBadRequest creates CreateDeclareCaptureBadRequest with default headers values
func NewCreateDeclareCaptureBadRequest() *CreateDeclareCaptureBadRequest {

	return &CreateDeclareCaptureBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create declare capture bad request response
func (o *CreateDeclareCaptureBadRequest) WithConfigurationVersion(configurationVersion string) *CreateDeclareCaptureBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create declare capture bad request response
func (o *CreateDeclareCaptureBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create declare capture bad request response
func (o *CreateDeclareCaptureBadRequest) WithPayload(payload *models.Error) *CreateDeclareCaptureBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create declare capture bad request response
func (o *CreateDeclareCaptureBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeclareCaptureBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateDeclareCaptureConflictCode is the HTTP code returned for type CreateDeclareCaptureConflict
const CreateDeclareCaptureConflictCode int = 409

/*CreateDeclareCaptureConflict The specified resource already exists

swagger:response createDeclareCaptureConflict
*/
type CreateDeclareCaptureConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDeclareCaptureConflict creates CreateDeclareCaptureConflict with default headers values
func NewCreateDeclareCaptureConflict() *CreateDeclareCaptureConflict {

	return &CreateDeclareCaptureConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create declare capture conflict response
func (o *CreateDeclareCaptureConflict) WithConfigurationVersion(configurationVersion string) *CreateDeclareCaptureConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create declare capture conflict response
func (o *CreateDeclareCaptureConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create declare capture conflict response
func (o *CreateDeclareCaptureConflict) WithPayload(payload *models.Error) *CreateDeclareCaptureConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create declare capture conflict response
func (o *CreateDeclareCaptureConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeclareCaptureConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateDeclareCaptureDefault General Error

swagger:response createDeclareCaptureDefault
*/
type CreateDeclareCaptureDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDeclareCaptureDefault creates CreateDeclareCaptureDefault with default headers values
func NewCreateDeclareCaptureDefault(code int) *CreateDeclareCaptureDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDeclareCaptureDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create declare capture default response
func (o *CreateDeclareCaptureDefault) WithStatusCode(code int) *CreateDeclareCaptureDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create declare capture default response
func (o *CreateDeclareCaptureDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create declare capture default response
func (o *CreateDeclareCaptureDefault) WithConfigurationVersion(configurationVersion string) *CreateDeclareCaptureDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create declare capture default response
func (o *CreateDeclareCaptureDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create declare capture default response
func (o *CreateDeclareCaptureDefault) WithPayload(payload *models.Error) *CreateDeclareCaptureDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create declare capture default response
func (o *CreateDeclareCaptureDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeclareCaptureDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
