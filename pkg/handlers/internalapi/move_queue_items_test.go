package internalapi

import (
	"fmt"
	"net/http/httptest"

	"github.com/transcom/mymove/pkg/factory"
	queueop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/queues"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/models/roles"
	"github.com/transcom/mymove/pkg/testdatagen"
)

var statusToQueueMap = map[string]string{
	"SUBMITTED":         "new",
	"APPROVED":          "ppm_approved",
	"PAYMENT_REQUESTED": "ppm_payment_requested",
	"COMPLETED":         "ppm_completed",
}

func (suite *HandlerSuite) TestShowQueueHandler() {
	for status, queueType := range statusToQueueMap {
		var ppmStatus models.PPMStatus

		switch status {
		case "COMPLETED":
			ppmStatus = models.PPMStatusCOMPLETED
		case "APPROVED":
			ppmStatus = models.PPMStatusAPPROVED
		case "PAYMENT_REQUESTED":
			ppmStatus = models.PPMStatusPAYMENTREQUESTED
		case "SUBMITTED":
			ppmStatus = models.PPMStatusSUBMITTED
		}

		// Given: An office user
		officeUser := factory.BuildOfficeUserWithRoles(suite.DB(), nil, []roles.RoleType{roles.RoleTypeTOO})

		// Make a PPM
		moveShow := true
		ppm := testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
			Move: models.Move{
				Status: models.MoveStatus(status),
				Show:   &moveShow,
			},
			PersonallyProcuredMove: models.PersonallyProcuredMove{
				Status: ppmStatus,
			},
		})

		// And: the context contains the auth values
		path := "/queues/" + queueType
		req := httptest.NewRequest("GET", path, nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)

		params := queueop.ShowQueueParams{
			HTTPRequest: req,
			QueueType:   queueType,
		}

		// And: show Queue is queried
		showHandler := ShowQueueHandler{suite.HandlerConfig()}
		showResponse := showHandler.Handle(params)

		// Then: Expect a 200 status code
		okResponse := showResponse.(*queueop.ShowQueueOK)
		moveQueueItem := okResponse.Payload[0]

		// And: Returned query to include our added move
		// The moveQueueItems are produced by joining Moves, Orders and ServiceMember to each other, so we check the
		// furthest link in that chain
		serviceMember := ppm.Move.Orders.ServiceMember
		expectedCustomerName := fmt.Sprintf("%v, %v", *serviceMember.LastName, *serviceMember.FirstName)
		suite.Equal(expectedCustomerName, *moveQueueItem.CustomerName)
		suite.Equal(string(ppmStatus), *moveQueueItem.Status)
	}
}

func (suite *HandlerSuite) TestShowQueueHandlerForbidden() {
	for _, queueType := range statusToQueueMap {

		// Given: A non-office user
		user := factory.BuildServiceMember(suite.DB(), nil, nil)

		// And: the context contains the auth values
		path := "/queues/" + queueType
		req := httptest.NewRequest("GET", path, nil)
		req = suite.AuthenticateRequest(req, user)

		params := queueop.ShowQueueParams{
			HTTPRequest: req,
			QueueType:   queueType,
		}

		// And: show Queue is queried
		showHandler := ShowQueueHandler{suite.HandlerConfig()}
		showResponse := showHandler.Handle(params)

		// Then: Expect a 403 status code
		suite.Assertions.IsType(&queueop.ShowQueueForbidden{}, showResponse)
	}
}

func (suite *HandlerSuite) TestShowQueueHandlerNotFound() {

	// Given: An office user
	officeUser := factory.BuildOfficeUserWithRoles(suite.DB(), nil, []roles.RoleType{roles.RoleTypeTOO})

	// And: the context contains the auth values
	queueType := "queue_not_found"
	path := "/queues/" + queueType
	req := httptest.NewRequest("GET", path, nil)
	req = suite.AuthenticateOfficeRequest(req, officeUser)

	params := queueop.ShowQueueParams{
		HTTPRequest: req,
		QueueType:   queueType,
	}
	// And: show Queue is queried
	showHandler := ShowQueueHandler{suite.HandlerConfig()}
	showResponse := showHandler.Handle(params)

	// Then: Expect a 404 status code
	suite.CheckResponseNotFound(showResponse)
}
