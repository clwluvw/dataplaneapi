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

// GetServicesHaproxyRuntimeACLFileEntriesIDOKCode is the HTTP code returned for type GetServicesHaproxyRuntimeACLFileEntriesIDOK
const GetServicesHaproxyRuntimeACLFileEntriesIDOKCode int = 200

/*GetServicesHaproxyRuntimeACLFileEntriesIDOK Successful operation

swagger:response getServicesHaproxyRuntimeAclFileEntriesIdOK
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ACLFileEntry `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeACLFileEntriesIDOK creates GetServicesHaproxyRuntimeACLFileEntriesIDOK with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDOK() *GetServicesHaproxyRuntimeACLFileEntriesIDOK {

	return &GetServicesHaproxyRuntimeACLFileEntriesIDOK{}
}

// WithPayload adds the payload to the get services haproxy runtime Acl file entries Id o k response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDOK) WithPayload(payload *models.ACLFileEntry) *GetServicesHaproxyRuntimeACLFileEntriesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime Acl file entries Id o k response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDOK) SetPayload(payload *models.ACLFileEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServicesHaproxyRuntimeACLFileEntriesIDBadRequestCode is the HTTP code returned for type GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest
const GetServicesHaproxyRuntimeACLFileEntriesIDBadRequestCode int = 400

/*GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest Bad request

swagger:response getServicesHaproxyRuntimeAclFileEntriesIdBadRequest
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeACLFileEntriesIDBadRequest creates GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDBadRequest() *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest {

	return &GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the get services haproxy runtime Acl file entries Id bad request response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) WithConfigurationVersion(configurationVersion string) *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get services haproxy runtime Acl file entries Id bad request response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get services haproxy runtime Acl file entries Id bad request response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) WithPayload(payload *models.Error) *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime Acl file entries Id bad request response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetServicesHaproxyRuntimeACLFileEntriesIDNotFoundCode is the HTTP code returned for type GetServicesHaproxyRuntimeACLFileEntriesIDNotFound
const GetServicesHaproxyRuntimeACLFileEntriesIDNotFoundCode int = 404

/*GetServicesHaproxyRuntimeACLFileEntriesIDNotFound The specified resource was not found

swagger:response getServicesHaproxyRuntimeAclFileEntriesIdNotFound
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeACLFileEntriesIDNotFound creates GetServicesHaproxyRuntimeACLFileEntriesIDNotFound with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDNotFound() *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound {

	return &GetServicesHaproxyRuntimeACLFileEntriesIDNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get services haproxy runtime Acl file entries Id not found response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) WithConfigurationVersion(configurationVersion string) *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get services haproxy runtime Acl file entries Id not found response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get services haproxy runtime Acl file entries Id not found response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) WithPayload(payload *models.Error) *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime Acl file entries Id not found response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetServicesHaproxyRuntimeACLFileEntriesIDDefault General Error

swagger:response getServicesHaproxyRuntimeAclFileEntriesIdDefault
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServicesHaproxyRuntimeACLFileEntriesIDDefault creates GetServicesHaproxyRuntimeACLFileEntriesIDDefault with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDDefault(code int) *GetServicesHaproxyRuntimeACLFileEntriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServicesHaproxyRuntimeACLFileEntriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) WithStatusCode(code int) *GetServicesHaproxyRuntimeACLFileEntriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) WithConfigurationVersion(configurationVersion string) *GetServicesHaproxyRuntimeACLFileEntriesIDDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) WithPayload(payload *models.Error) *GetServicesHaproxyRuntimeACLFileEntriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
