// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListResourceOKCode is the HTTP code returned for type ListResourceOK
const ListResourceOKCode int = 200

/*ListResourceOK returns the list of resources belonging to a user

swagger:response listResourceOK
*/
type ListResourceOK struct {

	/*
	  In: Body
	*/
	Payload *ListResourceOKBody `json:"body,omitempty"`
}

// NewListResourceOK creates ListResourceOK with default headers values
func NewListResourceOK() *ListResourceOK {

	return &ListResourceOK{}
}

// WithPayload adds the payload to the list resource o k response
func (o *ListResourceOK) WithPayload(payload *ListResourceOKBody) *ListResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list resource o k response
func (o *ListResourceOK) SetPayload(payload *ListResourceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListResourceUnauthorizedCode is the HTTP code returned for type ListResourceUnauthorized
const ListResourceUnauthorizedCode int = 401

/*ListResourceUnauthorized Unauthorized

swagger:response listResourceUnauthorized
*/
type ListResourceUnauthorized struct {
}

// NewListResourceUnauthorized creates ListResourceUnauthorized with default headers values
func NewListResourceUnauthorized() *ListResourceUnauthorized {

	return &ListResourceUnauthorized{}
}

// WriteResponse to the client
func (o *ListResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListResourceNotFoundCode is the HTTP code returned for type ListResourceNotFound
const ListResourceNotFoundCode int = 404

/*ListResourceNotFound resource not found

swagger:response listResourceNotFound
*/
type ListResourceNotFound struct {
}

// NewListResourceNotFound creates ListResourceNotFound with default headers values
func NewListResourceNotFound() *ListResourceNotFound {

	return &ListResourceNotFound{}
}

// WriteResponse to the client
func (o *ListResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}