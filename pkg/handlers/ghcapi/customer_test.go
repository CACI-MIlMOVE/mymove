package ghcapi

import (
	"net/http/httptest"

	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/factory"
	customerops "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/customer"
	"github.com/transcom/mymove/pkg/gen/ghcmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models/roles"
	customerservice "github.com/transcom/mymove/pkg/services/office_user/customer"
)

func (suite *HandlerSuite) TestGetCustomerHandlerIntegration() {
	customer := factory.BuildServiceMember(suite.DB(), nil, nil)

	request := httptest.NewRequest("GET", "/customer/{customerID}", nil)
	params := customerops.GetCustomerParams{
		HTTPRequest: request,
		CustomerID:  strfmt.UUID(customer.ID.String()),
	}
	handlerConfig := suite.HandlerConfig()
	handler := GetCustomerHandler{
		handlerConfig,
		customerservice.NewCustomerFetcher(),
	}

	// Validate incoming payload: no body to validate

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&customerops.GetCustomerOK{}, response)
	getCustomerResponse := response.(*customerops.GetCustomerOK)
	getCustomerPayload := getCustomerResponse.Payload

	// Validate outgoing payload
	suite.NoError(getCustomerPayload.Validate(strfmt.Default))

	suite.Equal(strfmt.UUID(customer.ID.String()), getCustomerPayload.ID)
	suite.Equal(*customer.Edipi, getCustomerPayload.DodID)
	suite.Equal(strfmt.UUID(customer.UserID.String()), getCustomerPayload.UserID)
	suite.Equal(customer.Affiliation.String(), getCustomerPayload.Agency)
	suite.Equal(customer.PersonalEmail, getCustomerPayload.Email)
	suite.Equal(customer.Telephone, getCustomerPayload.Phone)
	suite.NotZero(getCustomerPayload.CurrentAddress)
}

func (suite *HandlerSuite) TestUpdateCustomerHandler() {
	officeUser := factory.BuildOfficeUserWithRoles(suite.DB(), nil, []roles.RoleType{roles.RoleTypeTOO})
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeServicesCounselor,
	})

	body := &ghcmessages.UpdateCustomerPayload{
		LastName:  "Newlastname",
		FirstName: "Newfirstname",
		Phone:     handlers.FmtString("223-455-3399"),
		BackupContact: &ghcmessages.BackupContact{
			Name:  handlers.FmtString("New Backup Contact"),
			Phone: handlers.FmtString("445-345-1212"),
			Email: handlers.FmtString("newbackup@mail.com"),
		},
	}
	currentAddress := ghcmessages.Address{
		StreetAddress1: handlers.FmtString("123 New Street"),
		City:           handlers.FmtString("Newcity"),
		State:          handlers.FmtString("MA"),
		PostalCode:     handlers.FmtString("12345"),
	}
	body.CurrentAddress.Address = currentAddress

	customer := factory.BuildExtendedServiceMember(suite.DB(), nil, nil)
	request := httptest.NewRequest("PATCH", "/orders/{customerID}", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := customerops.UpdateCustomerParams{
		HTTPRequest: request,
		CustomerID:  strfmt.UUID(customer.ID.String()),
		IfMatch:     etag.GenerateEtag(customer.UpdatedAt),
		Body:        body,
	}
	handlerConfig := suite.HandlerConfig()
	handler := UpdateCustomerHandler{
		handlerConfig,
		customerservice.NewCustomerUpdater(),
	}

	// Validate incoming payload
	suite.NoError(params.Body.Validate(strfmt.Default))

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	// TODO: test with actual updated customer?
	// updatedCustomer, _ := models.FetchServiceMember(suite.DB(), customer.ID)
	suite.Assertions.IsType(&customerops.UpdateCustomerOK{}, response)
	updateCustomerResponse := response.(*customerops.UpdateCustomerOK)
	updateCustomerPayload := updateCustomerResponse.Payload

	// Validate outgoing payload
	suite.NoError(updateCustomerPayload.Validate(strfmt.Default))

	suite.Equal(body.FirstName, updateCustomerPayload.FirstName)
	suite.Equal(body.LastName, updateCustomerPayload.LastName)
	suite.Equal(body.Phone, updateCustomerPayload.Phone)
	suite.Equal(body.CurrentAddress.StreetAddress1, updateCustomerPayload.CurrentAddress.StreetAddress1)
	suite.Equal(body.CurrentAddress.City, updateCustomerPayload.CurrentAddress.City)
	suite.Equal(body.CurrentAddress.PostalCode, updateCustomerPayload.CurrentAddress.PostalCode)
	suite.Equal(body.CurrentAddress.State, updateCustomerPayload.CurrentAddress.State)
	suite.Equal(body.BackupContact.Name, updateCustomerPayload.BackupContact.Name)
	suite.Equal(body.BackupContact.Phone, updateCustomerPayload.BackupContact.Phone)
	suite.Equal(body.BackupContact.Email, updateCustomerPayload.BackupContact.Email)
}
