// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserLoginHandlerFunc turns a function with the right signature into a user login handler
type UserLoginHandlerFunc func(UserLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserLoginHandlerFunc) Handle(params UserLoginParams) middleware.Responder {
	return fn(params)
}

// UserLoginHandler interface for that can handle valid user login params
type UserLoginHandler interface {
	Handle(UserLoginParams) middleware.Responder
}

// NewUserLogin creates a new http.Handler for the user login operation
func NewUserLogin(ctx *middleware.Context, handler UserLoginHandler) *UserLogin {
	return &UserLogin{Context: ctx, Handler: handler}
}

/*UserLogin swagger:route POST /users/login user userLogin

user login

*/
type UserLogin struct {
	Context *middleware.Context
	Handler UserLoginHandler
}

func (o *UserLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UserLoginBody user login body
//
// swagger:model UserLoginBody
type UserLoginBody struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user login body
func (o *UserLoginBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UserLoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserLoginBody) UnmarshalBinary(b []byte) error {
	var res UserLoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UserLoginOKBody user login o k body
//
// swagger:model UserLoginOKBody
type UserLoginOKBody struct {

	// token
	Token string `json:"token,omitempty"`

	// user info
	UserInfo *UserLoginOKBodyUserInfo `json:"userInfo,omitempty"`
}

// Validate validates this user login o k body
func (o *UserLoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUserInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserLoginOKBody) validateUserInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.UserInfo) { // not required
		return nil
	}

	if o.UserInfo != nil {
		if err := o.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userLoginOK" + "." + "userInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserLoginOKBody) UnmarshalBinary(b []byte) error {
	var res UserLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UserLoginOKBodyUserInfo user login o k body user info
//
// swagger:model UserLoginOKBodyUserInfo
type UserLoginOKBodyUserInfo struct {

	// priority
	// Enum: [Low Medium High]
	Priority string `json:"priority,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user login o k body user info
func (o *UserLoginOKBodyUserInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userLoginOKBodyUserInfoTypePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Low","Medium","High"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userLoginOKBodyUserInfoTypePriorityPropEnum = append(userLoginOKBodyUserInfoTypePriorityPropEnum, v)
	}
}

const (

	// UserLoginOKBodyUserInfoPriorityLow captures enum value "Low"
	UserLoginOKBodyUserInfoPriorityLow string = "Low"

	// UserLoginOKBodyUserInfoPriorityMedium captures enum value "Medium"
	UserLoginOKBodyUserInfoPriorityMedium string = "Medium"

	// UserLoginOKBodyUserInfoPriorityHigh captures enum value "High"
	UserLoginOKBodyUserInfoPriorityHigh string = "High"
)

// prop value enum
func (o *UserLoginOKBodyUserInfo) validatePriorityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userLoginOKBodyUserInfoTypePriorityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *UserLoginOKBodyUserInfo) validatePriority(formats strfmt.Registry) error {

	if swag.IsZero(o.Priority) { // not required
		return nil
	}

	// value enum
	if err := o.validatePriorityEnum("userLoginOK"+"."+"userInfo"+"."+"priority", "body", o.Priority); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserLoginOKBodyUserInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserLoginOKBodyUserInfo) UnmarshalBinary(b []byte) error {
	var res UserLoginOKBodyUserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UserLoginUnauthorizedBody user login unauthorized body
//
// swagger:model UserLoginUnauthorizedBody
type UserLoginUnauthorizedBody struct {

	// message
	// Enum: [Unauthorized]
	Message string `json:"message,omitempty"`
}

// Validate validates this user login unauthorized body
func (o *UserLoginUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userLoginUnauthorizedBodyTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unauthorized"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userLoginUnauthorizedBodyTypeMessagePropEnum = append(userLoginUnauthorizedBodyTypeMessagePropEnum, v)
	}
}

const (

	// UserLoginUnauthorizedBodyMessageUnauthorized captures enum value "Unauthorized"
	UserLoginUnauthorizedBodyMessageUnauthorized string = "Unauthorized"
)

// prop value enum
func (o *UserLoginUnauthorizedBody) validateMessageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userLoginUnauthorizedBodyTypeMessagePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *UserLoginUnauthorizedBody) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(o.Message) { // not required
		return nil
	}

	// value enum
	if err := o.validateMessageEnum("userLoginUnauthorized"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserLoginUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserLoginUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res UserLoginUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}