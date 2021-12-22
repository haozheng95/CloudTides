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

// AchieveHostHandlerFunc turns a function with the right signature into a achieve host handler
type AchieveHostHandlerFunc func(AchieveHostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AchieveHostHandlerFunc) Handle(params AchieveHostParams) middleware.Responder {
	return fn(params)
}

// AchieveHostHandler interface for that can handle valid achieve host params
type AchieveHostHandler interface {
	Handle(AchieveHostParams) middleware.Responder
}

// NewAchieveHost creates a new http.Handler for the achieve host operation
func NewAchieveHost(ctx *middleware.Context, handler AchieveHostHandler) *AchieveHost {
	return &AchieveHost{Context: ctx, Handler: handler}
}

/* AchieveHost swagger:route GET /application/instance/hosts application achieveHost

achieve host

*/
type AchieveHost struct {
	Context *middleware.Context
	Handler AchieveHostHandler
}

func (o *AchieveHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAchieveHostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AchieveHostOKBodyItems0 achieve host o k body items0
//
// swagger:model AchieveHostOKBodyItems0
type AchieveHostOKBodyItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// ssh pass
	SSHPass string `json:"sshPass,omitempty"`

	// ssh port
	SSHPort string `json:"sshPort,omitempty"`

	// ssh user
	SSHUser string `json:"sshUser,omitempty"`
}

// Validate validates this achieve host o k body items0
func (o *AchieveHostOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this achieve host o k body items0 based on context it is used
func (o *AchieveHostOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AchieveHostOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AchieveHostOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res AchieveHostOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}