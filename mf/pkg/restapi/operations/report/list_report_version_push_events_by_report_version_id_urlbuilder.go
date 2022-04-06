// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// ListReportVersionPushEventsByReportVersionIDURL generates an URL for the list report version push events by report version Id operation
type ListReportVersionPushEventsByReportVersionIDURL struct {
	ReportVersionID int64

	CurrentPage *int64
	PageSize    *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListReportVersionPushEventsByReportVersionIDURL) WithBasePath(bp string) *ListReportVersionPushEventsByReportVersionIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListReportVersionPushEventsByReportVersionIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListReportVersionPushEventsByReportVersionIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/mf/v1/reportVersion/Push/{reportVersionId}"

	reportVersionID := swag.FormatInt64(o.ReportVersionID)
	if reportVersionID != "" {
		_path = strings.Replace(_path, "{reportVersionId}", reportVersionID, -1)
	} else {
		return nil, errors.New("reportVersionId is required on ListReportVersionPushEventsByReportVersionIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var currentPageQ string
	if o.CurrentPage != nil {
		currentPageQ = swag.FormatInt64(*o.CurrentPage)
	}
	if currentPageQ != "" {
		qs.Set("currentPage", currentPageQ)
	}

	var pageSizeQ string
	if o.PageSize != nil {
		pageSizeQ = swag.FormatInt64(*o.PageSize)
	}
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListReportVersionPushEventsByReportVersionIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListReportVersionPushEventsByReportVersionIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListReportVersionPushEventsByReportVersionIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListReportVersionPushEventsByReportVersionIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListReportVersionPushEventsByReportVersionIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListReportVersionPushEventsByReportVersionIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
