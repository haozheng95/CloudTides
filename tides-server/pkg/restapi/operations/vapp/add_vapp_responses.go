// Code generated by go-swagger; DO NOT EDIT.

package vapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddVappOKCode is the HTTP code returned for type AddVappOK
const AddVappOKCode int = 200

/*AddVappOK OK

swagger:response addVappOK
*/
type AddVappOK struct {

	/*
	  In: Body
	*/
	Payload *AddVappOKBody `json:"body,omitempty"`
}

// NewAddVappOK creates AddVappOK with default headers values
func NewAddVappOK() *AddVappOK {

	return &AddVappOK{}
}

// WithPayload adds the payload to the add vapp o k response
func (o *AddVappOK) WithPayload(payload *AddVappOKBody) *AddVappOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add vapp o k response
func (o *AddVappOK) SetPayload(payload *AddVappOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddVappOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddVappBadRequestCode is the HTTP code returned for type AddVappBadRequest
const AddVappBadRequestCode int = 400

/*AddVappBadRequest bad request

swagger:response addVappBadRequest
*/
type AddVappBadRequest struct {
}

// NewAddVappBadRequest creates AddVappBadRequest with default headers values
func NewAddVappBadRequest() *AddVappBadRequest {

	return &AddVappBadRequest{}
}

// WriteResponse to the client
func (o *AddVappBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddVappUnauthorizedCode is the HTTP code returned for type AddVappUnauthorized
const AddVappUnauthorizedCode int = 401

/*AddVappUnauthorized Unauthorized

swagger:response addVappUnauthorized
*/
type AddVappUnauthorized struct {
}

// NewAddVappUnauthorized creates AddVappUnauthorized with default headers values
func NewAddVappUnauthorized() *AddVappUnauthorized {

	return &AddVappUnauthorized{}
}

// WriteResponse to the client
func (o *AddVappUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AddVappNotFoundCode is the HTTP code returned for type AddVappNotFound
const AddVappNotFoundCode int = 404

/*AddVappNotFound resource not found

swagger:response addVappNotFound
*/
type AddVappNotFound struct {

	/*
	  In: Body
	*/
	Payload *AddVappNotFoundBody `json:"body,omitempty"`
}

// NewAddVappNotFound creates AddVappNotFound with default headers values
func NewAddVappNotFound() *AddVappNotFound {

	return &AddVappNotFound{}
}

// WithPayload adds the payload to the add vapp not found response
func (o *AddVappNotFound) WithPayload(payload *AddVappNotFoundBody) *AddVappNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add vapp not found response
func (o *AddVappNotFound) SetPayload(payload *AddVappNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddVappNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}