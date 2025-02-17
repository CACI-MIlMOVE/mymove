// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PersonallyProcuredMovePayload personally procured move payload
//
// swagger:model PersonallyProcuredMovePayload
type PersonallyProcuredMovePayload struct {

	// When did you actually move?
	// Example: 2018-04-26
	// Format: date
	ActualMoveDate *strfmt.Date `json:"actual_move_date,omitempty"`

	// ZIP code
	// Example: 90210
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	AdditionalPickupPostalCode *string `json:"additional_pickup_postal_code,omitempty"`

	// advance
	Advance *Reimbursement `json:"advance,omitempty"`

	// advance worksheet
	AdvanceWorksheet *Document `json:"advance_worksheet,omitempty"`

	// When was the ppm move approved?
	// Example: 2019-03-26T13:19:56-04:00
	// Format: date-time
	ApproveDate *strfmt.DateTime `json:"approve_date,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// How many days of storage do you think you'll need?
	// Maximum: 90
	// Minimum: 0
	DaysInStorage *int64 `json:"days_in_storage,omitempty"`

	// ZIP code
	// Example: 90210
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	DestinationPostalCode *string `json:"destination_postal_code,omitempty"`

	// Estimated Storage Reimbursement
	EstimatedStorageReimbursement *string `json:"estimated_storage_reimbursement,omitempty"`

	// Will you move anything from another pickup location?
	HasAdditionalPostalCode *bool `json:"has_additional_postal_code,omitempty"`

	// Has Pro-Gear
	// Enum: [NOT SURE YES NO]
	HasProGear *string `json:"has_pro_gear,omitempty"`

	// Has Pro-Gear Over Thousand Pounds
	// Enum: [NOT SURE YES NO]
	HasProGearOverThousand *string `json:"has_pro_gear_over_thousand,omitempty"`

	// Would you like an advance of up to 60% of your PPM incentive?
	HasRequestedAdvance *bool `json:"has_requested_advance,omitempty"`

	// Will you put anything in storage?
	HasSit *bool `json:"has_sit,omitempty"`

	// id
	// Example: c56a4180-65aa-42ec-a945-5fd21dec0538
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Estimated incentive maximum in cents
	IncentiveEstimateMax *int64 `json:"incentive_estimate_max,omitempty"`

	// Estimated incentive minimum in cents
	IncentiveEstimateMin *int64 `json:"incentive_estimate_min,omitempty"`

	// Distance between origin and destination in miles
	Mileage *int64 `json:"mileage,omitempty"`

	// move id
	// Example: c56a4180-65aa-42ec-a945-5fd21dec0538
	// Format: uuid
	MoveID strfmt.UUID `json:"move_id,omitempty"`

	// Net Weight
	// Minimum: 1
	NetWeight *int64 `json:"net_weight,omitempty"`

	// When do you plan to move?
	// Example: 2018-04-26
	// Format: date
	OriginalMoveDate *strfmt.Date `json:"original_move_date,omitempty"`

	// ZIP code
	// Example: 90210
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	PickupPostalCode *string `json:"pickup_postal_code,omitempty"`

	// Maximum SIT reimbursement for the planned SIT duration
	PlannedSitMax *int64 `json:"planned_sit_max,omitempty"`

	// Maximum SIT reimbursement for maximum SIT duration
	SitMax *int64 `json:"sit_max,omitempty"`

	// size
	Size *TShirtSize `json:"size,omitempty"`

	// status
	Status PPMStatus `json:"status,omitempty"`

	// When was the ppm move submitted?
	// Example: 2019-03-26T13:19:56-04:00
	// Format: date-time
	SubmitDate *strfmt.DateTime `json:"submit_date,omitempty"`

	// Total cost for the planned SIT duration
	TotalSitCost *int64 `json:"total_sit_cost,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// Weight Estimate
	// Minimum: 0
	WeightEstimate *int64 `json:"weight_estimate,omitempty"`
}

// Validate validates this personally procured move payload
func (m *PersonallyProcuredMovePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalPickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvanceWorksheet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApproveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaysInStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasProGear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasProGearOverThousand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmitDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightEstimate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonallyProcuredMovePayload) validateActualMoveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ActualMoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("actual_move_date", "body", "date", m.ActualMoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateAdditionalPickupPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalPickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("additional_pickup_postal_code", "body", *m.AdditionalPickupPostalCode, `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateAdvance(formats strfmt.Registry) error {
	if swag.IsZero(m.Advance) { // not required
		return nil
	}

	if m.Advance != nil {
		if err := m.Advance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advance")
			}
			return err
		}
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateAdvanceWorksheet(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvanceWorksheet) { // not required
		return nil
	}

	if m.AdvanceWorksheet != nil {
		if err := m.AdvanceWorksheet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advance_worksheet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advance_worksheet")
			}
			return err
		}
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateApproveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ApproveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("approve_date", "body", "date-time", m.ApproveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateDaysInStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.DaysInStorage) { // not required
		return nil
	}

	if err := validate.MinimumInt("days_in_storage", "body", *m.DaysInStorage, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("days_in_storage", "body", *m.DaysInStorage, 90, false); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateDestinationPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("destination_postal_code", "body", *m.DestinationPostalCode, `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

var personallyProcuredMovePayloadTypeHasProGearPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT SURE","YES","NO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		personallyProcuredMovePayloadTypeHasProGearPropEnum = append(personallyProcuredMovePayloadTypeHasProGearPropEnum, v)
	}
}

