// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GBLOC g b l o c
//
// swagger:model GBLOC
type GBLOC string

func NewGBLOC(value GBLOC) *GBLOC {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GBLOC.
func (m GBLOC) Pointer() *GBLOC {
	return &m
}

const (

	// GBLOCAGFM captures enum value "AGFM"
	GBLOCAGFM GBLOC = "AGFM"

	// GBLOCAPAT captures enum value "APAT"
	GBLOCAPAT GBLOC = "APAT"

	// GBLOCBGAC captures enum value "BGAC"
	GBLOCBGAC GBLOC = "BGAC"

	// GBLOCBGNC captures enum value "BGNC"
	GBLOCBGNC GBLOC = "BGNC"

	// GBLOCBKAS captures enum value "BKAS"
	GBLOCBKAS GBLOC = "BKAS"

	// GBLOCCFMQ captures enum value "CFMQ"
	GBLOCCFMQ GBLOC = "CFMQ"

	// GBLOCCLPK captures enum value "CLPK"
	GBLOCCLPK GBLOC = "CLPK"

	// GBLOCCNNQ captures enum value "CNNQ"
	GBLOCCNNQ GBLOC = "CNNQ"

	// GBLOCDMAT captures enum value "DMAT"
	GBLOCDMAT GBLOC = "DMAT"

	// GBLOCGSAT captures enum value "GSAT"
	GBLOCGSAT GBLOC = "GSAT"

	// GBLOCHAFC captures enum value "HAFC"
	GBLOCHAFC GBLOC = "HAFC"

	// GBLOCHBAT captures enum value "HBAT"
	GBLOCHBAT GBLOC = "HBAT"

	// GBLOCJEAT captures enum value "JEAT"
	GBLOCJEAT GBLOC = "JEAT"

	// GBLOCJENQ captures enum value "JENQ"
	GBLOCJENQ GBLOC = "JENQ"

	// GBLOCKKFA captures enum value "KKFA"
	GBLOCKKFA GBLOC = "KKFA"

	// GBLOCLHNQ captures enum value "LHNQ"
	GBLOCLHNQ GBLOC = "LHNQ"

	// GBLOCLKNQ captures enum value "LKNQ"
	GBLOCLKNQ GBLOC = "LKNQ"

	// GBLOCMAPK captures enum value "MAPK"
	GBLOCMAPK GBLOC = "MAPK"

	// GBLOCMAPS captures enum value "MAPS"
	GBLOCMAPS GBLOC = "MAPS"

	// GBLOCMBFL captures enum value "MBFL"
	GBLOCMBFL GBLOC = "MBFL"

	// GBLOCMLNQ captures enum value "MLNQ"
	GBLOCMLNQ GBLOC = "MLNQ"

	// GBLOCXXXX captures enum value "XXXX"
	GBLOCXXXX GBLOC = "XXXX"
)

// for schema
var gBLOCEnum []interface{}

func init() {
	var res []GBLOC
	if err := json.Unmarshal([]byte(`["AGFM","APAT","BGAC","BGNC","BKAS","CFMQ","CLPK","CNNQ","DMAT","GSAT","HAFC","HBAT","JEAT","JENQ","KKFA","LHNQ","LKNQ","MAPK","MAPS","MBFL","MLNQ","XXXX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gBLOCEnum = append(gBLOCEnum, v)
	}
}

func (m GBLOC) validateGBLOCEnum(path, location string, value GBLOC) error {
	if err := validate.EnumCase(path, location, value, gBLOCEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this g b l o c
func (m GBLOC) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGBLOCEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this g b l o c based on context it is used
func (m GBLOC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
