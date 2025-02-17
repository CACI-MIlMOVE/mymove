package evaluationreport

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/transcom/mymove/pkg/apperror"
	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/factory"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite EvaluationReportSuite) TestSubmitEvaluationReport() {
	updater := NewEvaluationReportUpdater()

	suite.Run("Successful submission", func() {
		// Create office user
		officeUser := factory.BuildOfficeUser(suite.DB(), nil, nil)
		// Create a report
		inspectionType := models.EvaluationReportInspectionTypeVirtual
		location := models.EvaluationReportLocationTypeOrigin
		inspectionTime := time.Now().AddDate(0, 0, -4)
		evalStart := inspectionTime
		evalEnd := inspectionTime

		evaluationReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model:    officeUser,
				LinkOnly: true,
			},
			{
				Model: models.EvaluationReport{
					InspectionDate:     models.TimePointer(time.Now()),
					InspectionType:     &inspectionType,
					Location:           &location,
					EvalStart:          &evalStart,
					EvalEnd:            &evalEnd,
					ViolationsObserved: models.BoolPointer(false),
					Remarks:            models.StringPointer("This is a remark."),
				},
			},
		}, nil)
		// Generate an etag
		eTag := etag.GenerateEtag(evaluationReport.UpdatedAt)
		// Submit the report
		err := updater.SubmitEvaluationReport(suite.AppContextForTest(), evaluationReport.ID, officeUser.ID, eTag)

		suite.NoError(err)
	})

	suite.Run("Bad etag", func() {
		// Create office user
		officeUser := factory.BuildOfficeUser(suite.DB(), nil, nil)
		// Create a report
		inspectionType := models.EvaluationReportInspectionTypeVirtual
		location := models.EvaluationReportLocationTypeOrigin
		evaluationReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model:    officeUser,
				LinkOnly: true,
			},
			{
				Model: models.EvaluationReport{
					InspectionDate:     models.TimePointer(time.Now()),
					InspectionType:     &inspectionType,
					Location:           &location,
					ViolationsObserved: models.BoolPointer(false),
					Remarks:            models.StringPointer("This is a remark."),
				},
			},
		}, nil)
		// Generate an etag
		eTag := ""
		// Submit the report
		err := updater.SubmitEvaluationReport(suite.AppContextForTest(), evaluationReport.ID, officeUser.ID, eTag)
		suite.IsType(apperror.PreconditionFailedError{}, err)
	})

	suite.Run("Missing inspection date", func() {
		// Create office user
		officeUser := factory.BuildOfficeUser(suite.DB(), nil, nil)
		// Create a report
		inspectionType := models.EvaluationReportInspectionTypeVirtual
		location := models.EvaluationReportLocationTypeOrigin
		inspectionTime := time.Now().AddDate(0, 0, -4)
		evalStart := inspectionTime
		evalEnd := inspectionTime

		// Missing inspection date
		evaluationReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model:    officeUser,
				LinkOnly: true,
			},
			{
				Model: models.EvaluationReport{
					InspectionType:     &inspectionType,
					Location:           &location,
					EvalStart:          &evalStart,
					EvalEnd:            &evalEnd,
					ViolationsObserved: models.BoolPointer(false),
					Remarks:            models.StringPointer("This is a remark."),
				},
			},
		}, nil)
		// Generate an etag
		eTag := etag.GenerateEtag(evaluationReport.UpdatedAt)
		// Submit the report
		err := updater.SubmitEvaluationReport(suite.AppContextForTest(), evaluationReport.ID, officeUser.ID, eTag)
		suite.Equal(models.ErrInvalidTransition, errors.Cause(err))
	})

	suite.Run("Missing location description for other location", func() {
		// Create office user
		officeUser := factory.BuildOfficeUser(suite.DB(), nil, nil)
		// Create a report
		inspectionType := models.EvaluationReportInspectionTypeVirtual
		location := models.EvaluationReportLocationTypeOther
		inspectionTime := time.Now().AddDate(0, 0, -4)
		evalStart := inspectionTime
		evalEnd := inspectionTime
		// Missing location description
		evaluationReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model:    officeUser,
				LinkOnly: true,
			},
			{
				Model: models.EvaluationReport{
					InspectionDate:     models.TimePointer(time.Now()),
					InspectionType:     &inspectionType,
					Location:           &location,
					EvalStart:          &evalStart,
					EvalEnd:            &evalEnd,
					ViolationsObserved: models.BoolPointer(false),
					Remarks:            models.StringPointer("This is a remark."),
				},
			},
		}, nil)
		// Generate an etag
		eTag := etag.GenerateEtag(evaluationReport.UpdatedAt)
		// Submit the report
		err := updater.SubmitEvaluationReport(suite.AppContextForTest(), evaluationReport.ID, officeUser.ID, eTag)
		suite.Equal(models.ErrInvalidTransition, errors.Cause(err))
	})

	suite.Run("Missing time depart for physical inspection", func() {
		// Create office user
		officeUser := factory.BuildOfficeUser(suite.DB(), nil, nil)
		// Create a report
		inspectionType := models.EvaluationReportInspectionTypePhysical
		location := models.EvaluationReportLocationTypeOther
		inspectionTime := time.Now().AddDate(0, 0, -4)
		evalStart := inspectionTime
		evalEnd := inspectionTime
		// Missing location description
		evaluationReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model:    officeUser,
				LinkOnly: true,
			},
			{
				Model: models.EvaluationReport{
					InspectionDate:     models.TimePointer(time.Now()),
					InspectionType:     &inspectionType,
					Location:           &location,
					EvalStart:          &evalStart,
					EvalEnd:            &evalEnd,
					ViolationsObserved: models.BoolPointer(false),
					Remarks:            models.StringPointer("This is a remark."),
				},
			},
		}, nil)
		// Generate an etag
		eTag := etag.GenerateEtag(evaluationReport.UpdatedAt)
		// Submit the report
		err := updater.SubmitEvaluationReport(suite.AppContextForTest(), evaluationReport.ID, officeUser.ID, eTag)
		suite.Equal(models.ErrInvalidTransition, errors.Cause(err))
	})
}

