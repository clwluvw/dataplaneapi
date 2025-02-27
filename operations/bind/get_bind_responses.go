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

// GetBindOKCode is the HTTP code returned for type GetBindOK
const GetBindOKCode int = 200

/*GetBindOK Successful operation

swagger:response getBindOK
*/
type GetBindOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetBindOKBody `json:"body,omitempty"`
}

// NewGetBindOK creates GetBindOK with default headers values
func NewGetBindOK() *GetBindOK {

	return &GetBindOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get bind o k response
func (o *GetBindOK) WithConfigurationVersion(configurationVersion string) *GetBindOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get bind o k response
func (o *GetBindOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get bind o k response
func (o *GetBindOK) WithPayload(payload *GetBindOKBody) *GetBindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bind o k response
func (o *GetBindOK) SetPayload(payload *GetBindOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetBindNotFoundCode is the HTTP code returned for type GetBindNotFound
const GetBindNotFoundCode int = 404

/*GetBindNotFound The specified resource already exists

swagger:response getBindNotFound
*/
type GetBindNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBindNotFound creates GetBindNotFound with default headers values
func NewGetBindNotFound() *GetBindNotFound {

	return &GetBindNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get bind not found response
func (o *GetBindNotFound) WithConfigurationVersion(configurationVersion string) *GetBindNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get bind not found response
func (o *GetBindNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get bind not found response
func (o *GetBindNotFound) WithPayload(payload *models.Error) *GetBindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bind not found response
func (o *GetBindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetBindDefault General Error

swagger:response getBindDefault
*/
type GetBindDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBindDefault creates GetBindDefault with default headers values
func NewGetBindDefault(code int) *GetBindDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get bind default response
func (o *GetBindDefault) WithStatusCode(code int) *GetBindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get bind default response
func (o *GetBindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get bind default response
func (o *GetBindDefault) WithConfigurationVersion(configurationVersion string) *GetBindDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get bind default response
func (o *GetBindDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get bind default response
func (o *GetBindDefault) WithPayload(payload *models.Error) *GetBindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bind default response
func (o *GetBindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
