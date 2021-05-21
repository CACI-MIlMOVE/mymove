package payloads

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
	"github.com/transcom/mymove/pkg/models"
)

// CustomerToServiceMember transforms UpdateCustomerPayload to ServiceMember model
func CustomerToServiceMember(payload ghcmessages.UpdateCustomerPayload) models.ServiceMember {

	var address = models.Address{
		ID:             uuid.FromStringOrNil(payload.CurrentAddress.ID.String()),
		StreetAddress1: *payload.CurrentAddress.StreetAddress1,
		StreetAddress2: payload.CurrentAddress.StreetAddress2,
		StreetAddress3: payload.CurrentAddress.StreetAddress3,
		City:           *payload.CurrentAddress.City,
		State:          *payload.CurrentAddress.State,
		PostalCode:     *payload.CurrentAddress.PostalCode,
		Country:        payload.CurrentAddress.Country,
	}

	var backupContact = models.BackupContact{
		Email: *payload.BackupContact.Email,
		Name:  *payload.BackupContact.Name,
		Phone: payload.BackupContact.Phone,
	}

	var backupContacts []models.BackupContact
	backupContacts = append(backupContacts, backupContact)

	return models.ServiceMember{
		ResidentialAddress: &address,
		BackupContacts:     backupContacts,
		FirstName:          &payload.FirstName,
		LastName:           &payload.LastName,
		PersonalEmail:      payload.Email,
		Telephone:          payload.Phone,
	}
}

// AddressModel model
func AddressModel(address *ghcmessages.Address) *models.Address {
	// To check if the model is intended to be blank, we'll look at both ID and StreetAddress1
	// We should always have ID if the user intends to update an Address,
	// and StreetAddress1 is a required field on creation. If both are blank, it should be treated as nil.
	var blankSwaggerID strfmt.UUID
	if address == nil || (address.ID == blankSwaggerID && address.StreetAddress1 == nil) {
		return nil
	}

	modelAddress := &models.Address{
		ID:             uuid.FromStringOrNil(address.ID.String()),
		StreetAddress2: address.StreetAddress2,
		StreetAddress3: address.StreetAddress3,
		Country:        address.Country,
	}
	if address.StreetAddress1 != nil {
		modelAddress.StreetAddress1 = *address.StreetAddress1
	}
	if address.City != nil {
		modelAddress.City = *address.City
	}
	if address.State != nil {
		modelAddress.State = *address.State
	}
	if address.PostalCode != nil {
		modelAddress.PostalCode = *address.PostalCode
	}
	return modelAddress
}

// MTOAgentModel model
func MTOAgentModel(mtoAgent *ghcmessages.MTOAgent) *models.MTOAgent {
	if mtoAgent == nil {
		return nil
	}

	return &models.MTOAgent{
		ID:            uuid.FromStringOrNil(mtoAgent.ID.String()),
		MTOShipmentID: uuid.FromStringOrNil(mtoAgent.MtoShipmentID.String()),
		FirstName:     mtoAgent.FirstName,
		LastName:      mtoAgent.LastName,
		Email:         mtoAgent.Email,
		Phone:         mtoAgent.Phone,
		MTOAgentType:  models.MTOAgentType(mtoAgent.AgentType),
	}
}

// MTOAgentsModel model
func MTOAgentsModel(mtoAgents *ghcmessages.MTOAgents) *models.MTOAgents {
	if mtoAgents == nil {
		return nil
	}

	agents := make(models.MTOAgents, len(*mtoAgents))

	for i, m := range *mtoAgents {
		agents[i] = *MTOAgentModel(m)
	}

	return &agents
}

// MTOShipmentModelFromCreate model
func MTOShipmentModelFromCreate(mtoShipment *ghcmessages.CreateMTOShipment) *models.MTOShipment {
	if mtoShipment == nil {
		return nil
	}

	model := &models.MTOShipment{
		MoveTaskOrderID:  uuid.FromStringOrNil(mtoShipment.MoveTaskOrderID.String()),
		ShipmentType:     models.MTOShipmentTypeHHG,
		Status:           models.MTOShipmentStatusSubmitted,
		CustomerRemarks:  mtoShipment.CustomerRemarks,
		CounselorRemarks: mtoShipment.CounselorRemarks,
	}

	if mtoShipment.RequestedPickupDate != nil {
		model.RequestedPickupDate = swag.Time(time.Time(*mtoShipment.RequestedPickupDate))
	}

	// Set up address models
	var addressModel *models.Address

	addressModel = AddressModel(&mtoShipment.PickupAddress.Address)
	if addressModel != nil {
		model.PickupAddress = addressModel
	}

	addressModel = AddressModel(&mtoShipment.DestinationAddress.Address)
	if addressModel != nil {
		model.DestinationAddress = addressModel
	}

	if mtoShipment.Agents != nil {
		model.MTOAgents = *MTOAgentsModel(&mtoShipment.Agents)
	}

	return model
}
