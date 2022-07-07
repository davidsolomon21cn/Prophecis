// Code generated by go-swagger; DO NOT EDIT.

package experiment_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetExperimentRunHandlerFunc turns a function with the right signature into a get experiment run handler
type GetExperimentRunHandlerFunc func(GetExperimentRunParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetExperimentRunHandlerFunc) Handle(params GetExperimentRunParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetExperimentRunHandler interface for that can handle valid get experiment run params
type GetExperimentRunHandler interface {
	Handle(GetExperimentRunParams, interface{}) middleware.Responder
}

// NewGetExperimentRun creates a new http.Handler for the get experiment run operation
func NewGetExperimentRun(ctx *middleware.Context, handler GetExperimentRunHandler) *GetExperimentRun {
	return &GetExperimentRun{Context: ctx, Handler: handler}
}

/*GetExperimentRun swagger:route GET /di/v1/experimentRun/{id} ExperimentRuns getExperimentRun

Get an Experiment Flow Exec Record By Experiment Id & Flow_Exec_ID.

Get an Experiment Flow Exec Record By Experiment Id & Flow_Exec_ID.

*/
type GetExperimentRun struct {
	Context *middleware.Context
	Handler GetExperimentRunHandler
}

func (o *GetExperimentRun) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetExperimentRunParams()

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