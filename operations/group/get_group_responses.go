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

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetGroupOKCode is the HTTP code returned for type GetGroupOK
const GetGroupOKCode int = 200

/*GetGroupOK Successful operation

swagger:response getGroupOK
*/
type GetGroupOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetGroupOKBody `json:"body,omitempty"`
}

// NewGetGroupOK creates GetGroupOK with default headers values
func NewGetGroupOK() *GetGroupOK {

	return &GetGroupOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get group o k response
func (o *GetGroupOK) WithConfigurationVersion(configurationVersion string) *GetGroupOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get group o k response
func (o *GetGroupOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get group o k response
func (o *GetGroupOK) WithPayload(payload *GetGroupOKBody) *GetGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group o k response
func (o *GetGroupOK) SetPayload(payload *GetGroupOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotFoundCode is the HTTP code returned for type GetGroupNotFound
const GetGroupNotFoundCode int = 404

/*GetGroupNotFound The specified resource already exists

swagger:response getGroupNotFound
*/
type GetGroupNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGroupNotFound creates GetGroupNotFound with default headers values
func NewGetGroupNotFound() *GetGroupNotFound {

	return &GetGroupNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get group not found response
func (o *GetGroupNotFound) WithConfigurationVersion(configurationVersion string) *GetGroupNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get group not found response
func (o *GetGroupNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get group not found response
func (o *GetGroupNotFound) WithPayload(payload *models.Error) *GetGroupNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group not found response
func (o *GetGroupNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetGroupDefault General Error

swagger:response getGroupDefault
*/
type GetGroupDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGroupDefault creates GetGroupDefault with default headers values
func NewGetGroupDefault(code int) *GetGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &GetGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get group default response
func (o *GetGroupDefault) WithStatusCode(code int) *GetGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get group default response
func (o *GetGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get group default response
func (o *GetGroupDefault) WithConfigurationVersion(configurationVersion string) *GetGroupDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get group default response
func (o *GetGroupDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get group default response
func (o *GetGroupDefault) WithPayload(payload *models.Error) *GetGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group default response
func (o *GetGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
