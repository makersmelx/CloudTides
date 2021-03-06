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

// DeleteTemplateHandlerFunc turns a function with the right signature into a delete template handler
type DeleteTemplateHandlerFunc func(DeleteTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTemplateHandlerFunc) Handle(params DeleteTemplateParams) middleware.Responder {
	return fn(params)
}

// DeleteTemplateHandler interface for that can handle valid delete template params
type DeleteTemplateHandler interface {
	Handle(DeleteTemplateParams) middleware.Responder
}

// NewDeleteTemplate creates a new http.Handler for the delete template operation
func NewDeleteTemplate(ctx *middleware.Context, handler DeleteTemplateHandler) *DeleteTemplate {
	return &DeleteTemplate{Context: ctx, Handler: handler}
}

/*DeleteTemplate swagger:route DELETE /template/delete template deleteTemplate

delete specified template

*/
type DeleteTemplate struct {
	Context *middleware.Context
	Handler DeleteTemplateHandler
}

func (o *DeleteTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteTemplateBody delete template body
//
// swagger:model DeleteTemplateBody
type DeleteTemplateBody struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this delete template body
func (o *DeleteTemplateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTemplateBody) UnmarshalBinary(b []byte) error {
	var res DeleteTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteTemplateOKBody delete template o k body
//
// swagger:model DeleteTemplateOKBody
type DeleteTemplateOKBody struct {

	// message
	// Enum: [success]
	Message string `json:"message,omitempty"`
}

// Validate validates this delete template o k body
func (o *DeleteTemplateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteTemplateOKBodyTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteTemplateOKBodyTypeMessagePropEnum = append(deleteTemplateOKBodyTypeMessagePropEnum, v)
	}
}

const (

	// DeleteTemplateOKBodyMessageSuccess captures enum value "success"
	DeleteTemplateOKBodyMessageSuccess string = "success"
)

// prop value enum
func (o *DeleteTemplateOKBody) validateMessageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deleteTemplateOKBodyTypeMessagePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *DeleteTemplateOKBody) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(o.Message) { // not required
		return nil
	}

	// value enum
	if err := o.validateMessageEnum("deleteTemplateOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteTemplateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTemplateOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteTemplateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
