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

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetServicesHaproxyRuntimeAclsIDOKCode is the HTTP code returned for type GetServicesHaproxyRuntimeAclsIDOK
const GetServicesHaproxyRuntimeAclsIDOKCode int = 200

/*GetServicesHaproxyRuntimeAclsIDOK Successful operation

swagger:response getServicesHaproxyRuntimeAclsIdOK
*/
type GetServicesHaproxyRuntimeAclsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ACLFile `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeAclsIDOK creates GetServicesHaproxyRuntimeAclsIDOK with default headers values
func NewGetServicesHaproxyRuntimeAclsIDOK() *GetServicesHaproxyRuntimeAclsIDOK {

	return &GetServicesHaproxyRuntimeAclsIDOK{}
}

// WithPayload adds the payload to the get services haproxy runtime acls Id o k response
func (o *GetServicesHaproxyRuntimeAclsIDOK) WithPayload(payload *models.ACLFile) *GetServicesHaproxyRuntimeAclsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime acls Id o k response
func (o *GetServicesHaproxyRuntimeAclsIDOK) SetPayload(payload *models.ACLFile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeAclsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServicesHaproxyRuntimeAclsIDNotFoundCode is the HTTP code returned for type GetServicesHaproxyRuntimeAclsIDNotFound
const GetServicesHaproxyRuntimeAclsIDNotFoundCode int = 404

/*GetServicesHaproxyRuntimeAclsIDNotFound The specified resource was not found

swagger:response getServicesHaproxyRuntimeAclsIdNotFound
*/
type GetServicesHaproxyRuntimeAclsIDNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeAclsIDNotFound creates GetServicesHaproxyRuntimeAclsIDNotFound with default headers values
func NewGetServicesHaproxyRuntimeAclsIDNotFound() *GetServicesHaproxyRuntimeAclsIDNotFound {

	return &GetServicesHaproxyRuntimeAclsIDNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get services haproxy runtime acls Id not found response
func (o *GetServicesHaproxyRuntimeAclsIDNotFound) WithConfigurationVersion(configurationVersion string) *GetServicesHaproxyRuntimeAclsIDNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get services haproxy runtime acls Id not found response
func (o *GetServicesHaproxyRuntimeAclsIDNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get services haproxy runtime acls Id not found response
func (o *GetServicesHaproxyRuntimeAclsIDNotFound) WithPayload(payload *models.Error) *GetServicesHaproxyRuntimeAclsIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime acls Id not found response
func (o *GetServicesHaproxyRuntimeAclsIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeAclsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetServicesHaproxyRuntimeAclsIDDefault General Error

swagger:response getServicesHaproxyRuntimeAclsIdDefault
*/
type GetServicesHaproxyRuntimeAclsIDDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeAclsIDDefault creates GetServicesHaproxyRuntimeAclsIDDefault with default headers values
func NewGetServicesHaproxyRuntimeAclsIDDefault(code int) *GetServicesHaproxyRuntimeAclsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServicesHaproxyRuntimeAclsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) WithStatusCode(code int) *GetServicesHaproxyRuntimeAclsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) WithConfigurationVersion(configurationVersion string) *GetServicesHaproxyRuntimeAclsIDDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) WithPayload(payload *models.Error) *GetServicesHaproxyRuntimeAclsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime acls ID default response
func (o *GetServicesHaproxyRuntimeAclsIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeAclsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
