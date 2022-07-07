// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewExportExperimentParams creates a new ExportExperimentParams object
// with the default values initialized.
func NewExportExperimentParams() *ExportExperimentParams {
	var ()
	return &ExportExperimentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportExperimentParamsWithTimeout creates a new ExportExperimentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportExperimentParamsWithTimeout(timeout time.Duration) *ExportExperimentParams {
	var ()
	return &ExportExperimentParams{

		timeout: timeout,
	}
}

// NewExportExperimentParamsWithContext creates a new ExportExperimentParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportExperimentParamsWithContext(ctx context.Context) *ExportExperimentParams {
	var ()
	return &ExportExperimentParams{

		Context: ctx,
	}
}

// NewExportExperimentParamsWithHTTPClient creates a new ExportExperimentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportExperimentParamsWithHTTPClient(client *http.Client) *ExportExperimentParams {
	var ()
	return &ExportExperimentParams{
		HTTPClient: client,
	}
}

/*ExportExperimentParams contains all the parameters to send to the API endpoint
for the export experiment operation typically these are written to a http.Request
*/
type ExportExperimentParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export experiment params
func (o *ExportExperimentParams) WithTimeout(timeout time.Duration) *ExportExperimentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export experiment params
func (o *ExportExperimentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export experiment params
func (o *ExportExperimentParams) WithContext(ctx context.Context) *ExportExperimentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export experiment params
func (o *ExportExperimentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export experiment params
func (o *ExportExperimentParams) WithHTTPClient(client *http.Client) *ExportExperimentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export experiment params
func (o *ExportExperimentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the export experiment params
func (o *ExportExperimentParams) WithID(id int64) *ExportExperimentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the export experiment params
func (o *ExportExperimentParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExportExperimentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
