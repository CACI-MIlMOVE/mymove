// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdatePPMShipment The PPM specific fields of the shipment with values being changed
//
// swagger:model UpdatePPMShipment
type UpdatePPMShipment struct {

	// ZIP
	//
	// The postal code of the destination location where goods are being delivered to.
	// Example: 90210
	// Pattern: ^(\d{5})$
	DestinationPostalCode *string `json:"destinationPostalCode,omitempty"`

	// The estimated weight of the PPM shipment goods being moved.
	// Example: 4200
	EstimatedWeight *int64 `json:"estimatedWeight,omitempty"`

	// Date the customer expects to begin moving from their origin.
	//
	// Format: date
	ExpectedDepartureDate *strfmt.Date `json:"expectedDepartureDate,omitempty"`

	// Indicates whether PPM shipment has pro gear for themselves or their spouse.
	//
	HasProGear *bool `json:"hasProGear,omitempty"`

	// The postal code of the origin location where goods are being moved from.
	// Example: 90210
	// Pattern: ^(\d{5})$
	PickupPostalCode *string `json:"pickupPostalCode,omitempty"`

	// The estimated weight of the pro-gear being moved belonging to the service member.
	ProGearWeight *int64 `json:"proGearWeight,omitempty"`

	// An optional secondary location near the destination where goods will be dropped off.
	// Example: 90210
	// Pattern: ^(\d{5})$
	SecondaryDestinationPostalCode *string `json:"secondaryDestinationPostalCode,omitempty"`

	// An optional secondary pickup location near the origin where additional goods exist.
	// Example: 90210
	// Pattern: ^(\d{5})$
	SecondaryPickupPostalCode *string `json:"secondaryPickupPostalCode,omitempty"`

	// The date that goods will exit the storage location.
	// Format: date
	SitEstimatedDepartureDate *strfmt.Date `json:"sitEstimatedDepartureDate,omitempty"`

	// The date that goods will first enter the storage location.
	// Format: date
	SitEstimatedEntryDate *strfmt.Date `json:"sitEstimatedEntryDate,omitempty"`

	// The estimated weight of the goods being put into storage.
	// Example: 2000
	SitEstimatedWeight *int64 `json:"sitEstimatedWeight,omitempty"`

	// Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.
	//
	// Must be set to `true` when providing `sitLocation`, `sitEstimatedWeight`, `sitEstimatedEntryDate`, and `sitEstimatedDepartureDate` values to calculate the `sitEstimatedCost`.
	//
	SitExpected *bool `json:"sitExpected,omitempty"`

	// sit location
	SitLocation *SITLocationType `json:"sitLocation,omitempty"`

	// The estimated weight of the pro-gear being moved belonging to a spouse.
	SpouseProGearWeight *int64 `json:"spouseProGearWeight,omitempty"`
}

// Validate validates this update p p m shipment
func (m *UpdatePPMShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedDepartureDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDestinationPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitEstimatedDepartureDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitEstimatedEntryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePPMShipment) validateDestinationPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("destinationPostalCode", "body", *m.DestinationPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validateExpectedDepartureDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpectedDepartureDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedDepartureDate", "body", "date", m.ExpectedDepartureDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validatePickupPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.PickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("pickupPostalCode", "body", *m.PickupPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validateSecondaryDestinationPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryDestinationPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("secondaryDestinationPostalCode", "body", *m.SecondaryDestinationPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validateSecondaryPickupPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryPickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("secondaryPickupPostalCode", "body", *m.SecondaryPickupPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validateSitEstimatedDepartureDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SitEstimatedDepartureDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitEstimatedDepartureDate", "body", "date", m.SitEstimatedDepartureDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validateSitEstimatedEntryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SitEstimatedEntryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitEstimatedEntryDate", "body", "date", m.SitEstimatedEntryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePPMShipment) validateSitLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.SitLocation) { // not required
		return nil
	}

	if m.SitLocation != nil {
		if err := m.SitLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sitLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sitLocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update p p m shipment based on the context it is used
func (m *UpdatePPMShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSitLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePPMShipment) contextValidateSitLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.SitLocation != nil {

		if swag.IsZero(m.SitLocation) { // not required
			return nil
		}

		if err := m.SitLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sitLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sitLocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePPMShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePPMShipment) UnmarshalBinary(b []byte) error {
	var res UpdatePPMShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
