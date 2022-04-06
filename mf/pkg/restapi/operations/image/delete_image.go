// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteImageHandlerFunc turns a function with the right signature into a delete image handler
type DeleteImageHandlerFunc func(DeleteImageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteImageHandlerFunc) Handle(params DeleteImageParams) middleware.Responder {
	return fn(params)
}

// DeleteImageHandler interface for that can handle valid delete image params
type DeleteImageHandler interface {
	Handle(DeleteImageParams) middleware.Responder
}

// NewDeleteImage creates a new http.Handler for the delete image operation
func NewDeleteImage(ctx *middleware.Context, handler DeleteImageHandler) *DeleteImage {
	return &DeleteImage{Context: ctx, Handler: handler}
}

/*DeleteImage swagger:route DELETE /mf/v1/image/{image_id} image deleteImage

Delete the Image in given image id

Delete Image.

*/
type DeleteImage struct {
	Context *middleware.Context
	Handler DeleteImageHandler
}

func (o *DeleteImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteImageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
