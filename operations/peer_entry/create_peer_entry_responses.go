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

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreatePeerEntryCreatedCode is the HTTP code returned for type CreatePeerEntryCreated
const CreatePeerEntryCreatedCode int = 201

/*CreatePeerEntryCreated PeerEntry created

swagger:response createPeerEntryCreated
*/
type CreatePeerEntryCreated struct {

	/*
	  In: Body
	*/
	Payload *models.PeerEntry `json:"body,omitempty"`
}

// NewCreatePeerEntryCreated creates CreatePeerEntryCreated with default headers values
func NewCreatePeerEntryCreated() *CreatePeerEntryCreated {

	return &CreatePeerEntryCreated{}
}

// WithPayload adds the payload to the create peer entry created response
func (o *CreatePeerEntryCreated) WithPayload(payload *models.PeerEntry) *CreatePeerEntryCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer entry created response
func (o *CreatePeerEntryCreated) SetPayload(payload *models.PeerEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerEntryCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePeerEntryAcceptedCode is the HTTP code returned for type CreatePeerEntryAccepted
const CreatePeerEntryAcceptedCode int = 202

/*CreatePeerEntryAccepted Configuration change accepted and reload requested

swagger:response createPeerEntryAccepted
*/
type CreatePeerEntryAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.PeerEntry `json:"body,omitempty"`
}

// NewCreatePeerEntryAccepted creates CreatePeerEntryAccepted with default headers values
func NewCreatePeerEntryAccepted() *CreatePeerEntryAccepted {

	return &CreatePeerEntryAccepted{}
}

// WithReloadID adds the reloadId to the create peer entry accepted response
func (o *CreatePeerEntryAccepted) WithReloadID(reloadID string) *CreatePeerEntryAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create peer entry accepted response
func (o *CreatePeerEntryAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create peer entry accepted response
func (o *CreatePeerEntryAccepted) WithPayload(payload *models.PeerEntry) *CreatePeerEntryAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer entry accepted response
func (o *CreatePeerEntryAccepted) SetPayload(payload *models.PeerEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerEntryAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreatePeerEntryBadRequestCode is the HTTP code returned for type CreatePeerEntryBadRequest
const CreatePeerEntryBadRequestCode int = 400

/*CreatePeerEntryBadRequest Bad request

swagger:response createPeerEntryBadRequest
*/
type CreatePeerEntryBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePeerEntryBadRequest creates CreatePeerEntryBadRequest with default headers values
func NewCreatePeerEntryBadRequest() *CreatePeerEntryBadRequest {

	return &CreatePeerEntryBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create peer entry bad request response
func (o *CreatePeerEntryBadRequest) WithConfigurationVersion(configurationVersion string) *CreatePeerEntryBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create peer entry bad request response
func (o *CreatePeerEntryBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create peer entry bad request response
func (o *CreatePeerEntryBadRequest) WithPayload(payload *models.Error) *CreatePeerEntryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer entry bad request response
func (o *CreatePeerEntryBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerEntryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreatePeerEntryConflictCode is the HTTP code returned for type CreatePeerEntryConflict
const CreatePeerEntryConflictCode int = 409

/*CreatePeerEntryConflict The specified resource already exists

swagger:response createPeerEntryConflict
*/
type CreatePeerEntryConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePeerEntryConflict creates CreatePeerEntryConflict with default headers values
func NewCreatePeerEntryConflict() *CreatePeerEntryConflict {

	return &CreatePeerEntryConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create peer entry conflict response
func (o *CreatePeerEntryConflict) WithConfigurationVersion(configurationVersion string) *CreatePeerEntryConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create peer entry conflict response
func (o *CreatePeerEntryConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create peer entry conflict response
func (o *CreatePeerEntryConflict) WithPayload(payload *models.Error) *CreatePeerEntryConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer entry conflict response
func (o *CreatePeerEntryConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerEntryConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreatePeerEntryDefault General Error

swagger:response createPeerEntryDefault
*/
type CreatePeerEntryDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePeerEntryDefault creates CreatePeerEntryDefault with default headers values
func NewCreatePeerEntryDefault(code int) *CreatePeerEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &CreatePeerEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create peer entry default response
func (o *CreatePeerEntryDefault) WithStatusCode(code int) *CreatePeerEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create peer entry default response
func (o *CreatePeerEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create peer entry default response
func (o *CreatePeerEntryDefault) WithConfigurationVersion(configurationVersion string) *CreatePeerEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create peer entry default response
func (o *CreatePeerEntryDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create peer entry default response
func (o *CreatePeerEntryDefault) WithPayload(payload *models.Error) *CreatePeerEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer entry default response
func (o *CreatePeerEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
