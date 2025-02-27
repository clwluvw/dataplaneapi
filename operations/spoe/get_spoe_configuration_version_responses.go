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

// GetSpoeConfigurationVersionOKCode is the HTTP code returned for type GetSpoeConfigurationVersionOK
const GetSpoeConfigurationVersionOKCode int = 200

/*GetSpoeConfigurationVersionOK SPOE configuration version

swagger:response getSpoeConfigurationVersionOK
*/
type GetSpoeConfigurationVersionOK struct {

	/*
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewGetSpoeConfigurationVersionOK creates GetSpoeConfigurationVersionOK with default headers values
func NewGetSpoeConfigurationVersionOK() *GetSpoeConfigurationVersionOK {

	return &GetSpoeConfigurationVersionOK{}
}

// WithPayload adds the payload to the get spoe configuration version o k response
func (o *GetSpoeConfigurationVersionOK) WithPayload(payload int64) *GetSpoeConfigurationVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe configuration version o k response
func (o *GetSpoeConfigurationVersionOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeConfigurationVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSpoeConfigurationVersionNotFoundCode is the HTTP code returned for type GetSpoeConfigurationVersionNotFound
const GetSpoeConfigurationVersionNotFoundCode int = 404

/*GetSpoeConfigurationVersionNotFound The specified resource was not found

swagger:response getSpoeConfigurationVersionNotFound
*/
type GetSpoeConfigurationVersionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeConfigurationVersionNotFound creates GetSpoeConfigurationVersionNotFound with default headers values
func NewGetSpoeConfigurationVersionNotFound() *GetSpoeConfigurationVersionNotFound {

	return &GetSpoeConfigurationVersionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe configuration version not found response
func (o *GetSpoeConfigurationVersionNotFound) WithConfigurationVersion(configurationVersion string) *GetSpoeConfigurationVersionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe configuration version not found response
func (o *GetSpoeConfigurationVersionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe configuration version not found response
func (o *GetSpoeConfigurationVersionNotFound) WithPayload(payload *models.Error) *GetSpoeConfigurationVersionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe configuration version not found response
func (o *GetSpoeConfigurationVersionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeConfigurationVersionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetSpoeConfigurationVersionDefault General Error

swagger:response getSpoeConfigurationVersionDefault
*/
type GetSpoeConfigurationVersionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeConfigurationVersionDefault creates GetSpoeConfigurationVersionDefault with default headers values
func NewGetSpoeConfigurationVersionDefault(code int) *GetSpoeConfigurationVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoeConfigurationVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoe configuration version default response
func (o *GetSpoeConfigurationVersionDefault) WithStatusCode(code int) *GetSpoeConfigurationVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe configuration version default response
func (o *GetSpoeConfigurationVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe configuration version default response
func (o *GetSpoeConfigurationVersionDefault) WithConfigurationVersion(configurationVersion string) *GetSpoeConfigurationVersionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe configuration version default response
func (o *GetSpoeConfigurationVersionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe configuration version default response
func (o *GetSpoeConfigurationVersionDefault) WithPayload(payload *models.Error) *GetSpoeConfigurationVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe configuration version default response
func (o *GetSpoeConfigurationVersionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeConfigurationVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
