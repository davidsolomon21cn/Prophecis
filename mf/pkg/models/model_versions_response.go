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

// ModelVersionsResponse model versions response
// swagger:model ModelVersionsResponse
type ModelVersionsResponse struct {

	// model versions
	ModelVersions []*ModelVersion `json:"modelVersions"`

	// page number
	PageNumber int64 `json:"page_number,omitempty"`

	// page size
	PageSize int64 `json:"page_size,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`

	// total page
	TotalPage int64 `json:"total_page,omitempty"`
}

// Validate validates this model versions response
func (m *ModelVersionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModelVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelVersionsResponse) validateModelVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.ModelVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelVersions); i++ {
		if swag.IsZero(m.ModelVersions[i]) { // not required
			continue
		}

		if m.ModelVersions[i] != nil {
			if err := m.ModelVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modelVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelVersionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelVersionsResponse) UnmarshalBinary(b []byte) error {
	var res ModelVersionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
