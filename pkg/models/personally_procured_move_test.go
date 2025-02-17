// RA Summary: gosec - errcheck - Unchecked return value
// RA: Linter flags errcheck error: Ignoring a method's return value can cause the program to overlook unexpected states and conditions.
// RA: Functions with unchecked return values in the file are used to generate stub data for a localized version of the application.
// RA: Given the data is being generated for local use and does not contain any sensitive information, there are no unexpected states and conditions
// RA: in which this would be considered a risk
// RA Developer Status: Mitigated
// RA Validator Status: Mitigated
// RA Modified Severity: N/A
// nolint:errcheck
package models_test

import (
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/factory"
	. "github.com/transcom/mymove/pkg/models"
)

func (suite *ModelSuite) TestPPMValidation() {
	ppm := &PersonallyProcuredMove{}

	expErrors := map[string][]string{
		"status": {"Status can not be blank."},
	}

	suite.verifyValidationErrors(ppm, expErrors)
}

func (suite *ModelSuite) TestPPMAdvance() {

	move := factory.BuildMove(suite.DB(), nil, nil)
	serviceMember := move.Orders.ServiceMember

	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)

	ppm, verrs, err := move.CreatePPM(suite.DB(), nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.NoError(err)
	suite.False(verrs.HasAny())

	advance.Request()
	SavePersonallyProcuredMove(suite.DB(), ppm)
	session := auth.Session{
		UserID:          serviceMember.User.ID,
		ApplicationName: auth.MilApp,
		ServiceMemberID: serviceMember.ID,
	}
	fetchedPPM, err := FetchPersonallyProcuredMove(suite.DB(), &session, ppm.ID)
	suite.NoError(err)
	suite.Equal(fetchedPPM.Advance.Status, ReimbursementStatusREQUESTED, "expected Requested")
}

// TODO: Fix test now that we capture transaction error
/* func (suite *ModelSuite) TestPPMAdvanceNoGTCC() {
	move := factory.BuildMove(suite.DB(), nil, nil)

	advance := BuildDraftReimbursement(1000, MethodOfReceiptGTCC)

	_, verrs, err := move.CreatePPM(suite.DB(), nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.NoError(err)
	suite.True(verrs.HasAny())
} */

func (suite *ModelSuite) TestPPMStateMachine() {
	orders := factory.BuildOrder(suite.DB(), nil, nil)
	orders.Status = OrderStatusSUBMITTED // NEVER do this outside of a test.
	suite.MustSave(&orders)
	factory.FetchOrBuildDefaultContractor(suite.DB(), nil, nil)

	moveOptions := MoveOptions{
		Show: BoolPointer(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	move.Orders = orders

	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)

	ppm, verrs, err := move.CreatePPM(suite.DB(), nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.NoError(err)
	suite.False(verrs.HasAny())

	ppm.Status = PPMStatusSUBMITTED // NEVER do this outside of a test.

	// Can cancel ppm
	err = ppm.Cancel()
	suite.NoError(err)
	suite.Equal(PPMStatusCANCELED, ppm.Status, "expected Canceled")
}

func (suite *ModelSuite) TestFetchPersonallyProcuredMoveByOrderID() {
	orderID := uuid.Must(uuid.NewV4())
	moveID, _ := uuid.FromString("7112b18b-7e03-4b28-adde-532b541bba8d")
	invalidID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")

	move := factory.BuildMove(suite.DB(), []factory.Customization{
		{
			Model: Move{
				ID: moveID,
			},
		},
		{
			Model: Order{
				ID: orderID,
			},
		},
	}, nil)

	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)

	ppm, verrs, err := move.CreatePPM(suite.DB(), nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.NoError(err)
	suite.False(verrs.HasAny())

	advance.Request()
	SavePersonallyProcuredMove(suite.DB(), ppm)

	tests := []struct {
		lookupID  uuid.UUID
		resultID  uuid.UUID
		resultErr bool
	}{
		{lookupID: orderID, resultID: moveID, resultErr: false},
		{lookupID: invalidID, resultID: invalidID, resultErr: true},
	}

	for _, ts := range tests {
		ppm, err := FetchPersonallyProcuredMoveByOrderID(suite.DB(), ts.lookupID)
		if ts.resultErr {
			suite.Error(err)
		} else {
			suite.NoError(err)
		}
		suite.Equal(ppm.MoveID, ts.resultID, "Wrong moveID: %s", ts.lookupID)
	}
}
