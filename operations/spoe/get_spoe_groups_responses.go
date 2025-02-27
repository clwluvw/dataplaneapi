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

// GetSpoeGroupsOKCode is the HTTP code returned for type GetSpoeGroupsOK
const GetSpoeGroupsOKCode int = 200

/*GetSpoeGroupsOK Successful operation

swagger:response getSpoeGroupsOK
*/
type GetSpoeGroupsOK struct {
	/*Spoe configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetSpoeGroupsOKBody `json:"body,omitempty"`
}

// NewGetSpoeGroupsOK creates GetSpoeGroupsOK with default headers values
func NewGetSpoeGroupsOK() *GetSpoeGroupsOK {

	return &GetSpoeGroupsOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe groups o k response
func (o *GetSpoeGroupsOK) WithConfigurationVersion(configurationVersion string) *GetSpoeGroupsOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe groups o k response
func (o *GetSpoeGroupsOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe groups o k response
func (o *GetSpoeGroupsOK) WithPayload(payload *GetSpoeGroupsOKBody) *GetSpoeGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe groups o k response
func (o *GetSpoeGroupsOK) SetPayload(payload *GetSpoeGroupsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetSpoeGroupsDefault General Error

swagger:response getSpoeGroupsDefault
*/
type GetSpoeGroupsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeGroupsDefault creates GetSpoeGroupsDefault with default headers values
func NewGetSpoeGroupsDefault(code int) *GetSpoeGroupsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoeGroupsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoe groups default response
func (o *GetSpoeGroupsDefault) WithStatusCode(code int) *GetSpoeGroupsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe groups default response
func (o *GetSpoeGroupsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe groups default response
func (o *GetSpoeGroupsDefault) WithConfigurationVersion(configurationVersion string) *GetSpoeGroupsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe groups default response
func (o *GetSpoeGroupsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe groups default response
func (o *GetSpoeGroupsDefault) WithPayload(payload *models.Error) *GetSpoeGroupsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe groups default response
func (o *GetSpoeGroupsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeGroupsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
