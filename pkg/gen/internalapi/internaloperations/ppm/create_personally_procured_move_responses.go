// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// CreatePersonallyProcuredMoveCreatedCode is the HTTP code returned for type CreatePersonallyProcuredMoveCreated
const CreatePersonallyProcuredMoveCreatedCode int = 201

/*
CreatePersonallyProcuredMoveCreated created instance of personally_procured_move

swagger:response createPersonallyProcuredMoveCreated
*/
type CreatePersonallyProcuredMoveCreated struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PersonallyProcuredMovePayload `json:"body,omitempty"`
}

// NewCreatePersonallyProcuredMoveCreated creates CreatePersonallyProcuredMoveCreated with default headers values
func NewCreatePersonallyProcuredMoveCreated() *CreatePersonallyProcuredMoveCreated {

	return &CreatePersonallyProcuredMoveCreated{}
}

// WithPayload adds the payload to the create personally procured move created response
func (o *CreatePersonallyProcuredMoveCreated) WithPayload(payload *internalmessages.PersonallyProcuredMovePayload) *CreatePersonallyProcuredMoveCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create personally procured move created response
func (o *CreatePersonallyProcuredMoveCreated) SetPayload(payload *internalmessages.PersonallyProcuredMovePayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePersonallyProcuredMoveCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePersonallyProcuredMoveBadRequestCode is the HTTP code returned for type CreatePersonallyProcuredMoveBadRequest
const CreatePersonallyProcuredMoveBadRequestCode int = 400

/*
CreatePersonallyProcuredMoveBadRequest invalid request

swagger:response createPersonallyProcuredMoveBadRequest
*/
type CreatePersonallyProcuredMoveBadRequest struct {
}

// NewCreatePersonallyProcuredMoveBadRequest creates CreatePersonallyProcuredMoveBadRequest with default headers values
func NewCreatePersonallyProcuredMoveBadRequest() *CreatePersonallyProcuredMoveBadRequest {

	return &CreatePersonallyProcuredMoveBadRequest{}
}

// WriteResponse to the client
func (o *CreatePersonallyProcuredMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreatePersonallyProcuredMoveUnauthorizedCode is the HTTP code returned for type CreatePersonallyProcuredMoveUnauthorized
const CreatePersonallyProcuredMoveUnauthorizedCode int = 401

/*
CreatePersonallyProcuredMoveUnauthorized request requires user authentication

swagger:response createPersonallyProcuredMoveUnauthorized
*/
type CreatePersonallyProcuredMoveUnauthorized struct {
}

// NewCreatePersonallyProcuredMoveUnauthorized creates CreatePersonallyProcuredMoveUnauthorized with default headers values
func NewCreatePersonallyProcuredMoveUnauthorized() *CreatePersonallyProcuredMoveUnauthorized {

	return &CreatePersonallyProcuredMoveUnauthorized{}
}

// WriteResponse to the client
func (o *CreatePersonallyProcuredMoveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// CreatePersonallyProcuredMoveForbiddenCode is the HTTP code returned for type CreatePersonallyProcuredMoveForbidden
const CreatePersonallyProcuredMoveForbiddenCode int = 403

/*
CreatePersonallyProcuredMoveForbidden user is not authorized

swagger:response createPersonallyProcuredMoveForbidden
*/
type CreatePersonallyProcuredMoveForbidden struct {
}

// NewCreatePersonallyProcuredMoveForbidden creates CreatePersonallyProcuredMoveForbidden with default headers values
func NewCreatePersonallyProcuredMoveForbidden() *CreatePersonallyProcuredMoveForbidden {

	return &CreatePersonallyProcuredMoveForbidden{}
}

// WriteResponse to the client
func (o *CreatePersonallyProcuredMoveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// CreatePersonallyProcuredMoveNotFoundCode is the HTTP code returned for type CreatePersonallyProcuredMoveNotFound
const CreatePersonallyProcuredMoveNotFoundCode int = 404

/*
CreatePersonallyProcuredMoveNotFound move not found

swagger:response createPersonallyProcuredMoveNotFound
*/
type CreatePersonallyProcuredMoveNotFound struct {
}

// NewCreatePersonallyProcuredMoveNotFound creates CreatePersonallyProcuredMoveNotFound with default headers values
func NewCreatePersonallyProcuredMoveNotFound() *CreatePersonallyProcuredMoveNotFound {

	return &CreatePersonallyProcuredMoveNotFound{}
}

// WriteResponse to the client
func (o *CreatePersonallyProcuredMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// CreatePersonallyProcuredMoveInternalServerErrorCode is the HTTP code returned for type CreatePersonallyProcuredMoveInternalServerError
const CreatePersonallyProcuredMoveInternalServerErrorCode int = 500

/*
CreatePersonallyProcuredMoveInternalServerError server error

swagger:response createPersonallyProcuredMoveInternalServerError
*/
type CreatePersonallyProcuredMoveInternalServerError struct {
}

// NewCreatePersonallyProcuredMoveInternalServerError creates CreatePersonallyProcuredMoveInternalServerError with default headers values
func NewCreatePersonallyProcuredMoveInternalServerError() *CreatePersonallyProcuredMoveInternalServerError {

	return &CreatePersonallyProcuredMoveInternalServerError{}
}

// WriteResponse to the client
func (o *CreatePersonallyProcuredMoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
