// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListApplicationInstanceHandlerFunc turns a function with the right signature into a list application instance handler
type ListApplicationInstanceHandlerFunc func(ListApplicationInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListApplicationInstanceHandlerFunc) Handle(params ListApplicationInstanceParams) middleware.Responder {
	return fn(params)
}

// ListApplicationInstanceHandler interface for that can handle valid list application instance params
type ListApplicationInstanceHandler interface {
	Handle(ListApplicationInstanceParams) middleware.Responder
}

// NewListApplicationInstance creates a new http.Handler for the list application instance operation
func NewListApplicationInstance(ctx *middleware.Context, handler ListApplicationInstanceHandler) *ListApplicationInstance {
	return &ListApplicationInstance{Context: ctx, Handler: handler}
}

/* ListApplicationInstance swagger:route GET /application/instance application listApplicationInstance

list application instance

*/
type ListApplicationInstance struct {
	Context *middleware.Context
	Handler ListApplicationInstanceHandler
}

func (o *ListApplicationInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListApplicationInstanceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListApplicationInstanceOKBodyItems0 list application instance o k body items0
//
// swagger:model ListApplicationInstanceOKBodyItems0
type ListApplicationInstanceOKBodyItems0 struct {

	// instance name
	InstanceName string `json:"instanceName,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this list application instance o k body items0
func (o *ListApplicationInstanceOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list application instance o k body items0 based on context it is used
func (o *ListApplicationInstanceOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListApplicationInstanceOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListApplicationInstanceOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ListApplicationInstanceOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
