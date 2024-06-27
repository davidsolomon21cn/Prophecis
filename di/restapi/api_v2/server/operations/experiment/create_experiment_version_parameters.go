// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	restmodels "webank/DI/restapi/api_v2/restmodels"
)

// NewCreateExperimentVersionParams creates a new CreateExperimentVersionParams object
// no default values defined in spec.
func NewCreateExperimentVersionParams() CreateExperimentVersionParams {

	return CreateExperimentVersionParams{}
}

// CreateExperimentVersionParams contains all the bound params for the create experiment version operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateExperimentVersion
type CreateExperimentVersionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*experiment id
	  Required: true
	  In: path
	*/
	ExpID string
	/*新建实验的具体内容
	  In: body
	*/
	Experiment *restmodels.CreateExperimentVersionRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateExperimentVersionParams() beforehand.
func (o *CreateExperimentVersionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rExpID, rhkExpID, _ := route.Params.GetOK("exp_id")
	if err := o.bindExpID(rExpID, rhkExpID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body restmodels.CreateExperimentVersionRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("experiment", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Experiment = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExpID binds and validates parameter ExpID from path.
func (o *CreateExperimentVersionParams) bindExpID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ExpID = raw

	return nil
}
