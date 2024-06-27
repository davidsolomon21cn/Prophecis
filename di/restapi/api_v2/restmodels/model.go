// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Model Model类型的全局变量的值类型
// swagger:model Model
type Model struct {

	// 类似GPU节点的 参数设置-任务执行设置-高级参数配置-模型用户组
	GroupID string `json:"group_id,omitempty"`

	// 类似GPU节点的 参数设置-任务执行设置-高级参数配置-模型名称
	ModelID string `json:"model_id,omitempty"`

	// 类似GPU节点的 参数设置-任务执行设置-高级参数配置-模型版本
	ModelVersionID string `json:"model_version_id,omitempty"`
}

// Validate validates this model
func (m *Model) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model) UnmarshalBinary(b []byte) error {
	var res Model
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
