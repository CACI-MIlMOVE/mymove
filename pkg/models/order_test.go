package models_test

import (
	"time"

	"github.com/gobuffalo/uuid"

	"github.com/transcom/mymove/pkg/app"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) TestBasicOrderInstantiation() {
	order := &Order{}

	expErrors := map[string][]string{
		"orders_type":         {"OrdersType can not be blank."},
		"issue_date":          {"IssueDate can not be blank."},
		"report_by_date":      {"ReportByDate can not be blank."},
		"service_member_id":   {"ServiceMemberID can not be blank."},
		"new_duty_station_id": {"NewDutyStationID can not be blank."},
	}

	suite.verifyValidationErrors(order, expErrors)
}

func (suite *ModelSuite) TestFetchOrder() {

	serviceMember1, _ := testdatagen.MakeServiceMember(suite.db)
	serviceMember2, _ := testdatagen.MakeServiceMember(suite.db)
	reqApp := app.MyApp

	dutyStation := testdatagen.MakeAnyDutyStation(suite.db)
	issueDate := time.Date(2018, time.March, 10, 0, 0, 0, 0, time.UTC)
	reportByDate := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)
	ordersType := internalmessages.OrdersTypeBLUEBARK
	hasDependents := true
	uploadedOrder := Document{
		ServiceMember:   serviceMember1,
		ServiceMemberID: serviceMember1.ID,
		Name:            UploadedOrdersDocumentName,
	}
	suite.mustSave(&uploadedOrder)
	order := Order{
		ServiceMemberID:  serviceMember1.ID,
		ServiceMember:    serviceMember1,
		IssueDate:        issueDate,
		ReportByDate:     reportByDate,
		OrdersType:       ordersType,
		HasDependents:    hasDependents,
		NewDutyStationID: dutyStation.ID,
		NewDutyStation:   dutyStation,
		UploadedOrdersID: uploadedOrder.ID,
		UploadedOrders:   uploadedOrder,
	}
	suite.mustSave(&order)

	// User is authorized to fetch order
	goodOrder, err := FetchOrder(suite.db, serviceMember1.User, reqApp, order.ID)
	if suite.NoError(err) {
		suite.True(order.IssueDate.Equal(goodOrder.IssueDate))
		suite.True(order.ReportByDate.Equal(goodOrder.ReportByDate))
		suite.Equal(order.OrdersType, goodOrder.OrdersType)
		suite.Equal(order.HasDependents, goodOrder.HasDependents)
		suite.Equal(order.NewDutyStation.ID, goodOrder.NewDutyStation.ID)
	}

	// User is forbidden from fetching order
	_, err = FetchOrder(suite.db, serviceMember2.User, reqApp, order.ID)
	if suite.Error(err) {
		suite.Equal(ErrFetchForbidden, err)
	}

	// Wrong Order ID
	wrongID, _ := uuid.NewV4()
	_, err = FetchOrder(suite.db, serviceMember1.User, reqApp, wrongID)
	if suite.Error(err) {
		suite.Equal(ErrFetchNotFound, err)
	}
}
