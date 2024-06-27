// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateExperimentTestHandlerFunc turns a function with the right signature into a create experiment test handler
type CreateExperimentTestHandlerFunc func(CreateExperimentTestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateExperimentTestHandlerFunc) Handle(params CreateExperimentTestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateExperimentTestHandler interface for that can handle valid create experiment test params
type CreateExperimentTestHandler interface {
	Handle(CreateExperimentTestParams, interface{}) middleware.Responder
}

// NewCreateExperimentTest creates a new http.Handler for the create experiment test operation
func NewCreateExperimentTest(ctx *middleware.Context, handler CreateExperimentTestHandler) *CreateExperimentTest {
	return &CreateExperimentTest{Context: ctx, Handler: handler}
}

/*CreateExperimentTest swagger:route POST /di/v2/experiment_test Experiment createExperimentTest

CreateExperimentTest

仅仅是为了测试接口

*/
type CreateExperimentTest struct {
	Context *middleware.Context
	Handler CreateExperimentTestHandler
}

func (o *CreateExperimentTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateExperimentTestParams()

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