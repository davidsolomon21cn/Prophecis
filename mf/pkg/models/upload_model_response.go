// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UploadModelResponse upload model response
// swagger:model UploadModelResponse
type UploadModelResponse struct {

	// file name
	FileName string `json:"fileName,omitempty"`

	// S3 Path
	S3Path string `json:"s3Path,omitempty"`
}

// Validate validates this upload model response
func (m *UploadModelResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadModelResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadModelResponse) UnmarshalBinary(b []byte) error {
	var res UploadModelResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
