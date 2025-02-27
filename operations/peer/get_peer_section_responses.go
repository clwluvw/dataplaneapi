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

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetPeerSectionOKCode is the HTTP code returned for type GetPeerSectionOK
const GetPeerSectionOKCode int = 200

/*GetPeerSectionOK Successful operation

swagger:response getPeerSectionOK
*/
type GetPeerSectionOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetPeerSectionOKBody `json:"body,omitempty"`
}

// NewGetPeerSectionOK creates GetPeerSectionOK with default headers values
func NewGetPeerSectionOK() *GetPeerSectionOK {

	return &GetPeerSectionOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get peer section o k response
func (o *GetPeerSectionOK) WithConfigurationVersion(configurationVersion string) *GetPeerSectionOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer section o k response
func (o *GetPeerSectionOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer section o k response
func (o *GetPeerSectionOK) WithPayload(payload *GetPeerSectionOKBody) *GetPeerSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer section o k response
func (o *GetPeerSectionOK) SetPayload(payload *GetPeerSectionOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPeerSectionNotFoundCode is the HTTP code returned for type GetPeerSectionNotFound
const GetPeerSectionNotFoundCode int = 404

/*GetPeerSectionNotFound The specified resource was not found

swagger:response getPeerSectionNotFound
*/
type GetPeerSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPeerSectionNotFound creates GetPeerSectionNotFound with default headers values
func NewGetPeerSectionNotFound() *GetPeerSectionNotFound {

	return &GetPeerSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get peer section not found response
func (o *GetPeerSectionNotFound) WithConfigurationVersion(configurationVersion string) *GetPeerSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer section not found response
func (o *GetPeerSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer section not found response
func (o *GetPeerSectionNotFound) WithPayload(payload *models.Error) *GetPeerSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer section not found response
func (o *GetPeerSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetPeerSectionDefault General Error

swagger:response getPeerSectionDefault
*/
type GetPeerSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPeerSectionDefault creates GetPeerSectionDefault with default headers values
func NewGetPeerSectionDefault(code int) *GetPeerSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPeerSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get peer section default response
func (o *GetPeerSectionDefault) WithStatusCode(code int) *GetPeerSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get peer section default response
func (o *GetPeerSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get peer section default response
func (o *GetPeerSectionDefault) WithConfigurationVersion(configurationVersion string) *GetPeerSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer section default response
func (o *GetPeerSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer section default response
func (o *GetPeerSectionDefault) WithPayload(payload *models.Error) *GetPeerSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer section default response
func (o *GetPeerSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
