// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HostPastUsageOKCode is the HTTP code returned for type HostPastUsageOK
const HostPastUsageOKCode int = 200

/*HostPastUsageOK OK

swagger:response hostPastUsageOK
*/
type HostPastUsageOK struct {
}

// NewHostPastUsageOK creates HostPastUsageOK with default headers values
func NewHostPastUsageOK() *HostPastUsageOK {

	return &HostPastUsageOK{}
}

// WriteResponse to the client
func (o *HostPastUsageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// HostPastUsageBadRequestCode is the HTTP code returned for type HostPastUsageBadRequest
const HostPastUsageBadRequestCode int = 400

/*HostPastUsageBadRequest bad request

swagger:response hostPastUsageBadRequest
*/
type HostPastUsageBadRequest struct {
}

// NewHostPastUsageBadRequest creates HostPastUsageBadRequest with default headers values
func NewHostPastUsageBadRequest() *HostPastUsageBadRequest {

	return &HostPastUsageBadRequest{}
}

// WriteResponse to the client
func (o *HostPastUsageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}