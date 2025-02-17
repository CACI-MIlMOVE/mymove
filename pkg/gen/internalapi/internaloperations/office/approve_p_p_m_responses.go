// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ApprovePPMOKCode is the HTTP code returned for type ApprovePPMOK
const ApprovePPMOKCode int = 200

/*
ApprovePPMOK updated instance of personally_procured_move

swagger:response approvePPMOK
*/
type ApprovePPMOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PersonallyProcuredMovePayload `json:"body,omitempty"`
}

// NewApprovePPMOK creates ApprovePPMOK with default headers values
func NewApprovePPMOK() *ApprovePPMOK {

	return &ApprovePPMOK{}
}

// WithPayload adds the payload to the approve p p m o k response
func (o *ApprovePPMOK) WithPayload(payload *internalmessages.PersonallyProcuredMovePayload) *ApprovePPMOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the approve p p m o k response
func (o *ApprovePPMOK) SetPayload(payload *internalmessages.PersonallyProcuredMovePayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApprovePPMOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ApprovePPMBadRequestCode is the HTTP code returned for type ApprovePPMBadRequest
const ApprovePPMBadRequestCode int = 400

/*
ApprovePPMBadRequest invalid request

swagger:response approvePPMBadRequest
*/
type ApprovePPMBadRequest struct {
}

// NewApprovePPMBadRequest creates ApprovePPMBadRequest with default headers values
func NewApprovePPMBadRequest() *ApprovePPMBadRequest {

	return &ApprovePPMBadRequest{}
}

// WriteResponse to the client
func (o *ApprovePPMBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ApprovePPMUnauthorizedCode is the HTTP code returned for type ApprovePPMUnauthorized
const ApprovePPMUnauthorizedCode int = 401

/*
ApprovePPMUnauthorized request requires user authentication

swagger:response approvePPMUnauthorized
*/
type ApprovePPMUnauthorized struct {
}

// NewApprovePPMUnauthorized creates ApprovePPMUnauthorized with default headers values
func NewApprovePPMUnauthorized() *ApprovePPMUnauthorized {

	return &ApprovePPMUnauthorized{}
}

// WriteResponse to the client
func (o *ApprovePPMUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ApprovePPMForbiddenCode is the HTTP code returned for type ApprovePPMForbidden
const ApprovePPMForbiddenCode int = 403

/*
ApprovePPMForbidden user is not authorized

swagger:response approvePPMForbidden
*/
type ApprovePPMForbidden struct {
}

// NewApprovePPMForbidden creates ApprovePPMForbidden with default headers values
func NewApprovePPMForbidden() *ApprovePPMForbidden {

	return &ApprovePPMForbidden{}
}

// WriteResponse to the client
func (o *ApprovePPMForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ApprovePPMInternalServerErrorCode is the HTTP code returned for type ApprovePPMInternalServerError
const ApprovePPMInternalServerErrorCode int = 500

/*
ApprovePPMInternalServerError internal server error

swagger:response approvePPMInternalServerError
*/
type ApprovePPMInternalServerError struct {
}

// NewApprovePPMInternalServerError creates ApprovePPMInternalServerError with default headers values
func NewApprovePPMInternalServerError() *ApprovePPMInternalServerError {

	return &ApprovePPMInternalServerError{}
}

// WriteResponse to the client
func (o *ApprovePPMInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
