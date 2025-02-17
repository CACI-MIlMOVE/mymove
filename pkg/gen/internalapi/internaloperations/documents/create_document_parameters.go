// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// NewCreateDocumentParams creates a new CreateDocumentParams object
//
// There are no default values defined in the spec.
func NewCreateDocumentParams() CreateDocumentParams {

	return CreateDocumentParams{}
}

// CreateDocumentParams contains all the bound params for the create document operation
// typically these are obtained from a http.Request
//
// swagger:parameters createDocument
type CreateDocumentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	DocumentPayload *internalmessages.PostDocumentPayload
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateDocumentParams() beforehand.
func (o *CreateDocumentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body internalmessages.PostDocumentPayload
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("documentPayload", "body", ""))
			} else {
				res = append(res, errors.NewParseError("documentPayload", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.DocumentPayload = &body
			}
		}
	} else {
		res = append(res, errors.Required("documentPayload", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
