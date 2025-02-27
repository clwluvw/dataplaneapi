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

package http_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetHTTPChecksOKCode is the HTTP code returned for type GetHTTPChecksOK
const GetHTTPChecksOKCode int = 200

/*GetHTTPChecksOK Successful operation

swagger:response getHttpChecksOK
*/
type GetHTTPChecksOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHTTPChecksOKBody `json:"body,omitempty"`
}

// NewGetHTTPChecksOK creates GetHTTPChecksOK with default headers values
func NewGetHTTPChecksOK() *GetHTTPChecksOK {

	return &GetHTTPChecksOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http checks o k response
func (o *GetHTTPChecksOK) WithConfigurationVersion(configurationVersion string) *GetHTTPChecksOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http checks o k response
func (o *GetHTTPChecksOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http checks o k response
func (o *GetHTTPChecksOK) WithPayload(payload *GetHTTPChecksOKBody) *GetHTTPChecksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http checks o k response
func (o *GetHTTPChecksOK) SetPayload(payload *GetHTTPChecksOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPChecksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetHTTPChecksDefault General Error

swagger:response getHttpChecksDefault
*/
type GetHTTPChecksDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPChecksDefault creates GetHTTPChecksDefault with default headers values
func NewGetHTTPChecksDefault(code int) *GetHTTPChecksDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPChecksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP checks default response
func (o *GetHTTPChecksDefault) WithStatusCode(code int) *GetHTTPChecksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP checks default response
func (o *GetHTTPChecksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP checks default response
func (o *GetHTTPChecksDefault) WithConfigurationVersion(configurationVersion string) *GetHTTPChecksDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP checks default response
func (o *GetHTTPChecksDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP checks default response
func (o *GetHTTPChecksDefault) WithPayload(payload *models.Error) *GetHTTPChecksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP checks default response
func (o *GetHTTPChecksDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPChecksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
