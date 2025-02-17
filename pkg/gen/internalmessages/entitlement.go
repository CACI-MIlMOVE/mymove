// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Entitlement entitlement
//
// swagger:model Entitlement
type Entitlement struct {

	// Pro-gear weight limit as set by an Office user, distinct from the service member's default weight allotment determined by rank
	//
	// Example: 2000
	ProGear *int64 `json:"proGear,omitempty"`

	// Spouse's pro-gear weight limit as set by an Office user, distinct from the service member's default weight allotment determined by rank
	//
	// Example: 500
	ProGearSpouse *int64 `json:"proGearSpouse,omitempty"`
}

// Validate validates this entitlement
func (m *Entitlement) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this entitlement based on context it is used
func (m *Entitlement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Entitlement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entitlement) UnmarshalBinary(b []byte) error {
	var res Entitlement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
