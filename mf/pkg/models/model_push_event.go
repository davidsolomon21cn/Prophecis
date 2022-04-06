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

// ModelPushEvent model push event
// swagger:model ModelPushEvent
type ModelPushEvent struct {

	// Timestamp recorded when metadata for this model was last updated.
	CreationTimestamp string `json:"creation_timestamp,omitempty"`

	// enable flag
	EnableFlag int8 `json:"enable_flag,omitempty"`

	// group
	Group *Group `json:"group,omitempty"`

	// Group Id of User.
	GroupID int64 `json:"group_id,omitempty"`

	// id for the model.
	ID int64 `json:"id,omitempty"`

	// models of servier.
	ModelLatestVersion *ModelVersionInfo `json:"model_latest_version,omitempty"`

	// model latest version id
	ModelLatestVersionID int64 `json:"model_latest_version_id,omitempty"`

	// Name for the model.
	ModelName string `json:"model_name,omitempty"`

	// type of model, eg. tensorflow, sklearn, xgboost
	ModelType string `json:"model_type,omitempty"`

	// model version events
	ModelVersionEvents []*ModelVersionInfo `json:"model_version_events"`

	// Timestamp recorded when metadata for this model was last updated.
	Position int64 `json:"position,omitempty"`

	// Description of this model
	Reamrk string `json:"reamrk,omitempty"`

	// User that created this model.
	ServiceID int64 `json:"service_id,omitempty"`

	// Update time
	UpdateTimestamp string `json:"update_timestamp,omitempty"`

	// User that created this model.
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this model push event
func (m *ModelPushEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelLatestVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelVersionEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelPushEvent) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *ModelPushEvent) validateModelLatestVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.ModelLatestVersion) { // not required
		return nil
	}

	if m.ModelLatestVersion != nil {
		if err := m.ModelLatestVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_latest_version")
			}
			return err
		}
	}

	return nil
}

func (m *ModelPushEvent) validateModelVersionEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.ModelVersionEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelVersionEvents); i++ {
		if swag.IsZero(m.ModelVersionEvents[i]) { // not required
			continue
		}

		if m.ModelVersionEvents[i] != nil {
			if err := m.ModelVersionEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("model_version_events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelPushEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelPushEvent) UnmarshalBinary(b []byte) error {
	var res ModelPushEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
