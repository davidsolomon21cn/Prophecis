// Code generated by go-swagger; DO NOT EDIT.

package model_deploy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "mlss-mf/pkg/models"
)

// NewCreateNamespacedServiceRunParams creates a new CreateNamespacedServiceRunParams object
// no default values defined in spec.
func NewCreateNamespacedServiceRunParams() CreateNamespacedServiceRunParams {

	return CreateNamespacedServiceRunParams{}
}

// CreateNamespacedServiceRunParams contains all the bound params for the create namespaced service run operation
// typically these are obtained from a http.Request
//
// swagger:parameters createNamespacedServiceRun
type CreateNamespacedServiceRunParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The Model Request
	  Required: true
	  In: body
	*/
	Service *models.ServicePost
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateNamespacedServiceRunParams() beforehand.
func (o *CreateNamespacedServiceRunParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ServicePost
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("service", "body"))
			} else {
				res = append(res, errors.NewParseError("service", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Service = &body
			}
		}
	} else {
		res = append(res, errors.Required("service", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
