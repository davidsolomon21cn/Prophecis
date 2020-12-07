// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserFromGroupHandlerFunc turns a function with the right signature into a delete user from group handler
type DeleteUserFromGroupHandlerFunc func(DeleteUserFromGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserFromGroupHandlerFunc) Handle(params DeleteUserFromGroupParams) middleware.Responder {
	return fn(params)
}

// DeleteUserFromGroupHandler interface for that can handle valid delete user from group params
type DeleteUserFromGroupHandler interface {
	Handle(DeleteUserFromGroupParams) middleware.Responder
}

// NewDeleteUserFromGroup creates a new http.Handler for the delete user from group operation
func NewDeleteUserFromGroup(ctx *middleware.Context, handler DeleteUserFromGroupHandler) *DeleteUserFromGroup {
	return &DeleteUserFromGroup{Context: ctx, Handler: handler}
}

/*DeleteUserFromGroup swagger:route DELETE /cc/v1/groups/userGroup/id/{id} Groups deleteUserFromGroup

Returns a user.

Optional extended description in Markdown.

*/
type DeleteUserFromGroup struct {
	Context *middleware.Context
	Handler DeleteUserFromGroupHandler
}

func (o *DeleteUserFromGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserFromGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}