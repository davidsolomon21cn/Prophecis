// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ComputeResource compute resource
// swagger:model ComputeResource
type ComputeResource struct {

	// 该计算资源类型所申请的cpu个数
	Cpus float64 `json:"cpus,omitempty"`

	// 该计算资源所申请的gpu个数
	Gpus float64 `json:"gpus,omitempty"`

	// 该计算资源类型所申请的内存大小
	Memory string `json:"memory,omitempty"`

	// 该计算资源类型的个数
	Nums float64 `json:"nums,omitempty"`
}

// Validate validates this compute resource
func (m *ComputeResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputeResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeResource) UnmarshalBinary(b []byte) error {
	var res ComputeResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}