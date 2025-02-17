// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateShipment update shipment
//
// swagger:model UpdateShipment
type UpdateShipment struct {

	// agents
	Agents MTOAgents `json:"agents,omitempty"`

	// customer remarks
	// Example: handle with care
	CustomerRemarks *string `json:"customerRemarks,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// has secondary delivery address
	HasSecondaryDeliveryAddress *bool `json:"hasSecondaryDeliveryAddress"`

	// has secondary pickup address
	HasSecondaryPickupAddress *bool `json:"hasSecondaryPickupAddress"`

	// pickup address
	PickupAddress *Address `json:"pickupAddress,omitempty"`

	// ppm shipment
	PpmShipment *UpdatePPMShipment `json:"ppmShipment,omitempty"`

	// requested delivery date
	// Format: date
	RequestedDeliveryDate *strfmt.Date `json:"requestedDeliveryDate,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate *strfmt.Date `json:"requestedPickupDate,omitempty"`

	// secondary delivery address
	SecondaryDeliveryAddress *Address `json:"secondaryDeliveryAddress,omitempty"`

	// secondary pickup address
	SecondaryPickupAddress *Address `json:"secondaryPickupAddress,omitempty"`

	// shipment type
	ShipmentType MTOShipmentType `json:"shipmentType,omitempty"`

	// status
	Status MTOShipmentStatus `json:"status,omitempty"`
}

// Validate validates this update shipment
func (m *UpdateShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpmShipment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentType(formats); err != nil {
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

func (m *UpdateShipment) validateAgents(formats strfmt.Registry) error {
	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	if err := m.Agents.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("agents")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("agents")
		}
		return err
	}

	return nil
}

func (m *UpdateShipment) validateDestinationAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationAddress) { // not required
		return nil
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) validatePickupAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PickupAddress) { // not required
		return nil
	}

	if m.PickupAddress != nil {
		if err := m.PickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickupAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) validatePpmShipment(formats strfmt.Registry) error {
	if swag.IsZero(m.PpmShipment) { // not required
		return nil
	}

	if m.PpmShipment != nil {
		if err := m.PpmShipment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ppmShipment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ppmShipment")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) validateRequestedDeliveryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedDeliveryDate", "body", "date", m.RequestedDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateShipment) validateRequestedPickupDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateShipment) validateSecondaryDeliveryAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryDeliveryAddress) { // not required
		return nil
	}

	if m.SecondaryDeliveryAddress != nil {
		if err := m.SecondaryDeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryDeliveryAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryDeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) validateSecondaryPickupAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryPickupAddress) { // not required
		return nil
	}

	if m.SecondaryPickupAddress != nil {
		if err := m.SecondaryPickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryPickupAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryPickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) validateShipmentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentType) { // not required
		return nil
	}

	if err := m.ShipmentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

func (m *UpdateShipment) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this update shipment based on the context it is used
func (m *UpdateShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestinationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePickupAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePpmShipment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryDeliveryAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryPickupAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateShipment) contextValidateAgents(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Agents.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("agents")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("agents")
		}
		return err
	}

	return nil
}

func (m *UpdateShipment) contextValidateDestinationAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationAddress != nil {

		if swag.IsZero(m.DestinationAddress) { // not required
			return nil
		}

		if err := m.DestinationAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) contextValidatePickupAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.PickupAddress != nil {

		if swag.IsZero(m.PickupAddress) { // not required
			return nil
		}

		if err := m.PickupAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickupAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) contextValidatePpmShipment(ctx context.Context, formats strfmt.Registry) error {

	if m.PpmShipment != nil {

		if swag.IsZero(m.PpmShipment) { // not required
			return nil
		}

		if err := m.PpmShipment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ppmShipment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ppmShipment")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) contextValidateSecondaryDeliveryAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.SecondaryDeliveryAddress != nil {

		if swag.IsZero(m.SecondaryDeliveryAddress) { // not required
			return nil
		}

		if err := m.SecondaryDeliveryAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryDeliveryAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryDeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) contextValidateSecondaryPickupAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.SecondaryPickupAddress != nil {

		if swag.IsZero(m.SecondaryPickupAddress) { // not required
			return nil
		}

		if err := m.SecondaryPickupAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryPickupAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryPickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateShipment) contextValidateShipmentType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ShipmentType) { // not required
		return nil
	}

	if err := m.ShipmentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

func (m *UpdateShipment) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UpdateShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateShipment) UnmarshalBinary(b []byte) error {
	var res UpdateShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