const (

	// PersonallyProcuredMovePayloadHasProGearNOTSURE captures enum value "NOT SURE"
	PersonallyProcuredMovePayloadHasProGearNOTSURE string = "NOT SURE"

	// PersonallyProcuredMovePayloadHasProGearYES captures enum value "YES"
	PersonallyProcuredMovePayloadHasProGearYES string = "YES"

	// PersonallyProcuredMovePayloadHasProGearNO captures enum value "NO"
	PersonallyProcuredMovePayloadHasProGearNO string = "NO"
)

// prop value enum
func (m *PersonallyProcuredMovePayload) validateHasProGearEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, personallyProcuredMovePayloadTypeHasProGearPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PersonallyProcuredMovePayload) validateHasProGear(formats strfmt.Registry) error {
	if swag.IsZero(m.HasProGear) { // not required
		return nil
	}

	// value enum
	if err := m.validateHasProGearEnum("has_pro_gear", "body", *m.HasProGear); err != nil {
		return err
	}

	return nil
}

var personallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT SURE","YES","NO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		personallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum = append(personallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum, v)
	}
}

const (

	// PersonallyProcuredMovePayloadHasProGearOverThousandNOTSURE captures enum value "NOT SURE"
	PersonallyProcuredMovePayloadHasProGearOverThousandNOTSURE string = "NOT SURE"

	// PersonallyProcuredMovePayloadHasProGearOverThousandYES captures enum value "YES"
	PersonallyProcuredMovePayloadHasProGearOverThousandYES string = "YES"

	// PersonallyProcuredMovePayloadHasProGearOverThousandNO captures enum value "NO"
	PersonallyProcuredMovePayloadHasProGearOverThousandNO string = "NO"
)

// prop value enum
func (m *PersonallyProcuredMovePayload) validateHasProGearOverThousandEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, personallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PersonallyProcuredMovePayload) validateHasProGearOverThousand(formats strfmt.Registry) error {
	if swag.IsZero(m.HasProGearOverThousand) { // not required
		return nil
	}

	// value enum
	if err := m.validateHasProGearOverThousandEnum("has_pro_gear_over_thousand", "body", *m.HasProGearOverThousand); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateMoveID(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveID) { // not required
		return nil
	}

	if err := validate.FormatOf("move_id", "body", "uuid", m.MoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateNetWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.NetWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("net_weight", "body", *m.NetWeight, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateOriginalMoveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalMoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("original_move_date", "body", "date", m.OriginalMoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validatePickupPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.PickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("pickup_postal_code", "body", *m.PickupPostalCode, `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateStatus(formats strfmt.Registry) error {
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

func (m *PersonallyProcuredMovePayload) validateSubmitDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SubmitDate) { // not required
		return nil
	}

	if err := validate.FormatOf("submit_date", "body", "date-time", m.SubmitDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) validateWeightEstimate(formats strfmt.Registry) error {
	if swag.IsZero(m.WeightEstimate) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight_estimate", "body", *m.WeightEstimate, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this personally procured move payload based on the context it is used
func (m *PersonallyProcuredMovePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdvanceWorksheet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
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

func (m *PersonallyProcuredMovePayload) contextValidateAdvance(ctx context.Context, formats strfmt.Registry) error {

	if m.Advance != nil {

		if swag.IsZero(m.Advance) { // not required
			return nil
		}

		if err := m.Advance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advance")
			}
			return err
		}
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) contextValidateAdvanceWorksheet(ctx context.Context, formats strfmt.Registry) error {

	if m.AdvanceWorksheet != nil {

		if swag.IsZero(m.AdvanceWorksheet) { // not required
			return nil
		}

		if err := m.AdvanceWorksheet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advance_worksheet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advance_worksheet")
			}
			return err
		}
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {

		if swag.IsZero(m.Size) { // not required
			return nil
		}

		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *PersonallyProcuredMovePayload) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PersonallyProcuredMovePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonallyProcuredMovePayload) UnmarshalBinary(b []byte) error {
	var res PersonallyProcuredMovePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
