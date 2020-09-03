// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SysAdminFlag sys admin flag
//
// swagger:model SysAdminFlag
type SysAdminFlag struct {

	// true-admin, false-not admin.
	SysadminFlag bool `json:"sysadmin_flag,omitempty"`
}

// Validate validates this sys admin flag
func (m *SysAdminFlag) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SysAdminFlag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SysAdminFlag) UnmarshalBinary(b []byte) error {
	var res SysAdminFlag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}