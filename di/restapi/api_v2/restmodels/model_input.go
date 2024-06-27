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

// ModelInput Model类型的全局变量的值类型
// swagger:model ModelInput
type ModelInput struct {

	// 类似GPU节点的 参数设置-任务执行设置-高级参数配置-模型用户组
	GroupID string `json:"group_id,omitempty"`

	// 类似GPU节点的 参数设置-任务执行设置-高级参数配置-模型名称
	ModelID string `json:"model_id,omitempty"`

	// 类似GPU节点的 参数设置-任务执行设置-高级参数配置-模型版本
	ModelVersionID int64 `json:"model_version_id,omitempty"`

	// 模型数据的来源类型，枚举值
	//
	// 选择全局变量要求在全局参数里定义了模型类型的变量，这里便可以选择该变量，然后在执行工作流和发布工作流的时候可以动态修改该变量；选择自定义则是普通的设置模型相关的设置
	// Enum: [Direct Variable]
	SourceType string `json:"source_type,omitempty"`

	// 当source_type为variable的时候，需填充该值为全局变量的某个model类型的key
	VariableKey string `json:"variable_key,omitempty"`
}

// Validate validates this model input
func (m *ModelInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modelInputTypeSourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Direct","Variable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelInputTypeSourceTypePropEnum = append(modelInputTypeSourceTypePropEnum, v)
	}
}

const (

	// ModelInputSourceTypeDirect captures enum value "Direct"
	ModelInputSourceTypeDirect string = "Direct"

	// ModelInputSourceTypeVariable captures enum value "Variable"
	ModelInputSourceTypeVariable string = "Variable"
)

// prop value enum
func (m *ModelInput) validateSourceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, modelInputTypeSourceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ModelInput) validateSourceType(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceTypeEnum("source_type", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelInput) UnmarshalBinary(b []byte) error {
	var res ModelInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
