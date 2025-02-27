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

package userlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetUserlistOKCode is the HTTP code returned for type GetUserlistOK
const GetUserlistOKCode int = 200

/*GetUserlistOK Successful operation

swagger:response getUserlistOK
*/
type GetUserlistOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetUserlistOKBody `json:"body,omitempty"`
}

// NewGetUserlistOK creates GetUserlistOK with default headers values
func NewGetUserlistOK() *GetUserlistOK {

	return &GetUserlistOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get userlist o k response
func (o *GetUserlistOK) WithConfigurationVersion(configurationVersion string) *GetUserlistOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get userlist o k response
func (o *GetUserlistOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get userlist o k response
func (o *GetUserlistOK) WithPayload(payload *GetUserlistOKBody) *GetUserlistOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get userlist o k response
func (o *GetUserlistOK) SetPayload(payload *GetUserlistOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserlistOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUserlistNotFoundCode is the HTTP code returned for type GetUserlistNotFound
const GetUserlistNotFoundCode int = 404

/*GetUserlistNotFound The specified resource already exists

swagger:response getUserlistNotFound
*/
type GetUserlistNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserlistNotFound creates GetUserlistNotFound with default headers values
func NewGetUserlistNotFound() *GetUserlistNotFound {

	return &GetUserlistNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get userlist not found response
func (o *GetUserlistNotFound) WithConfigurationVersion(configurationVersion string) *GetUserlistNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get userlist not found response
func (o *GetUserlistNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get userlist not found response
func (o *GetUserlistNotFound) WithPayload(payload *models.Error) *GetUserlistNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get userlist not found response
func (o *GetUserlistNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserlistNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetUserlistDefault General Error

swagger:response getUserlistDefault
*/
type GetUserlistDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserlistDefault creates GetUserlistDefault with default headers values
func NewGetUserlistDefault(code int) *GetUserlistDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserlistDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get userlist default response
func (o *GetUserlistDefault) WithStatusCode(code int) *GetUserlistDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get userlist default response
func (o *GetUserlistDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get userlist default response
func (o *GetUserlistDefault) WithConfigurationVersion(configurationVersion string) *GetUserlistDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get userlist default response
func (o *GetUserlistDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get userlist default response
func (o *GetUserlistDefault) WithPayload(payload *models.Error) *GetUserlistDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get userlist default response
func (o *GetUserlistDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserlistDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
