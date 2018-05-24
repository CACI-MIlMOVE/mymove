package handlers

import (
	"fmt"
	"net/http/httptest"
	"time"

	"github.com/gobuffalo/uuid"

	ordersop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/orders"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestCreateOrder() {
	sm, _ := testdatagen.MakeServiceMember(suite.db)
	station := testdatagen.MakeAnyDutyStation(suite.db)

	req := httptest.NewRequest("POST", "/orders", nil)
	req = suite.authenticateRequest(req, sm)

	hasDependents := true
	issueDate := time.Date(2018, time.March, 10, 0, 0, 0, 0, time.UTC)
	reportByDate := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)
	ordersType := internalmessages.OrdersTypePERMANENTCHANGEOFSTATION
	payload := &internalmessages.CreateUpdateOrders{
		HasDependents:    fmtBool(hasDependents),
		IssueDate:        fmtDate(issueDate),
		ReportByDate:     fmtDate(reportByDate),
		OrdersType:       ordersType,
		NewDutyStationID: fmtUUID(station.ID),
		ServiceMemberID:  fmtUUID(sm.ID),
	}

	params := ordersop.CreateOrdersParams{
		HTTPRequest:  req,
		CreateOrders: payload,
	}
	createHandler := CreateOrdersHandler(NewHandlerContext(suite.db, suite.logger))
	response := createHandler.Handle(params)

	suite.Assertions.IsType(&ordersop.CreateOrdersCreated{}, response)
	okResponse := response.(*ordersop.CreateOrdersCreated)

	suite.Assertions.Equal(sm.ID.String(), okResponse.Payload.ServiceMemberID.String())
	suite.Assertions.Equal(ordersType, okResponse.Payload.OrdersType)
}

func (suite *HandlerSuite) TestShowOrder() {
	order, _ := testdatagen.MakeOrder(suite.db)

	path := fmt.Sprintf("/orders/%v", order.ID.String())
	req := httptest.NewRequest("GET", path, nil)
	req = suite.authenticateRequest(req, order.ServiceMember)

	params := ordersop.ShowOrdersParams{
		HTTPRequest: req,
		OrdersID:    *fmtUUID(order.ID),
	}
	showHandler := ShowOrdersHandler(NewHandlerContext(suite.db, suite.logger))
	response := showHandler.Handle(params)

	suite.Assertions.IsType(&ordersop.ShowOrdersOK{}, response)
	okResponse := response.(*ordersop.ShowOrdersOK)

	suite.Assertions.Equal(order.ServiceMember.ID.String(), okResponse.Payload.ServiceMemberID.String())
	suite.Assertions.Equal(order.OrdersType, okResponse.Payload.OrdersType)
}

func (suite *HandlerSuite) TestUpdateOrder() {
	order, _ := testdatagen.MakeOrder(suite.db)

	path := fmt.Sprintf("/orders/%v", order.ID.String())
	req := httptest.NewRequest("PUT", path, nil)
	req = suite.authenticateRequest(req, order.ServiceMember)

	newOrdersType := internalmessages.OrdersTypePERMANENTCHANGEOFSTATION
	newOrdersTypeDetail := internalmessages.OrdersTypeDetailHHGPERMITTED
	departmentIndicator := internalmessages.DeptIndicatorAIRFORCE
	otherServiceMemberUUID := uuid.Must(uuid.NewV4())

	payload := &internalmessages.CreateUpdateOrders{
		OrdersNumber:        fmtString("123456"),
		HasDependents:       fmtBool(order.HasDependents),
		IssueDate:           fmtDate(order.IssueDate),
		ReportByDate:        fmtDate(order.ReportByDate),
		OrdersType:          newOrdersType,
		OrdersTypeDetail:    &newOrdersTypeDetail,
		NewDutyStationID:    fmtUUID(order.NewDutyStationID),
		Tac:                 order.TAC,
		DepartmentIndicator: &departmentIndicator,
		// Attempt to assign to another service member
		ServiceMemberID: fmtUUID(otherServiceMemberUUID),
	}

	params := ordersop.UpdateOrdersParams{
		HTTPRequest:  req,
		OrdersID:     *fmtUUID(order.ID),
		UpdateOrders: payload,
	}
	updateHandler := UpdateOrdersHandler(NewHandlerContext(suite.db, suite.logger))
	response := updateHandler.Handle(params)

	suite.Assertions.IsType(&ordersop.UpdateOrdersOK{}, response)
	okResponse := response.(*ordersop.UpdateOrdersOK)

	suite.Assertions.Equal(fmtString("123456"), okResponse.Payload.OrdersNumber)
	suite.Assertions.Equal(order.ServiceMember.ID.String(), okResponse.Payload.ServiceMemberID.String(), "service member id should not change")
	suite.Assertions.Equal(newOrdersType, okResponse.Payload.OrdersType)
	suite.Assertions.Equal(newOrdersTypeDetail, *okResponse.Payload.OrdersTypeDetail)
}
