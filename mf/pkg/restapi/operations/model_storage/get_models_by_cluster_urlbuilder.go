// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetModelsByClusterURL generates an URL for the get models by cluster operation
type GetModelsByClusterURL struct {
	Cluster string

	Page     *int64
	QueryStr *string
	Size     *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetModelsByClusterURL) WithBasePath(bp string) *GetModelsByClusterURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetModelsByClusterURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetModelsByClusterURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/mf/v1/models/{cluster}"

	cluster := o.Cluster
	if cluster != "" {
		_path = strings.Replace(_path, "{cluster}", cluster, -1)
	} else {
		return nil, errors.New("cluster is required on GetModelsByClusterURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt64(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var queryStrQ string
	if o.QueryStr != nil {
		queryStrQ = *o.QueryStr
	}
	if queryStrQ != "" {
		qs.Set("query_str", queryStrQ)
	}

	var sizeQ string
	if o.Size != nil {
		sizeQ = swag.FormatInt64(*o.Size)
	}
	if sizeQ != "" {
		qs.Set("size", sizeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetModelsByClusterURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetModelsByClusterURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetModelsByClusterURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetModelsByClusterURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetModelsByClusterURL")
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
func (o *GetModelsByClusterURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
