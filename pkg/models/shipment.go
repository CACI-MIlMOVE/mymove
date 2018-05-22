package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Shipment represents a single shipment within a Service Member's move.
// PickupDate: when the shipment is currently scheduled to be picked up by the TSP
// RequestedPickupDate: when the shipment was originally scheduled to be picked up
// DeliveryDate: when the shipment is to be delivered
// BookDate: when the shipment was most recently offered to a TSP
type Shipment struct {
	ID                        uuid.UUID `json:"id" db:"id"`
	CreatedAt                 time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at" db:"updated_at"`
	PickupDate                time.Time `json:"pickup_date" db:"pickup_date"`
	RequestedPickupDate       time.Time `json:"requested_pickup_date" db:"requested_pickup_date"`
	DeliveryDate              time.Time `json:"delivery_date" db:"delivery_date"`
	BookDate                  time.Time `json:"book_date" db:"book_date"`
	TrafficDistributionListID uuid.UUID `json:"traffic_distribution_list_id" db:"traffic_distribution_list_id"`
	SourceGBLOC               string    `json:"source_gbloc" db:"source_gbloc"`
	Market                    *string   `json:"market" db:"market"`
}

// ShipmentWithOffer represents a single offered shipment within a Service Member's move.
type ShipmentWithOffer struct {
	ID                              uuid.UUID  `db:"id"`
	CreatedAt                       time.Time  `db:"created_at"`
	UpdatedAt                       time.Time  `db:"updated_at"`
	BookDate                        time.Time  `db:"book_date"`
	PickupDate                      time.Time  `db:"pickup_date"`
	RequestedPickupDate             time.Time  `db:"requested_pickup_date"`
	TrafficDistributionListID       uuid.UUID  `db:"traffic_distribution_list_id"`
	TransportationServiceProviderID *uuid.UUID `db:"transportation_service_provider_id"`
	SourceGBLOC                     *string    `db:"source_gbloc"`
	Market                          *string    `db:"market"`
	Accepted                        *bool      `db:"accepted"`
	RejectionReason                 *string    `db:"rejection_reason"`
	AdministrativeShipment          *bool      `db:"administrative_shipment"`
}

// FetchShipments looks up all shipments joined with their offer information in a
// ShipmentWithOffer struct. Optionally, you can only query for unassigned
// shipments with the `onlyUnassigned` parameter.
func FetchShipments(dbConnection *pop.Connection, onlyUnassigned bool) ([]ShipmentWithOffer, error) {
	shipments := []ShipmentWithOffer{}

	var sql string

	if onlyUnassigned {
		sql = `SELECT
				shipments.id,
				shipments.created_at,
				shipments.updated_at,
				shipments.pickup_date,
				shipments.requested_pickup_date,
				shipments.book_date,
				shipments.traffic_distribution_list_id,
				shipments.source_gbloc,
				shipments.market,
				shipment_offers.transportation_service_provider_id,
				shipment_offers.administrative_shipment
			FROM shipments
			LEFT JOIN shipment_offers ON
				shipment_offers.shipment_id=shipments.id
			WHERE shipment_offers.id IS NULL`
	} else {
		sql = `SELECT
				shipments.id,
				shipments.created_at,
				shipments.updated_at,
				shipments.pickup_date,
				shipments.requested_pickup_date,
				shipments.book_date,
				shipments.traffic_distribution_list_id,
				shipments.source_gbloc,
				shipments.market,
				shipment_offers.transportation_service_provider_id,
				shipment_offers.administrative_shipment
			FROM shipments
			LEFT JOIN shipment_offers ON
				shipment_offers.shipment_id=shipments.id`
	}

	err := dbConnection.RawQuery(sql).All(&shipments)

	return shipments, err
}

// Shipments is not required by pop and may be deleted
type Shipments []Shipment

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (s *Shipment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: s.TrafficDistributionListID, Name: "traffic_distribution_list_id"},
		&validators.StringIsPresent{Field: s.SourceGBLOC, Name: "source_gbloc"},
	), nil
}
