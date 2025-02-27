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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetTransactionsOKCode is the HTTP code returned for type GetTransactionsOK
const GetTransactionsOKCode int = 200

/*GetTransactionsOK Success

swagger:response getTransactionsOK
*/
type GetTransactionsOK struct {

	/*
	  In: Body
	*/
	Payload models.Transactions `json:"body,omitempty"`
}

// NewGetTransactionsOK creates GetTransactionsOK with default headers values
func NewGetTransactionsOK() *GetTransactionsOK {

	return &GetTransactionsOK{}
}

// WithPayload adds the payload to the get transactions o k response
func (o *GetTransactionsOK) WithPayload(payload models.Transactions) *GetTransactionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions o k response
func (o *GetTransactionsOK) SetPayload(payload models.Transactions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Transactions{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetTransactionsDefault General Error

swagger:response getTransactionsDefault
*/
type GetTransactionsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionsDefault creates GetTransactionsDefault with default headers values
func NewGetTransactionsDefault(code int) *GetTransactionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTransactionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get transactions default response
func (o *GetTransactionsDefault) WithStatusCode(code int) *GetTransactionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get transactions default response
func (o *GetTransactionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get transactions default response
func (o *GetTransactionsDefault) WithConfigurationVersion(configurationVersion string) *GetTransactionsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get transactions default response
func (o *GetTransactionsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get transactions default response
func (o *GetTransactionsDefault) WithPayload(payload *models.Error) *GetTransactionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions default response
func (o *GetTransactionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
