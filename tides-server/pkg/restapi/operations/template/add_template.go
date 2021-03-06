// Code generated by go-swagger; DO NOT EDIT.

package template

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

// AddTemplateHandlerFunc turns a function with the right signature into a add template handler
type AddTemplateHandlerFunc func(AddTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTemplateHandlerFunc) Handle(params AddTemplateParams) middleware.Responder {
	return fn(params)
}

// AddTemplateHandler interface for that can handle valid add template params
type AddTemplateHandler interface {
	Handle(AddTemplateParams) middleware.Responder
}

// NewAddTemplate creates a new http.Handler for the add template operation
func NewAddTemplate(ctx *middleware.Context, handler AddTemplateHandler) *AddTemplate {
	return &AddTemplate{Context: ctx, Handler: handler}
}

/*AddTemplate swagger:route POST /template/add template addTemplate

upload a VM template

*/
type AddTemplate struct {
	Context *middleware.Context
	Handler AddTemplateHandler
}

func (o *AddTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddTemplateBody add template body
//
// swagger:model AddTemplateBody
type AddTemplateBody struct {

	// compat
	Compat string `json:"compat,omitempty"`

	// memsize
	Memsize float64 `json:"memsize,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// os
	Os string `json:"os,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// space
	Space float64 `json:"space,omitempty"`
}

// Validate validates this add template body
func (o *AddTemplateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddTemplateBody) UnmarshalBinary(b []byte) error {
	var res AddTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddTemplateOKBody add template o k body
//
// swagger:model AddTemplateOKBody
type AddTemplateOKBody struct {

	// id
	ID int64 `json:"id,omitempty"`

	// message
	// Enum: [success]
	Message string `json:"message,omitempty"`
}

// Validate validates this add template o k body
func (o *AddTemplateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addTemplateOKBodyTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addTemplateOKBodyTypeMessagePropEnum = append(addTemplateOKBodyTypeMessagePropEnum, v)
	}
}

const (

	// AddTemplateOKBodyMessageSuccess captures enum value "success"
	AddTemplateOKBodyMessageSuccess string = "success"
)

// prop value enum
func (o *AddTemplateOKBody) validateMessageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addTemplateOKBodyTypeMessagePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddTemplateOKBody) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(o.Message) { // not required
		return nil
	}

	// value enum
	if err := o.validateMessageEnum("addTemplateOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddTemplateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddTemplateOKBody) UnmarshalBinary(b []byte) error {
	var res AddTemplateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
