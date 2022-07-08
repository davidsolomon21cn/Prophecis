// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProphecisExperimentPutBasicInfo prophecis experiment put basic info
// swagger:model ProphecisExperimentPutBasicInfo
type ProphecisExperimentPutBasicInfo struct {

	// Experiment Description.
	ExpDesc string `json:"exp_desc,omitempty"`

	// Experiment Name.
	ExpName string `json:"exp_name,omitempty"`

	// Experiment group.
	GroupName string `json:"group_name,omitempty"`

	// Experiment Tags
	TagList []*ProphecisExperimentTagPutBasicInfo `json:"tag_list"`
}

// Validate validates this prophecis experiment put basic info
func (m *ProphecisExperimentPutBasicInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProphecisExperimentPutBasicInfo) validateTagList(formats strfmt.Registry) error {

	if swag.IsZero(m.TagList) { // not required
		return nil
	}

	for i := 0; i < len(m.TagList); i++ {
		if swag.IsZero(m.TagList[i]) { // not required
			continue
		}

		if m.TagList[i] != nil {
			if err := m.TagList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProphecisExperimentPutBasicInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProphecisExperimentPutBasicInfo) UnmarshalBinary(b []byte) error {
	var res ProphecisExperimentPutBasicInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
