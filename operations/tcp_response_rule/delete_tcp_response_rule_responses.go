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

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteTCPResponseRuleAcceptedCode is the HTTP code returned for type DeleteTCPResponseRuleAccepted
const DeleteTCPResponseRuleAcceptedCode int = 202

/*DeleteTCPResponseRuleAccepted Configuration change accepted and reload requested

swagger:response deleteTcpResponseRuleAccepted
*/
type DeleteTCPResponseRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteTCPResponseRuleAccepted creates DeleteTCPResponseRuleAccepted with default headers values
func NewDeleteTCPResponseRuleAccepted() *DeleteTCPResponseRuleAccepted {

	return &DeleteTCPResponseRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete Tcp response rule accepted response
func (o *DeleteTCPResponseRuleAccepted) WithReloadID(reloadID string) *DeleteTCPResponseRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Tcp response rule accepted response
func (o *DeleteTCPResponseRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteTCPResponseRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteTCPResponseRuleNoContentCode is the HTTP code returned for type DeleteTCPResponseRuleNoContent
const DeleteTCPResponseRuleNoContentCode int = 204

/*DeleteTCPResponseRuleNoContent TCP Response Rule deleted

swagger:response deleteTcpResponseRuleNoContent
*/
type DeleteTCPResponseRuleNoContent struct {
}

// NewDeleteTCPResponseRuleNoContent creates DeleteTCPResponseRuleNoContent with default headers values
func NewDeleteTCPResponseRuleNoContent() *DeleteTCPResponseRuleNoContent {

	return &DeleteTCPResponseRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteTCPResponseRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTCPResponseRuleNotFoundCode is the HTTP code returned for type DeleteTCPResponseRuleNotFound
const DeleteTCPResponseRuleNotFoundCode int = 404

/*DeleteTCPResponseRuleNotFound The specified resource was not found

swagger:response deleteTcpResponseRuleNotFound
*/
type DeleteTCPResponseRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTCPResponseRuleNotFound creates DeleteTCPResponseRuleNotFound with default headers values
func NewDeleteTCPResponseRuleNotFound() *DeleteTCPResponseRuleNotFound {

	return &DeleteTCPResponseRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Tcp response rule not found response
func (o *DeleteTCPResponseRuleNotFound) WithConfigurationVersion(configurationVersion string) *DeleteTCPResponseRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Tcp response rule not found response
func (o *DeleteTCPResponseRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Tcp response rule not found response
func (o *DeleteTCPResponseRuleNotFound) WithPayload(payload *models.Error) *DeleteTCPResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Tcp response rule not found response
func (o *DeleteTCPResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTCPResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteTCPResponseRuleDefault General Error

swagger:response deleteTcpResponseRuleDefault
*/
type DeleteTCPResponseRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTCPResponseRuleDefault creates DeleteTCPResponseRuleDefault with default headers values
func NewDeleteTCPResponseRuleDefault(code int) *DeleteTCPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTCPResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) WithStatusCode(code int) *DeleteTCPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) WithConfigurationVersion(configurationVersion string) *DeleteTCPResponseRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) WithPayload(payload *models.Error) *DeleteTCPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTCPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
