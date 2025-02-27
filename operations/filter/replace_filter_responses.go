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

// ReplaceFilterOKCode is the HTTP code returned for type ReplaceFilterOK
const ReplaceFilterOKCode int = 200

/*ReplaceFilterOK Filter replaced

swagger:response replaceFilterOK
*/
type ReplaceFilterOK struct {

	/*
	  In: Body
	*/
	Payload *models.Filter `json:"body,omitempty"`
}

// NewReplaceFilterOK creates ReplaceFilterOK with default headers values
func NewReplaceFilterOK() *ReplaceFilterOK {

	return &ReplaceFilterOK{}
}

// WithPayload adds the payload to the replace filter o k response
func (o *ReplaceFilterOK) WithPayload(payload *models.Filter) *ReplaceFilterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace filter o k response
func (o *ReplaceFilterOK) SetPayload(payload *models.Filter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFilterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceFilterAcceptedCode is the HTTP code returned for type ReplaceFilterAccepted
const ReplaceFilterAcceptedCode int = 202

/*ReplaceFilterAccepted Configuration change accepted and reload requested

swagger:response replaceFilterAccepted
*/
type ReplaceFilterAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Filter `json:"body,omitempty"`
}

// NewReplaceFilterAccepted creates ReplaceFilterAccepted with default headers values
func NewReplaceFilterAccepted() *ReplaceFilterAccepted {

	return &ReplaceFilterAccepted{}
}

// WithReloadID adds the reloadId to the replace filter accepted response
func (o *ReplaceFilterAccepted) WithReloadID(reloadID string) *ReplaceFilterAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace filter accepted response
func (o *ReplaceFilterAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace filter accepted response
func (o *ReplaceFilterAccepted) WithPayload(payload *models.Filter) *ReplaceFilterAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace filter accepted response
func (o *ReplaceFilterAccepted) SetPayload(payload *models.Filter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFilterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceFilterBadRequestCode is the HTTP code returned for type ReplaceFilterBadRequest
const ReplaceFilterBadRequestCode int = 400

/*ReplaceFilterBadRequest Bad request

swagger:response replaceFilterBadRequest
*/
type ReplaceFilterBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceFilterBadRequest creates ReplaceFilterBadRequest with default headers values
func NewReplaceFilterBadRequest() *ReplaceFilterBadRequest {

	return &ReplaceFilterBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace filter bad request response
func (o *ReplaceFilterBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceFilterBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace filter bad request response
func (o *ReplaceFilterBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace filter bad request response
func (o *ReplaceFilterBadRequest) WithPayload(payload *models.Error) *ReplaceFilterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace filter bad request response
func (o *ReplaceFilterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFilterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceFilterNotFoundCode is the HTTP code returned for type ReplaceFilterNotFound
const ReplaceFilterNotFoundCode int = 404

/*ReplaceFilterNotFound The specified resource was not found

swagger:response replaceFilterNotFound
*/
type ReplaceFilterNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceFilterNotFound creates ReplaceFilterNotFound with default headers values
func NewReplaceFilterNotFound() *ReplaceFilterNotFound {

	return &ReplaceFilterNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace filter not found response
func (o *ReplaceFilterNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceFilterNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace filter not found response
func (o *ReplaceFilterNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace filter not found response
func (o *ReplaceFilterNotFound) WithPayload(payload *models.Error) *ReplaceFilterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace filter not found response
func (o *ReplaceFilterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFilterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceFilterDefault General Error

swagger:response replaceFilterDefault
*/
type ReplaceFilterDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceFilterDefault creates ReplaceFilterDefault with default headers values
func NewReplaceFilterDefault(code int) *ReplaceFilterDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceFilterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace filter default response
func (o *ReplaceFilterDefault) WithStatusCode(code int) *ReplaceFilterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace filter default response
func (o *ReplaceFilterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace filter default response
func (o *ReplaceFilterDefault) WithConfigurationVersion(configurationVersion string) *ReplaceFilterDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace filter default response
func (o *ReplaceFilterDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace filter default response
func (o *ReplaceFilterDefault) WithPayload(payload *models.Error) *ReplaceFilterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace filter default response
func (o *ReplaceFilterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFilterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
