// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListPolicyOKCode is the HTTP code returned for type ListPolicyOK
const ListPolicyOKCode int = 200

/*ListPolicyOK OK

swagger:response listPolicyOK
*/
type ListPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *ListPolicyOKBody `json:"body,omitempty"`
}

// NewListPolicyOK creates ListPolicyOK with default headers values
func NewListPolicyOK() *ListPolicyOK {

	return &ListPolicyOK{}
}

// WithPayload adds the payload to the list policy o k response
func (o *ListPolicyOK) WithPayload(payload *ListPolicyOKBody) *ListPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list policy o k response
func (o *ListPolicyOK) SetPayload(payload *ListPolicyOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPolicyBadRequestCode is the HTTP code returned for type ListPolicyBadRequest
const ListPolicyBadRequestCode int = 400

/*ListPolicyBadRequest bad request

swagger:response listPolicyBadRequest
*/
type ListPolicyBadRequest struct {
}

// NewListPolicyBadRequest creates ListPolicyBadRequest with default headers values
func NewListPolicyBadRequest() *ListPolicyBadRequest {

	return &ListPolicyBadRequest{}
}

// WriteResponse to the client
func (o *ListPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ListPolicyUnauthorizedCode is the HTTP code returned for type ListPolicyUnauthorized
const ListPolicyUnauthorizedCode int = 401

/*ListPolicyUnauthorized Unauthorized

swagger:response listPolicyUnauthorized
*/
type ListPolicyUnauthorized struct {
}

// NewListPolicyUnauthorized creates ListPolicyUnauthorized with default headers values
func NewListPolicyUnauthorized() *ListPolicyUnauthorized {

	return &ListPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *ListPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListPolicyNotFoundCode is the HTTP code returned for type ListPolicyNotFound
const ListPolicyNotFoundCode int = 404

/*ListPolicyNotFound resource not found

swagger:response listPolicyNotFound
*/
type ListPolicyNotFound struct {
}

// NewListPolicyNotFound creates ListPolicyNotFound with default headers values
func NewListPolicyNotFound() *ListPolicyNotFound {

	return &ListPolicyNotFound{}
}

// WriteResponse to the client
func (o *ListPolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}