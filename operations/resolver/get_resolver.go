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

package resolver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// GetResolverHandlerFunc turns a function with the right signature into a get resolver handler
type GetResolverHandlerFunc func(GetResolverParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResolverHandlerFunc) Handle(params GetResolverParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetResolverHandler interface for that can handle valid get resolver params
type GetResolverHandler interface {
	Handle(GetResolverParams, interface{}) middleware.Responder
}

// NewGetResolver creates a new http.Handler for the get resolver operation
func NewGetResolver(ctx *middleware.Context, handler GetResolverHandler) *GetResolver {
	return &GetResolver{Context: ctx, Handler: handler}
}

/*GetResolver swagger:route GET /services/haproxy/configuration/resolvers/{name} Resolver getResolver

Return a resolver

Returns one resolver section configuration by it's name.

*/
type GetResolver struct {
	Context *middleware.Context
	Handler GetResolverHandler
}

func (o *GetResolver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResolverParams()

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

// GetResolverOKBody get resolver o k body
//
// swagger:model GetResolverOKBody
type GetResolverOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Resolver `json:"data,omitempty"`
}

// Validate validates this get resolver o k body
func (o *GetResolverOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResolverOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResolverOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResolverOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResolverOKBody) UnmarshalBinary(b []byte) error {
	var res GetResolverOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
