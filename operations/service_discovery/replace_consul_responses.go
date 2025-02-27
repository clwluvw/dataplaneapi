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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// ReplaceConsulOKCode is the HTTP code returned for type ReplaceConsulOK
const ReplaceConsulOKCode int = 200

/*ReplaceConsulOK Consul server replaced

swagger:response replaceConsulOK
*/
type ReplaceConsulOK struct {

	/*
	  In: Body
	*/
	Payload *models.Consul `json:"body,omitempty"`
}

// NewReplaceConsulOK creates ReplaceConsulOK with default headers values
func NewReplaceConsulOK() *ReplaceConsulOK {

	return &ReplaceConsulOK{}
}

// WithPayload adds the payload to the replace consul o k response
func (o *ReplaceConsulOK) WithPayload(payload *models.Consul) *ReplaceConsulOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace consul o k response
func (o *ReplaceConsulOK) SetPayload(payload *models.Consul) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceConsulOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceConsulBadRequestCode is the HTTP code returned for type ReplaceConsulBadRequest
const ReplaceConsulBadRequestCode int = 400

/*ReplaceConsulBadRequest Bad request

swagger:response replaceConsulBadRequest
*/
type ReplaceConsulBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceConsulBadRequest creates ReplaceConsulBadRequest with default headers values
func NewReplaceConsulBadRequest() *ReplaceConsulBadRequest {

	return &ReplaceConsulBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace consul bad request response
func (o *ReplaceConsulBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceConsulBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace consul bad request response
func (o *ReplaceConsulBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace consul bad request response
func (o *ReplaceConsulBadRequest) WithPayload(payload *models.Error) *ReplaceConsulBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace consul bad request response
func (o *ReplaceConsulBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceConsulBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceConsulNotFoundCode is the HTTP code returned for type ReplaceConsulNotFound
const ReplaceConsulNotFoundCode int = 404

/*ReplaceConsulNotFound The specified resource was not found

swagger:response replaceConsulNotFound
*/
type ReplaceConsulNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceConsulNotFound creates ReplaceConsulNotFound with default headers values
func NewReplaceConsulNotFound() *ReplaceConsulNotFound {

	return &ReplaceConsulNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace consul not found response
func (o *ReplaceConsulNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceConsulNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace consul not found response
func (o *ReplaceConsulNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace consul not found response
func (o *ReplaceConsulNotFound) WithPayload(payload *models.Error) *ReplaceConsulNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace consul not found response
func (o *ReplaceConsulNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceConsulNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceConsulDefault General Error

swagger:response replaceConsulDefault
*/
type ReplaceConsulDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceConsulDefault creates ReplaceConsulDefault with default headers values
func NewReplaceConsulDefault(code int) *ReplaceConsulDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceConsulDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace consul default response
func (o *ReplaceConsulDefault) WithStatusCode(code int) *ReplaceConsulDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace consul default response
func (o *ReplaceConsulDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace consul default response
func (o *ReplaceConsulDefault) WithConfigurationVersion(configurationVersion string) *ReplaceConsulDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace consul default response
func (o *ReplaceConsulDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace consul default response
func (o *ReplaceConsulDefault) WithPayload(payload *models.Error) *ReplaceConsulDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace consul default response
func (o *ReplaceConsulDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceConsulDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
