// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PermissionGrant A permission resolved in a given context, such as a user in an entity
// swagger:model PermissionGrant
type PermissionGrant struct {

	// Whether this permission has been granted in the context or a sub-context
	// Enum: [all some none]
	Granted string `json:"granted,omitempty"`

	// permission
	Permission *PermissionSummary `json:"permission,omitempty"`
}

// Validate validates this permission grant
func (m *PermissionGrant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGranted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var permissionGrantTypeGrantedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","some","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		permissionGrantTypeGrantedPropEnum = append(permissionGrantTypeGrantedPropEnum, v)
	}
}

const (

	// PermissionGrantGrantedAll captures enum value "all"
	PermissionGrantGrantedAll string = "all"

	// PermissionGrantGrantedSome captures enum value "some"
	PermissionGrantGrantedSome string = "some"

	// PermissionGrantGrantedNone captures enum value "none"
	PermissionGrantGrantedNone string = "none"
)

// prop value enum
func (m *PermissionGrant) validateGrantedEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, permissionGrantTypeGrantedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PermissionGrant) validateGranted(formats strfmt.Registry) error {

	if swag.IsZero(m.Granted) { // not required
		return nil
	}

	// value enum
	if err := m.validateGrantedEnum("granted", "body", m.Granted); err != nil {
		return err
	}

	return nil
}

func (m *PermissionGrant) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if m.Permission != nil {
		if err := m.Permission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PermissionGrant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionGrant) UnmarshalBinary(b []byte) error {
	var res PermissionGrant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}