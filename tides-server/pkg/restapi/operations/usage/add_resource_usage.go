// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddResourceUsageHandlerFunc turns a function with the right signature into a add resource usage handler
type AddResourceUsageHandlerFunc func(AddResourceUsageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddResourceUsageHandlerFunc) Handle(params AddResourceUsageParams) middleware.Responder {
	return fn(params)
}

// AddResourceUsageHandler interface for that can handle valid add resource usage params
type AddResourceUsageHandler interface {
	Handle(AddResourceUsageParams) middleware.Responder
}

// NewAddResourceUsage creates a new http.Handler for the add resource usage operation
func NewAddResourceUsage(ctx *middleware.Context, handler AddResourceUsageHandler) *AddResourceUsage {
	return &AddResourceUsage{Context: ctx, Handler: handler}
}

/* AddResourceUsage swagger:route POST /usage usage addResourceUsage

add resource usage info into database

*/
type AddResourceUsage struct {
	Context *middleware.Context
	Handler AddResourceUsageHandler
}

func (o *AddResourceUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddResourceUsageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddResourceUsageBody add resource usage body
//
// swagger:model AddResourceUsageBody
type AddResourceUsageBody struct {

	// current CPU
	CurrentCPU float64 `json:"currentCPU,omitempty"`

	// current disk
	CurrentDisk float64 `json:"currentDisk,omitempty"`

	// current RAM
	CurrentRAM float64 `json:"currentRAM,omitempty"`

	// host address
	HostAddress string `json:"hostAddress,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// total CPU
	TotalCPU float64 `json:"totalCPU,omitempty"`

	// total disk
	TotalDisk float64 `json:"totalDisk,omitempty"`

	// total RAM
	TotalRAM float64 `json:"totalRAM,omitempty"`
}

// Validate validates this add resource usage body
func (o *AddResourceUsageBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add resource usage body based on context it is used
func (o *AddResourceUsageBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddResourceUsageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddResourceUsageBody) UnmarshalBinary(b []byte) error {
	var res AddResourceUsageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddResourceUsageOKBody add resource usage o k body
//
// swagger:model AddResourceUsageOKBody
type AddResourceUsageOKBody struct {

	// message
	// Enum: [success]
	Message string `json:"message,omitempty"`
}

// Validate validates this add resource usage o k body
func (o *AddResourceUsageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addResourceUsageOKBodyTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addResourceUsageOKBodyTypeMessagePropEnum = append(addResourceUsageOKBodyTypeMessagePropEnum, v)
	}
}

const (

	// AddResourceUsageOKBodyMessageSuccess captures enum value "success"
	AddResourceUsageOKBodyMessageSuccess string = "success"
)

// prop value enum
func (o *AddResourceUsageOKBody) validateMessageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addResourceUsageOKBodyTypeMessagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddResourceUsageOKBody) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(o.Message) { // not required
		return nil
	}

	// value enum
	if err := o.validateMessageEnum("addResourceUsageOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add resource usage o k body based on context it is used
func (o *AddResourceUsageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddResourceUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddResourceUsageOKBody) UnmarshalBinary(b []byte) error {
	var res AddResourceUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
