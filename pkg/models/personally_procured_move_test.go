package models_test

import (
	"github.com/pkg/errors"

	"github.com/transcom/mymove/pkg/auth"
	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) TestPPMValidation() {
	ppm := &PersonallyProcuredMove{}

	expErrors := map[string][]string{
		"status": {"Status can not be blank."},
	}

	suite.verifyValidationErrors(ppm, expErrors)
}

func (suite *ModelSuite) TestPPMAdvance() {

	move, _ := testdatagen.MakeMove(suite.db)
	serviceMember := move.Orders.ServiceMember

	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)

	ppm, verrs, err := move.CreatePPM(suite.db, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.Nil(err)
	suite.False(verrs.HasAny())

	advance.Request()
	SavePersonallyProcuredMove(suite.db, ppm)
	session := auth.Session{
		UserID:          serviceMember.User.ID,
		ApplicationName: auth.MyApp,
		ServiceMemberID: serviceMember.ID,
	}
	fetchedPPM, err := FetchPersonallyProcuredMove(suite.db, &session, ppm.ID)
	suite.Nil(err)
	suite.Equal(fetchedPPM.Advance.Status, ReimbursementStatusREQUESTED, "expected Requested")
}

func (suite *ModelSuite) TestPPMAdvanceNoGTCC() {
	move, _ := testdatagen.MakeMove(suite.db)

	advance := BuildDraftReimbursement(1000, MethodOfReceiptGTCC)

	_, verrs, err := move.CreatePPM(suite.db, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.Nil(err)
	suite.True(verrs.HasAny())
}

func (suite *ModelSuite) TestPPMStateMachine() {
	move, _ := testdatagen.MakeMove(suite.db)

	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)

	ppm, verrs, err := move.CreatePPM(suite.db, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.Nil(err)
	suite.False(verrs.HasAny())

	// Can't cancel a PPM with DRAFT status
	err = ppm.Cancel()
	suite.Equal(ErrInvalidTransition, errors.Cause(err))

	ppm.Status = PPMStatusSUBMITTED // NEVER do this outside of a test.

	// Can cancel ppm
	err = ppm.Cancel()
	suite.Nil(err)
	suite.Equal(ppm.Status, PPMStatusCANCELED, "expected Canceled")

	// RESET PPM
	ppm.Status = PPMStatusSUBMITTED // NEVER do this outside of a test.
	suite.Equal(ppm.Status, PPMStatusSUBMITTED, "expected Submitted")
	// When move is canceled, expect associated PPM to be canceled
	err = move.Submit()
	suite.Nil(err)
	err = move.Cancel()
	suite.Nil(err)
	// PPM has also been canceled
	suite.Equal(ppm.Status, PPMStatusCANCELED, "expected Canceled")
}
