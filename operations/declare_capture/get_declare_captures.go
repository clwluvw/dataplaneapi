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

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v3/models"
)

// GetDeclareCapturesHandlerFunc turns a function with the right signature into a get declare captures handler
type GetDeclareCapturesHandlerFunc func(GetDeclareCapturesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDeclareCapturesHandlerFunc) Handle(params GetDeclareCapturesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDeclareCapturesHandler interface for that can handle valid get declare captures params
type GetDeclareCapturesHandler interface {
	Handle(GetDeclareCapturesParams, interface{}) middleware.Responder
}

// NewGetDeclareCaptures creates a new http.Handler for the get declare captures operation
func NewGetDeclareCaptures(ctx *middleware.Context, handler GetDeclareCapturesHandler) *GetDeclareCaptures {
	return &GetDeclareCaptures{Context: ctx, Handler: handler}
}

/*GetDeclareCaptures swagger:route GET /services/haproxy/configuration/captures DeclareCapture getDeclareCaptures

Return an array of declare captures

Returns an array of all declare capture records that are configured in specified frontend.

*/
type GetDeclareCaptures struct {
	Context *middleware.Context
	Handler GetDeclareCapturesHandler
}

func (o *GetDeclareCaptures) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDeclareCapturesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetDeclareCapturesOKBody get declare captures o k body
//
// swagger:model GetDeclareCapturesOKBody
type GetDeclareCapturesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Captures `json:"data"`
}

// Validate validates this get declare captures o k body
func (o *GetDeclareCapturesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeclareCapturesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getDeclareCapturesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getDeclareCapturesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeclareCapturesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeclareCapturesOKBody) UnmarshalBinary(b []byte) error {
	var res GetDeclareCapturesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
