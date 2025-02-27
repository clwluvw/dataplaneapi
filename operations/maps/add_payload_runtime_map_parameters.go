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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// NewAddPayloadRuntimeMapParams creates a new AddPayloadRuntimeMapParams object
// with the default values initialized.
func NewAddPayloadRuntimeMapParams() AddPayloadRuntimeMapParams {

	var (
		// initialize parameters with default values

		forceSyncDefault = bool(false)
	)

	return AddPayloadRuntimeMapParams{
		ForceSync: &forceSyncDefault,
	}
}

// AddPayloadRuntimeMapParams contains all the bound params for the add payload runtime map operation
// typically these are obtained from a http.Request
//
// swagger:parameters addPayloadRuntimeMap
type AddPayloadRuntimeMapParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Data models.MapEntries
	/*If true, immediately syncs changes to disk
	  In: query
	  Default: false
	*/
	ForceSync *bool
	/*Map file name
	  Required: true
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddPayloadRuntimeMapParams() beforehand.
func (o *AddPayloadRuntimeMapParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.MapEntries
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	qForceSync, qhkForceSync, _ := qs.GetOK("force_sync")
	if err := o.bindForceSync(qForceSync, qhkForceSync, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindForceSync binds and validates parameter ForceSync from query.
func (o *AddPayloadRuntimeMapParams) bindForceSync(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewAddPayloadRuntimeMapParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("force_sync", "query", "bool", raw)
	}
	o.ForceSync = &value

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *AddPayloadRuntimeMapParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Name = raw

	return nil
}