func (suite EvaluationReportSuite) TestUpdateEvaluationReport() {
	updater := NewEvaluationReportUpdater()

	suite.Run("successful save", func() {
		// Create a report
		originalReport := factory.BuildEvaluationReport(suite.DB(), nil, nil)

		// Copy report to new object
		report := originalReport

		report.Remarks = models.StringPointer("spectacular packing job!!")

		// Attempt to update the report
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, etag.GenerateEtag(report.UpdatedAt))
		suite.NoError(err)

		// Fetch the report from the database and make sure that it got updated
		var updatedReport models.EvaluationReport
		err = suite.DB().Find(&updatedReport, originalReport.ID)
		suite.NoError(err)

		suite.Equal(report.Remarks, updatedReport.Remarks)
		suite.Nil(updatedReport.SubmittedAt)
	})

	suite.Run("saving report with pre-existing violations should be removed or preserved based on observedViolations bool", func() {
		report := factory.BuildEvaluationReport(suite.DB(), nil, nil)
		testdatagen.MakeReportViolation(suite.DB(), testdatagen.Assertions{ReportViolation: models.ReportViolation{
			ReportID:    report.ID,
			Violation:   models.PWSViolation{},
			ViolationID: uuid.UUID{},
		}})
		report.ViolationsObserved = models.BoolPointer(true)

		// do the update
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, etag.GenerateEtag(report.UpdatedAt))
		suite.NoError(err)

		var reportViolations models.ReportViolations
		err = suite.DB().Where("report_id = ?", report.ID).All(&reportViolations)
		suite.NoError(err)
		// we should find any report violations, which means this object should have a 1 length
		suite.Equal(1, len(reportViolations))

		report.ViolationsObserved = models.BoolPointer(false)

		// do the update
		err = updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, etag.GenerateEtag(report.UpdatedAt))
		suite.NoError(err)

		err = suite.DB().Where("report_id = ?", report.ID).All(&reportViolations)
		suite.NoError(err)
		// we shouldn't find any report violations, which means this object should have a 0 length
		suite.Equal(0, len(reportViolations))
	})

	suite.Run("saving report does not overwrite readonly fields", func() {
		// Create a report and save it to the database
		move := factory.BuildMove(suite.DB(), nil, nil)
		originalReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model:    move,
				LinkOnly: true,
			},
			{
				Model: models.EvaluationReport{
					Type: models.EvaluationReportTypeShipment,
				},
			},
		}, nil)
		// Copy report to new object
		reportPayload := originalReport

		wrongUUID := uuid.Must(uuid.NewV4())
		reportPayload.Remarks = models.StringPointer("spectacular packing job!!")
		reportPayload.MoveID = wrongUUID
		reportPayload.ShipmentID = &wrongUUID

		// Attempt to update the reportPayload
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &reportPayload, reportPayload.OfficeUserID, etag.GenerateEtag(reportPayload.UpdatedAt))
		suite.NoError(err)

		// Fetch the reportPayload from the database and make sure that it got updated
		var updatedReport models.EvaluationReport
		err = suite.DB().Find(&updatedReport, originalReport.ID)
		suite.NoError(err)

		suite.Equal(reportPayload.Remarks, updatedReport.Remarks)
		suite.Equal(originalReport.MoveID, updatedReport.MoveID)
		suite.Equal(*originalReport.ShipmentID, *updatedReport.ShipmentID)

		swaggerTimeFormat := "2006-01-02T15:04:05.99Z07:00"
		suite.Equal(originalReport.CreatedAt.Format(swaggerTimeFormat), time.Time(updatedReport.CreatedAt).Format(swaggerTimeFormat))
	})

	suite.Run("saving evaluation report with bad report id should fail", func() {
		// Create a report
		report := factory.BuildEvaluationReport(suite.DB(), nil, nil)

		// Overwrite the report's ID with some nonsense
		report.ID = uuid.Must(uuid.NewV4())

		// Attempt to update the report
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, etag.GenerateEtag(report.UpdatedAt))

		// Our bogus report ID should cause an error
		suite.Error(err)
		suite.IsType(apperror.NotFoundError{}, err)
	})
	suite.Run("saving evaluation report created by another office user should fail", func() {
		// Create a report
		report := factory.BuildEvaluationReport(suite.DB(), nil, nil)

		otherOfficeUserID := uuid.Must(uuid.NewV4())

		// Attempt to update the report
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, otherOfficeUserID, etag.GenerateEtag(report.UpdatedAt))

		// Our bogus office user ID should cause an error
		suite.Error(err)
		suite.IsType(apperror.ForbiddenError{}, err)
	})

	suite.Run("updating a non-draft report should fail", func() {
		// Create a report
		originalReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model: models.EvaluationReport{
					SubmittedAt: models.TimePointer(time.Now()),
				},
			},
		}, nil)
		report := originalReport
		report.Remarks = models.StringPointer("spectacular packing job!!")

		// Attempt to update the report
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, etag.GenerateEtag(report.UpdatedAt))
		suite.Error(err)
	})
	suite.Run("updating a report with a bad ETag should fail", func() {
		// Create a report
		originalReport := factory.BuildEvaluationReport(suite.DB(), []factory.Customization{
			{
				Model: models.EvaluationReport{
					SubmittedAt: models.TimePointer(time.Now()),
				},
			},
		}, nil)

		report := originalReport
		report.Remarks = models.StringPointer("spectacular packing job!!")

		// Attempt to update the report
		err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, "not a real etag")
		suite.Error(err)
		suite.IsType(apperror.PreconditionFailedError{}, err)
	})

	physical := models.EvaluationReportInspectionTypePhysical
	virtual := models.EvaluationReportInspectionTypeVirtual
	dataReview := models.EvaluationReportInspectionTypeDataReview
	currentTime := time.Now()

	testCases := map[string]struct {
		inspectionType                     *models.EvaluationReportInspectionType
		observedShipmentDeliveryDate       *time.Time
		observedShipmentPhysicalPickupDate *time.Time
		expectedError                      bool
		location                           *models.EvaluationReportLocationType
	}{
		"observed shipment delivery date set for physical report type should succeed": {
			inspectionType:               &physical,
			observedShipmentDeliveryDate: &currentTime,
			expectedError:                false,
		},
		"observed shipment delivery date set for virtual report type should fail": {
			inspectionType:               &virtual,
			observedShipmentDeliveryDate: &currentTime,
			expectedError:                true,
		},
		"observed shipment phsyical pickup set for data review report type should fail": {
			inspectionType:                     &dataReview,
			observedShipmentPhysicalPickupDate: &currentTime,
			expectedError:                      true,
		},
	}

	for name, tc := range testCases {
		name := name
		tc := tc

		suite.Run(name, func() {
			report := factory.BuildEvaluationReport(suite.DB(), nil, nil)
			report.InspectionType = tc.inspectionType
			report.ObservedShipmentDeliveryDate = tc.observedShipmentDeliveryDate
			report.ObservedShipmentPhysicalPickupDate = tc.observedShipmentPhysicalPickupDate
			report.Location = tc.location
			err := updater.UpdateEvaluationReport(suite.AppContextForTest(), &report, report.OfficeUserID, etag.GenerateEtag(report.UpdatedAt))
			if tc.expectedError {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
