// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteModelParams creates a new DeleteModelParams object
// with the default values initialized.
func NewDeleteModelParams() *DeleteModelParams {
	var (
		versionDefault = string("2017-02-13")
	)
	return &DeleteModelParams{
		Version: &versionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModelParamsWithTimeout creates a new DeleteModelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteModelParamsWithTimeout(timeout time.Duration) *DeleteModelParams {
	var (
		versionDefault = string("2017-02-13")
	)
	return &DeleteModelParams{
		Version: &versionDefault,

		timeout: timeout,
	}
}

// NewDeleteModelParamsWithContext creates a new DeleteModelParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteModelParamsWithContext(ctx context.Context) *DeleteModelParams {
	var (
		versionDefault = string("2017-02-13")
	)
	return &DeleteModelParams{
		Version: &versionDefault,

		Context: ctx,
	}
}

// NewDeleteModelParamsWithHTTPClient creates a new DeleteModelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteModelParamsWithHTTPClient(client *http.Client) *DeleteModelParams {
	var (
		versionDefault = string("2017-02-13")
	)
	return &DeleteModelParams{
		Version:    &versionDefault,
		HTTPClient: client,
	}
}

/*DeleteModelParams contains all the parameters to send to the API endpoint
for the delete model operation typically these are written to a http.Request
*/
type DeleteModelParams struct {

	/*ModelID
	  The id of the model.

	*/
	ModelID string
	/*Version
	  The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete model params
func (o *DeleteModelParams) WithTimeout(timeout time.Duration) *DeleteModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete model params
func (o *DeleteModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete model params
func (o *DeleteModelParams) WithContext(ctx context.Context) *DeleteModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete model params
func (o *DeleteModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete model params
func (o *DeleteModelParams) WithHTTPClient(client *http.Client) *DeleteModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete model params
func (o *DeleteModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelID adds the modelID to the delete model params
func (o *DeleteModelParams) WithModelID(modelID string) *DeleteModelParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the delete model params
func (o *DeleteModelParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WithVersion adds the version to the delete model params
func (o *DeleteModelParams) WithVersion(version *string) *DeleteModelParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete model params
func (o *DeleteModelParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param model_id
	if err := r.SetPathParam("model_id", o.ModelID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
