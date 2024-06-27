// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateExperimentVersionResponse create experiment version response
// swagger:model CreateExperimentVersionResponse
type CreateExperimentVersionResponse struct {

	// 返回固化的版本
	ArchivedVersionName string `json:"archived_version_name,omitempty"`

	// 返回当前最新的版本（一般是固化的版本+1）
	LatestVersionName string `json:"latest_version_name,omitempty"`
}

// Validate validates this create experiment version response
func (m *CreateExperimentVersionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateExperimentVersionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateExperimentVersionResponse) UnmarshalBinary(b []byte) error {
	var res CreateExperimentVersionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
