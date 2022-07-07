// Code generated by go-swagger; DO NOT EDIT.

package experiment_runs

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

// NewGetExperimentRunsHistoryParams creates a new GetExperimentRunsHistoryParams object
// with the default values initialized.
func NewGetExperimentRunsHistoryParams() *GetExperimentRunsHistoryParams {
	var ()
	return &GetExperimentRunsHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExperimentRunsHistoryParamsWithTimeout creates a new GetExperimentRunsHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExperimentRunsHistoryParamsWithTimeout(timeout time.Duration) *GetExperimentRunsHistoryParams {
	var ()
	return &GetExperimentRunsHistoryParams{

		timeout: timeout,
	}
}

// NewGetExperimentRunsHistoryParamsWithContext creates a new GetExperimentRunsHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExperimentRunsHistoryParamsWithContext(ctx context.Context) *GetExperimentRunsHistoryParams {
	var ()
	return &GetExperimentRunsHistoryParams{

		Context: ctx,
	}
}

// NewGetExperimentRunsHistoryParamsWithHTTPClient creates a new GetExperimentRunsHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExperimentRunsHistoryParamsWithHTTPClient(client *http.Client) *GetExperimentRunsHistoryParams {
	var ()
	return &GetExperimentRunsHistoryParams{
		HTTPClient: client,
	}
}

/*GetExperimentRunsHistoryParams contains all the parameters to send to the API endpoint
for the get experiment runs history operation typically these are written to a http.Request
*/
type GetExperimentRunsHistoryParams struct {

	/*ExpID*/
	ExpID int64
	/*Page*/
	Page *int64
	/*QueryStr*/
	QueryStr *string
	/*Size*/
	Size *int64
	/*Username*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithTimeout(timeout time.Duration) *GetExperimentRunsHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithContext(ctx context.Context) *GetExperimentRunsHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithHTTPClient(client *http.Client) *GetExperimentRunsHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpID adds the expID to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithExpID(expID int64) *GetExperimentRunsHistoryParams {
	o.SetExpID(expID)
	return o
}

// SetExpID adds the expId to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetExpID(expID int64) {
	o.ExpID = expID
}

// WithPage adds the page to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithPage(page *int64) *GetExperimentRunsHistoryParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetPage(page *int64) {
	o.Page = page
}

// WithQueryStr adds the queryStr to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithQueryStr(queryStr *string) *GetExperimentRunsHistoryParams {
	o.SetQueryStr(queryStr)
	return o
}

// SetQueryStr adds the queryStr to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetQueryStr(queryStr *string) {
	o.QueryStr = queryStr
}

// WithSize adds the size to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithSize(size *int64) *GetExperimentRunsHistoryParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetSize(size *int64) {
	o.Size = size
}

// WithUsername adds the username to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) WithUsername(username *string) *GetExperimentRunsHistoryParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get experiment runs history params
func (o *GetExperimentRunsHistoryParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetExperimentRunsHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param exp_id
	if err := r.SetPathParam("exp_id", swag.FormatInt64(o.ExpID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.QueryStr != nil {

		// query param query_str
		var qrQueryStr string
		if o.QueryStr != nil {
			qrQueryStr = *o.QueryStr
		}
		qQueryStr := qrQueryStr
		if qQueryStr != "" {
			if err := r.SetQueryParam("query_str", qQueryStr); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int64
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if o.Username != nil {

		// query param username
		var qrUsername string
		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {
			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
