// Code generated by go-swagger; DO NOT EDIT.

package vendor_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteVendorHandlerFunc turns a function with the right signature into a delete vendor handler
type DeleteVendorHandlerFunc func(DeleteVendorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteVendorHandlerFunc) Handle(params DeleteVendorParams) middleware.Responder {
	return fn(params)
}

// DeleteVendorHandler interface for that can handle valid delete vendor params
type DeleteVendorHandler interface {
	Handle(DeleteVendorParams) middleware.Responder
}

// NewDeleteVendor creates a new http.Handler for the delete vendor operation
func NewDeleteVendor(ctx *middleware.Context, handler DeleteVendorHandler) *DeleteVendor {
	return &DeleteVendor{Context: ctx, Handler: handler}
}

/* DeleteVendor swagger:route DELETE /vendors/{id} vendor deleteVendor

delete vendor

*/
type DeleteVendor struct {
	Context *middleware.Context
	Handler DeleteVendorHandler
}

func (o *DeleteVendor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteVendorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteVendorNotFoundBody delete vendor not found body
//
// swagger:model DeleteVendorNotFoundBody
type DeleteVendorNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete vendor not found body
func (o *DeleteVendorNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete vendor not found body based on context it is used
func (o *DeleteVendorNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteVendorNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteVendorNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteVendorNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteVendorOKBody delete vendor o k body
//
// swagger:model DeleteVendorOKBody
type DeleteVendorOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete vendor o k body
func (o *DeleteVendorOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete vendor o k body based on context it is used
func (o *DeleteVendorOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteVendorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteVendorOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteVendorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
