// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// CreateProGearWeightTicketCreatedCode is the HTTP code returned for type CreateProGearWeightTicketCreated
const CreateProGearWeightTicketCreatedCode int = 201

/*
CreateProGearWeightTicketCreated returns a new pro-gear weight ticket object

swagger:response createProGearWeightTicketCreated
*/
type CreateProGearWeightTicketCreated struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ProGearWeightTicket `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketCreated creates CreateProGearWeightTicketCreated with default headers values
func NewCreateProGearWeightTicketCreated() *CreateProGearWeightTicketCreated {

	return &CreateProGearWeightTicketCreated{}
}

// WithPayload adds the payload to the create pro gear weight ticket created response
func (o *CreateProGearWeightTicketCreated) WithPayload(payload *internalmessages.ProGearWeightTicket) *CreateProGearWeightTicketCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket created response
func (o *CreateProGearWeightTicketCreated) SetPayload(payload *internalmessages.ProGearWeightTicket) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketBadRequestCode is the HTTP code returned for type CreateProGearWeightTicketBadRequest
const CreateProGearWeightTicketBadRequestCode int = 400

/*
CreateProGearWeightTicketBadRequest The request payload is invalid.

swagger:response createProGearWeightTicketBadRequest
*/
type CreateProGearWeightTicketBadRequest struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ClientError `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketBadRequest creates CreateProGearWeightTicketBadRequest with default headers values
func NewCreateProGearWeightTicketBadRequest() *CreateProGearWeightTicketBadRequest {

	return &CreateProGearWeightTicketBadRequest{}
}

// WithPayload adds the payload to the create pro gear weight ticket bad request response
func (o *CreateProGearWeightTicketBadRequest) WithPayload(payload *internalmessages.ClientError) *CreateProGearWeightTicketBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket bad request response
func (o *CreateProGearWeightTicketBadRequest) SetPayload(payload *internalmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketUnauthorizedCode is the HTTP code returned for type CreateProGearWeightTicketUnauthorized
const CreateProGearWeightTicketUnauthorizedCode int = 401

/*
CreateProGearWeightTicketUnauthorized The request was denied.

swagger:response createProGearWeightTicketUnauthorized
*/
type CreateProGearWeightTicketUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ClientError `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketUnauthorized creates CreateProGearWeightTicketUnauthorized with default headers values
func NewCreateProGearWeightTicketUnauthorized() *CreateProGearWeightTicketUnauthorized {

	return &CreateProGearWeightTicketUnauthorized{}
}

// WithPayload adds the payload to the create pro gear weight ticket unauthorized response
func (o *CreateProGearWeightTicketUnauthorized) WithPayload(payload *internalmessages.ClientError) *CreateProGearWeightTicketUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket unauthorized response
func (o *CreateProGearWeightTicketUnauthorized) SetPayload(payload *internalmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketForbiddenCode is the HTTP code returned for type CreateProGearWeightTicketForbidden
const CreateProGearWeightTicketForbiddenCode int = 403

/*
CreateProGearWeightTicketForbidden The request was denied.

swagger:response createProGearWeightTicketForbidden
*/
type CreateProGearWeightTicketForbidden struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ClientError `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketForbidden creates CreateProGearWeightTicketForbidden with default headers values
func NewCreateProGearWeightTicketForbidden() *CreateProGearWeightTicketForbidden {

	return &CreateProGearWeightTicketForbidden{}
}

// WithPayload adds the payload to the create pro gear weight ticket forbidden response
func (o *CreateProGearWeightTicketForbidden) WithPayload(payload *internalmessages.ClientError) *CreateProGearWeightTicketForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket forbidden response
func (o *CreateProGearWeightTicketForbidden) SetPayload(payload *internalmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketNotFoundCode is the HTTP code returned for type CreateProGearWeightTicketNotFound
const CreateProGearWeightTicketNotFoundCode int = 404

/*
CreateProGearWeightTicketNotFound The requested resource wasn't found.

swagger:response createProGearWeightTicketNotFound
*/
type CreateProGearWeightTicketNotFound struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ClientError `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketNotFound creates CreateProGearWeightTicketNotFound with default headers values
func NewCreateProGearWeightTicketNotFound() *CreateProGearWeightTicketNotFound {

	return &CreateProGearWeightTicketNotFound{}
}

// WithPayload adds the payload to the create pro gear weight ticket not found response
func (o *CreateProGearWeightTicketNotFound) WithPayload(payload *internalmessages.ClientError) *CreateProGearWeightTicketNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket not found response
func (o *CreateProGearWeightTicketNotFound) SetPayload(payload *internalmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketPreconditionFailedCode is the HTTP code returned for type CreateProGearWeightTicketPreconditionFailed
const CreateProGearWeightTicketPreconditionFailedCode int = 412

/*
CreateProGearWeightTicketPreconditionFailed Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.

swagger:response createProGearWeightTicketPreconditionFailed
*/
type CreateProGearWeightTicketPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ClientError `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketPreconditionFailed creates CreateProGearWeightTicketPreconditionFailed with default headers values
func NewCreateProGearWeightTicketPreconditionFailed() *CreateProGearWeightTicketPreconditionFailed {

	return &CreateProGearWeightTicketPreconditionFailed{}
}

// WithPayload adds the payload to the create pro gear weight ticket precondition failed response
func (o *CreateProGearWeightTicketPreconditionFailed) WithPayload(payload *internalmessages.ClientError) *CreateProGearWeightTicketPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket precondition failed response
func (o *CreateProGearWeightTicketPreconditionFailed) SetPayload(payload *internalmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketUnprocessableEntityCode is the HTTP code returned for type CreateProGearWeightTicketUnprocessableEntity
const CreateProGearWeightTicketUnprocessableEntityCode int = 422

/*
CreateProGearWeightTicketUnprocessableEntity The payload was unprocessable.

swagger:response createProGearWeightTicketUnprocessableEntity
*/
type CreateProGearWeightTicketUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ValidationError `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketUnprocessableEntity creates CreateProGearWeightTicketUnprocessableEntity with default headers values
func NewCreateProGearWeightTicketUnprocessableEntity() *CreateProGearWeightTicketUnprocessableEntity {

	return &CreateProGearWeightTicketUnprocessableEntity{}
}

// WithPayload adds the payload to the create pro gear weight ticket unprocessable entity response
func (o *CreateProGearWeightTicketUnprocessableEntity) WithPayload(payload *internalmessages.ValidationError) *CreateProGearWeightTicketUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket unprocessable entity response
func (o *CreateProGearWeightTicketUnprocessableEntity) SetPayload(payload *internalmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProGearWeightTicketInternalServerErrorCode is the HTTP code returned for type CreateProGearWeightTicketInternalServerError
const CreateProGearWeightTicketInternalServerErrorCode int = 500

/*
CreateProGearWeightTicketInternalServerError A server error occurred.

swagger:response createProGearWeightTicketInternalServerError
*/
type CreateProGearWeightTicketInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.Error `json:"body,omitempty"`
}

// NewCreateProGearWeightTicketInternalServerError creates CreateProGearWeightTicketInternalServerError with default headers values
func NewCreateProGearWeightTicketInternalServerError() *CreateProGearWeightTicketInternalServerError {

	return &CreateProGearWeightTicketInternalServerError{}
}

// WithPayload adds the payload to the create pro gear weight ticket internal server error response
func (o *CreateProGearWeightTicketInternalServerError) WithPayload(payload *internalmessages.Error) *CreateProGearWeightTicketInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pro gear weight ticket internal server error response
func (o *CreateProGearWeightTicketInternalServerError) SetPayload(payload *internalmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProGearWeightTicketInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
