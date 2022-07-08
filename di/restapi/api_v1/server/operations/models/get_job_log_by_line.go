// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetJobLogByLineHandlerFunc turns a function with the right signature into a get job log by line handler
type GetJobLogByLineHandlerFunc func(GetJobLogByLineParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetJobLogByLineHandlerFunc) Handle(params GetJobLogByLineParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetJobLogByLineHandler interface for that can handle valid get job log by line params
type GetJobLogByLineHandler interface {
	Handle(GetJobLogByLineParams, interface{}) middleware.Responder
}

// NewGetJobLogByLine creates a new http.Handler for the get job log by line operation
func NewGetJobLogByLine(ctx *middleware.Context, handler GetJobLogByLineHandler) *GetJobLogByLine {
	return &GetJobLogByLine{Context: ctx, Handler: handler}
}

/*GetJobLogByLine swagger:route GET /di/v1/job/{training_id}/log Models getJobLogByLine

Get Job Log By Exec Id And Line.

Get Job Log By Exec Id And Line.

*/
type GetJobLogByLine struct {
	Context *middleware.Context
	Handler GetJobLogByLineHandler
}

func (o *GetJobLogByLine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetJobLogByLineParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
