// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// SubmitPersonallyProcuredMoveOKCode is the HTTP code returned for type SubmitPersonallyProcuredMoveOK
const SubmitPersonallyProcuredMoveOKCode int = 200

/*
SubmitPersonallyProcuredMoveOK updated instance of personally_procured_move

swagger:response submitPersonallyProcuredMoveOK
*/
type SubmitPersonallyProcuredMoveOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PersonallyProcuredMovePayload `json:"body,omitempty"`
}

// NewSubmitPersonallyProcuredMoveOK creates SubmitPersonallyProcuredMoveOK with default headers values
func NewSubmitPersonallyProcuredMoveOK() *SubmitPersonallyProcuredMoveOK {

	return &SubmitPersonallyProcuredMoveOK{}
}

// WithPayload adds the payload to the submit personally procured move o k response
func (o *SubmitPersonallyProcuredMoveOK) WithPayload(payload *internalmessages.PersonallyProcuredMovePayload) *SubmitPersonallyProcuredMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit personally procured move o k response
func (o *SubmitPersonallyProcuredMoveOK) SetPayload(payload *internalmessages.PersonallyProcuredMovePayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitPersonallyProcuredMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubmitPersonallyProcuredMoveBadRequestCode is the HTTP code returned for type SubmitPersonallyProcuredMoveBadRequest
const SubmitPersonallyProcuredMoveBadRequestCode int = 400

/*
SubmitPersonallyProcuredMoveBadRequest invalid request

swagger:response submitPersonallyProcuredMoveBadRequest
*/
type SubmitPersonallyProcuredMoveBadRequest struct {
}

// NewSubmitPersonallyProcuredMoveBadRequest creates SubmitPersonallyProcuredMoveBadRequest with default headers values
func NewSubmitPersonallyProcuredMoveBadRequest() *SubmitPersonallyProcuredMoveBadRequest {

	return &SubmitPersonallyProcuredMoveBadRequest{}
}

// WriteResponse to the client
func (o *SubmitPersonallyProcuredMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SubmitPersonallyProcuredMoveUnauthorizedCode is the HTTP code returned for type SubmitPersonallyProcuredMoveUnauthorized
const SubmitPersonallyProcuredMoveUnauthorizedCode int = 401

/*
SubmitPersonallyProcuredMoveUnauthorized request requires user authentication

swagger:response submitPersonallyProcuredMoveUnauthorized
*/
type SubmitPersonallyProcuredMoveUnauthorized struct {
}

// NewSubmitPersonallyProcuredMoveUnauthorized creates SubmitPersonallyProcuredMoveUnauthorized with default headers values
func NewSubmitPersonallyProcuredMoveUnauthorized() *SubmitPersonallyProcuredMoveUnauthorized {

	return &SubmitPersonallyProcuredMoveUnauthorized{}
}

// WriteResponse to the client
func (o *SubmitPersonallyProcuredMoveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SubmitPersonallyProcuredMoveForbiddenCode is the HTTP code returned for type SubmitPersonallyProcuredMoveForbidden
const SubmitPersonallyProcuredMoveForbiddenCode int = 403

/*
SubmitPersonallyProcuredMoveForbidden user is not authorized

swagger:response submitPersonallyProcuredMoveForbidden
*/
type SubmitPersonallyProcuredMoveForbidden struct {
}

// NewSubmitPersonallyProcuredMoveForbidden creates SubmitPersonallyProcuredMoveForbidden with default headers values
func NewSubmitPersonallyProcuredMoveForbidden() *SubmitPersonallyProcuredMoveForbidden {

	return &SubmitPersonallyProcuredMoveForbidden{}
}

// WriteResponse to the client
func (o *SubmitPersonallyProcuredMoveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// SubmitPersonallyProcuredMoveNotFoundCode is the HTTP code returned for type SubmitPersonallyProcuredMoveNotFound
const SubmitPersonallyProcuredMoveNotFoundCode int = 404

/*
SubmitPersonallyProcuredMoveNotFound ppm is not found

swagger:response submitPersonallyProcuredMoveNotFound
*/
type SubmitPersonallyProcuredMoveNotFound struct {
}

// NewSubmitPersonallyProcuredMoveNotFound creates SubmitPersonallyProcuredMoveNotFound with default headers values
func NewSubmitPersonallyProcuredMoveNotFound() *SubmitPersonallyProcuredMoveNotFound {

	return &SubmitPersonallyProcuredMoveNotFound{}
}

// WriteResponse to the client
func (o *SubmitPersonallyProcuredMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// SubmitPersonallyProcuredMoveInternalServerErrorCode is the HTTP code returned for type SubmitPersonallyProcuredMoveInternalServerError
const SubmitPersonallyProcuredMoveInternalServerErrorCode int = 500

/*
SubmitPersonallyProcuredMoveInternalServerError internal server error

swagger:response submitPersonallyProcuredMoveInternalServerError
*/
type SubmitPersonallyProcuredMoveInternalServerError struct {
}

// NewSubmitPersonallyProcuredMoveInternalServerError creates SubmitPersonallyProcuredMoveInternalServerError with default headers values
func NewSubmitPersonallyProcuredMoveInternalServerError() *SubmitPersonallyProcuredMoveInternalServerError {

	return &SubmitPersonallyProcuredMoveInternalServerError{}
}

// WriteResponse to the client
func (o *SubmitPersonallyProcuredMoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
