package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Tariff400ngZip5RateArea repersents the mapping from a full 5-digit zipcode to a
// specific rate area. This is only needed for a small subset of zip3s.
type Tariff400ngZip5RateArea struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	RateArea  string    `json:"rate_area" db:"rate_area"`
	Zip5      string    `json:"zip5" db:"zip5"`
}

// String is not required by pop and may be deleted
func (t Tariff400ngZip5RateArea) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tariff400ngZip5RateAreas is not required by pop and may be deleted
type Tariff400ngZip5RateAreas []Tariff400ngZip5RateArea

// String is not required by pop and may be deleted
func (t Tariff400ngZip5RateAreas) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (t *Tariff400ngZip5RateArea) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.RateArea, Name: "RateArea"},
		&validators.RegexMatch{Field: t.RateArea, Name: "RateArea", Expr: "^(ZIP|US[0-9]+)$"},
		&validators.StringLengthInRange{Field: t.Zip5, Name: "Zip5", Min: 5, Max: 5},
	), nil
}
