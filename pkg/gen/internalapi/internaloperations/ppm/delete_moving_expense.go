// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMovingExpenseHandlerFunc turns a function with the right signature into a delete moving expense handler
type DeleteMovingExpenseHandlerFunc func(DeleteMovingExpenseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMovingExpenseHandlerFunc) Handle(params DeleteMovingExpenseParams) middleware.Responder {
	return fn(params)
}

// DeleteMovingExpenseHandler interface for that can handle valid delete moving expense params
type DeleteMovingExpenseHandler interface {
	Handle(DeleteMovingExpenseParams) middleware.Responder
}

// NewDeleteMovingExpense creates a new http.Handler for the delete moving expense operation
func NewDeleteMovingExpense(ctx *middleware.Context, handler DeleteMovingExpenseHandler) *DeleteMovingExpense {
	return &DeleteMovingExpense{Context: ctx, Handler: handler}
}

/*
	DeleteMovingExpense swagger:route DELETE /ppm-shipments/{ppmShipmentId}/moving-expenses/{movingExpenseId} ppm deleteMovingExpense

# Soft deletes a moving expense by ID

Removes a single moving expense receipt from the closeout line items for a PPM shipment. Soft deleted
records are not visible in milmove, but are kept in the database.
*/
type DeleteMovingExpense struct {
	Context *middleware.Context
	Handler DeleteMovingExpenseHandler
}

func (o *DeleteMovingExpense) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMovingExpenseParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
