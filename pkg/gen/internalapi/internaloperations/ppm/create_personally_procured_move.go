// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreatePersonallyProcuredMoveHandlerFunc turns a function with the right signature into a create personally procured move handler
type CreatePersonallyProcuredMoveHandlerFunc func(CreatePersonallyProcuredMoveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePersonallyProcuredMoveHandlerFunc) Handle(params CreatePersonallyProcuredMoveParams) middleware.Responder {
	return fn(params)
}

// CreatePersonallyProcuredMoveHandler interface for that can handle valid create personally procured move params
type CreatePersonallyProcuredMoveHandler interface {
	Handle(CreatePersonallyProcuredMoveParams) middleware.Responder
}

// NewCreatePersonallyProcuredMove creates a new http.Handler for the create personally procured move operation
func NewCreatePersonallyProcuredMove(ctx *middleware.Context, handler CreatePersonallyProcuredMoveHandler) *CreatePersonallyProcuredMove {
	return &CreatePersonallyProcuredMove{Context: ctx, Handler: handler}
}

/*
	CreatePersonallyProcuredMove swagger:route POST /moves/{moveId}/personally_procured_move ppm createPersonallyProcuredMove

# Creates a new PPM for the given move

Create an instance of personally_procured_move tied to the move ID
*/
type CreatePersonallyProcuredMove struct {
	Context *middleware.Context
	Handler CreatePersonallyProcuredMoveHandler
}

func (o *CreatePersonallyProcuredMove) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreatePersonallyProcuredMoveParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
