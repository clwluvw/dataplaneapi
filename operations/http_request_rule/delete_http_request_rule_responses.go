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

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteHTTPRequestRuleAcceptedCode is the HTTP code returned for type DeleteHTTPRequestRuleAccepted
const DeleteHTTPRequestRuleAcceptedCode int = 202

/*DeleteHTTPRequestRuleAccepted Configuration change accepted and reload requested

swagger:response deleteHttpRequestRuleAccepted
*/
type DeleteHTTPRequestRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteHTTPRequestRuleAccepted creates DeleteHTTPRequestRuleAccepted with default headers values
func NewDeleteHTTPRequestRuleAccepted() *DeleteHTTPRequestRuleAccepted {

	return &DeleteHTTPRequestRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete Http request rule accepted response
func (o *DeleteHTTPRequestRuleAccepted) WithReloadID(reloadID string) *DeleteHTTPRequestRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Http request rule accepted response
func (o *DeleteHTTPRequestRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteHTTPRequestRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteHTTPRequestRuleNoContentCode is the HTTP code returned for type DeleteHTTPRequestRuleNoContent
const DeleteHTTPRequestRuleNoContentCode int = 204

/*DeleteHTTPRequestRuleNoContent HTTP Request Rule deleted

swagger:response deleteHttpRequestRuleNoContent
*/
type DeleteHTTPRequestRuleNoContent struct {
}

// NewDeleteHTTPRequestRuleNoContent creates DeleteHTTPRequestRuleNoContent with default headers values
func NewDeleteHTTPRequestRuleNoContent() *DeleteHTTPRequestRuleNoContent {

	return &DeleteHTTPRequestRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteHTTPRequestRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteHTTPRequestRuleNotFoundCode is the HTTP code returned for type DeleteHTTPRequestRuleNotFound
const DeleteHTTPRequestRuleNotFoundCode int = 404

/*DeleteHTTPRequestRuleNotFound The specified resource was not found

swagger:response deleteHttpRequestRuleNotFound
*/
type DeleteHTTPRequestRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPRequestRuleNotFound creates DeleteHTTPRequestRuleNotFound with default headers values
func NewDeleteHTTPRequestRuleNotFound() *DeleteHTTPRequestRuleNotFound {

	return &DeleteHTTPRequestRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Http request rule not found response
func (o *DeleteHTTPRequestRuleNotFound) WithConfigurationVersion(configurationVersion string) *DeleteHTTPRequestRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Http request rule not found response
func (o *DeleteHTTPRequestRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Http request rule not found response
func (o *DeleteHTTPRequestRuleNotFound) WithPayload(payload *models.Error) *DeleteHTTPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Http request rule not found response
func (o *DeleteHTTPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteHTTPRequestRuleDefault General Error

swagger:response deleteHttpRequestRuleDefault
*/
type DeleteHTTPRequestRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPRequestRuleDefault creates DeleteHTTPRequestRuleDefault with default headers values
func NewDeleteHTTPRequestRuleDefault(code int) *DeleteHTTPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) WithStatusCode(code int) *DeleteHTTPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) WithConfigurationVersion(configurationVersion string) *DeleteHTTPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) WithPayload(payload *models.Error) *DeleteHTTPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
