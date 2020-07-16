// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddProjectOKCode is the HTTP code returned for type AddProjectOK
const AddProjectOKCode int = 200

/*AddProjectOK OK

swagger:response addProjectOK
*/
type AddProjectOK struct {

	/*
	  In: Body
	*/
	Payload *AddProjectOKBody `json:"body,omitempty"`
}

// NewAddProjectOK creates AddProjectOK with default headers values
func NewAddProjectOK() *AddProjectOK {

	return &AddProjectOK{}
}

// WithPayload adds the payload to the add project o k response
func (o *AddProjectOK) WithPayload(payload *AddProjectOKBody) *AddProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add project o k response
func (o *AddProjectOK) SetPayload(payload *AddProjectOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddProjectBadRequestCode is the HTTP code returned for type AddProjectBadRequest
const AddProjectBadRequestCode int = 400

/*AddProjectBadRequest bad request

swagger:response addProjectBadRequest
*/
type AddProjectBadRequest struct {
}

// NewAddProjectBadRequest creates AddProjectBadRequest with default headers values
func NewAddProjectBadRequest() *AddProjectBadRequest {

	return &AddProjectBadRequest{}
}

// WriteResponse to the client
func (o *AddProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddProjectUnauthorizedCode is the HTTP code returned for type AddProjectUnauthorized
const AddProjectUnauthorizedCode int = 401

/*AddProjectUnauthorized Unauthorized

swagger:response addProjectUnauthorized
*/
type AddProjectUnauthorized struct {
}

// NewAddProjectUnauthorized creates AddProjectUnauthorized with default headers values
func NewAddProjectUnauthorized() *AddProjectUnauthorized {

	return &AddProjectUnauthorized{}
}

// WriteResponse to the client
func (o *AddProjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}