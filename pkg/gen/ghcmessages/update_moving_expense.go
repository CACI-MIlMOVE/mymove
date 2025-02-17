// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateMovingExpense update moving expense
//
// swagger:model UpdateMovingExpense
type UpdateMovingExpense struct {

	// The total amount of the expense as indicated on the receipt
	Amount int64 `json:"amount,omitempty"`

	// The reason the services counselor has excluded or rejected the item.
	Reason string `json:"reason,omitempty"`

	// The date the shipment exited storage, applicable for the `STORAGE` movingExpenseType only
	// Format: date
	SitEndDate strfmt.Date `json:"sitEndDate,omitempty"`

	// The date the shipment entered storage, applicable for the `STORAGE` movingExpenseType only
	// Format: date
	SitStartDate strfmt.Date `json:"sitStartDate,omitempty"`

	// status
	Status PPMDocumentStatus `json:"status,omitempty"`
}

// Validate validates this update moving expense
func (m *UpdateMovingExpense) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSitEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMovingExpense) validateSitEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SitEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitEndDate", "body", "date", m.SitEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMovingExpense) validateSitStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SitStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitStartDate", "body", "date", m.SitStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMovingExpense) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this update moving expense based on the context it is used
func (m *UpdateMovingExpense) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMovingExpense) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMovingExpense) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMovingExpense) UnmarshalBinary(b []byte) error {
	var res UpdateMovingExpense
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
