// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServicePostModel service post model
// swagger:model ServicePostModel
type ServicePostModel struct {

	// serivce model name
	EndpointType string `json:"endpoint_type,omitempty"`

	// serivce model name
	GroupName string `json:"group_name,omitempty"`

	// Custom image
	Image string `json:"image,omitempty"`

	// serivce model name
	ModelName string `json:"model_name,omitempty"`

	// serivce model name
	ModelType string `json:"model_type,omitempty"`

	// serivce model name
	ModelVersion string `json:"model_version,omitempty"`

	// the modelversion id of this service.
	ModelversionID int64 `json:"modelversion_id,omitempty"`

	// parameters
	Parameters []*ModelParameters `json:"parameters"`

	// model s3 path
	Source string `json:"source,omitempty"`
}

// Validate validates this service post model
func (m *ServicePostModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePostModel) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServicePostModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePostModel) UnmarshalBinary(b []byte) error {
	var res ServicePostModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
