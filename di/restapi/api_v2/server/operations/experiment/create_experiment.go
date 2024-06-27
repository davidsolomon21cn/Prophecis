// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateExperimentHandlerFunc turns a function with the right signature into a create experiment handler
type CreateExperimentHandlerFunc func(CreateExperimentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateExperimentHandlerFunc) Handle(params CreateExperimentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateExperimentHandler interface for that can handle valid create experiment params
type CreateExperimentHandler interface {
	Handle(CreateExperimentParams, interface{}) middleware.Responder
}

// NewCreateExperiment creates a new http.Handler for the create experiment operation
func NewCreateExperiment(ctx *middleware.Context, handler CreateExperimentHandler) *CreateExperiment {
	return &CreateExperiment{Context: ctx, Handler: handler}
}

/*CreateExperiment swagger:route POST /di/v2/experiment Experiment createExperiment

CreateExperiment

创建一个新实验

*/
type CreateExperiment struct {
	Context *middleware.Context
	Handler CreateExperimentHandler
}

func (o *CreateExperiment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateExperimentParams()

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