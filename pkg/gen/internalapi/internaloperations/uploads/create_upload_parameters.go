// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CreateUploadMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var CreateUploadMaxParseMemory int64 = 32 << 20

// NewCreateUploadParams creates a new CreateUploadParams object
//
// There are no default values defined in the spec.
func NewCreateUploadParams() CreateUploadParams {

	return CreateUploadParams{}
}

// CreateUploadParams contains all the bound params for the create upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUpload
type CreateUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the document to add an upload to
	  In: query
	*/
	DocumentID *strfmt.UUID
	/*The file to upload.
	  Required: true
	  In: formData
	*/
	File io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateUploadParams() beforehand.
func (o *CreateUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := r.ParseMultipartForm(CreateUploadMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	qDocumentID, qhkDocumentID, _ := qs.GetOK("documentId")
	if err := o.bindDocumentID(qDocumentID, qhkDocumentID, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err := o.bindFile(file, fileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDocumentID binds and validates parameter DocumentID from query.
func (o *CreateUploadParams) bindDocumentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("documentId", "query", "strfmt.UUID", raw)
	}
	o.DocumentID = (value.(*strfmt.UUID))

	if err := o.validateDocumentID(formats); err != nil {
		return err
	}

	return nil
}

// validateDocumentID carries on validations for parameter DocumentID
func (o *CreateUploadParams) validateDocumentID(formats strfmt.Registry) error {

	if err := validate.FormatOf("documentId", "query", "uuid", o.DocumentID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *CreateUploadParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
