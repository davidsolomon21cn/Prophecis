// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GlobalVariable global variable
// swagger:model GlobalVariable
type GlobalVariable struct {

	// 全局变量，变量的key
	Key string `json:"key,omitempty"`

	// 全局变量，变量的类型
	// Enum: [string number model processline]
	Type string `json:"type,omitempty"`

	// 全局变量的值，根据类型的不同，可能是String, Number, Model, ProcessLine; Model的数据结构请看GetPipelineGlobalVariablesModel的返回，PorcessLine的数据结构请看GetPipelineGlobalVariablesProcessLine的返回
	// Enum: [String Number Model ProcessLine]
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this global variable
func (m *GlobalVariable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var globalVariableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","number","model","processline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalVariableTypeTypePropEnum = append(globalVariableTypeTypePropEnum, v)
	}
}

const (

	// GlobalVariableTypeString captures enum value "string"
	GlobalVariableTypeString string = "string"

	// GlobalVariableTypeNumber captures enum value "number"
	GlobalVariableTypeNumber string = "number"

	// GlobalVariableTypeModel captures enum value "model"
	GlobalVariableTypeModel string = "model"

	// GlobalVariableTypeProcessline captures enum value "processline"
	GlobalVariableTypeProcessline string = "processline"
)

// prop value enum
func (m *GlobalVariable) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalVariableTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalVariable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalVariable) UnmarshalBinary(b []byte) error {
	var res GlobalVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
