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

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// PostHAProxyConfigurationCreatedCode is the HTTP code returned for type PostHAProxyConfigurationCreated
const PostHAProxyConfigurationCreatedCode int = 201

/*PostHAProxyConfigurationCreated New HAProxy configuration pushed

swagger:response postHAProxyConfigurationCreated
*/
type PostHAProxyConfigurationCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostHAProxyConfigurationCreated creates PostHAProxyConfigurationCreated with default headers values
func NewPostHAProxyConfigurationCreated() *PostHAProxyConfigurationCreated {

	return &PostHAProxyConfigurationCreated{}
}

// WithPayload adds the payload to the post h a proxy configuration created response
func (o *PostHAProxyConfigurationCreated) WithPayload(payload string) *PostHAProxyConfigurationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post h a proxy configuration created response
func (o *PostHAProxyConfigurationCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostHAProxyConfigurationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostHAProxyConfigurationAcceptedCode is the HTTP code returned for type PostHAProxyConfigurationAccepted
const PostHAProxyConfigurationAcceptedCode int = 202

/*PostHAProxyConfigurationAccepted Configuration change accepted and reload requested

swagger:response postHAProxyConfigurationAccepted
*/
type PostHAProxyConfigurationAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostHAProxyConfigurationAccepted creates PostHAProxyConfigurationAccepted with default headers values
func NewPostHAProxyConfigurationAccepted() *PostHAProxyConfigurationAccepted {

	return &PostHAProxyConfigurationAccepted{}
}

// WithReloadID adds the reloadId to the post h a proxy configuration accepted response
func (o *PostHAProxyConfigurationAccepted) WithReloadID(reloadID string) *PostHAProxyConfigurationAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the post h a proxy configuration accepted response
func (o *PostHAProxyConfigurationAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the post h a proxy configuration accepted response
func (o *PostHAProxyConfigurationAccepted) WithPayload(payload string) *PostHAProxyConfigurationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post h a proxy configuration accepted response
func (o *PostHAProxyConfigurationAccepted) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostHAProxyConfigurationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostHAProxyConfigurationBadRequestCode is the HTTP code returned for type PostHAProxyConfigurationBadRequest
const PostHAProxyConfigurationBadRequestCode int = 400

/*PostHAProxyConfigurationBadRequest Bad request

swagger:response postHAProxyConfigurationBadRequest
*/
type PostHAProxyConfigurationBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostHAProxyConfigurationBadRequest creates PostHAProxyConfigurationBadRequest with default headers values
func NewPostHAProxyConfigurationBadRequest() *PostHAProxyConfigurationBadRequest {

	return &PostHAProxyConfigurationBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the post h a proxy configuration bad request response
func (o *PostHAProxyConfigurationBadRequest) WithConfigurationVersion(configurationVersion string) *PostHAProxyConfigurationBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the post h a proxy configuration bad request response
func (o *PostHAProxyConfigurationBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the post h a proxy configuration bad request response
func (o *PostHAProxyConfigurationBadRequest) WithPayload(payload *models.Error) *PostHAProxyConfigurationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post h a proxy configuration bad request response
func (o *PostHAProxyConfigurationBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostHAProxyConfigurationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*PostHAProxyConfigurationDefault General Error

swagger:response postHAProxyConfigurationDefault
*/
type PostHAProxyConfigurationDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostHAProxyConfigurationDefault creates PostHAProxyConfigurationDefault with default headers values
func NewPostHAProxyConfigurationDefault(code int) *PostHAProxyConfigurationDefault {
	if code <= 0 {
		code = 500
	}

	return &PostHAProxyConfigurationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) WithStatusCode(code int) *PostHAProxyConfigurationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) WithConfigurationVersion(configurationVersion string) *PostHAProxyConfigurationDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) WithPayload(payload *models.Error) *PostHAProxyConfigurationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostHAProxyConfigurationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
