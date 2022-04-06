// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PushModelByModelIDHandlerFunc turns a function with the right signature into a push model by model Id handler
type PushModelByModelIDHandlerFunc func(PushModelByModelIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PushModelByModelIDHandlerFunc) Handle(params PushModelByModelIDParams) middleware.Responder {
	return fn(params)
}

// PushModelByModelIDHandler interface for that can handle valid push model by model Id params
type PushModelByModelIDHandler interface {
	Handle(PushModelByModelIDParams) middleware.Responder
}

// NewPushModelByModelID creates a new http.Handler for the push model by model Id operation
func NewPushModelByModelID(ctx *middleware.Context, handler PushModelByModelIDHandler) *PushModelByModelID {
	return &PushModelByModelID{Context: ctx, Handler: handler}
}

/*PushModelByModelID swagger:route POST /mf/v1/model/push/{modelId} modelStorage pushModelByModelId

push model by model id

push model by model id

*/
type PushModelByModelID struct {
	Context *middleware.Context
	Handler PushModelByModelIDHandler
}

func (o *PushModelByModelID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPushModelByModelIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}