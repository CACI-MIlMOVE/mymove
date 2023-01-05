package movingexpense

import (
	"github.com/gobuffalo/validate/v3"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/models"
)

func checkID() movingExpenseValidator {
	return movingExpenseValidatorFunc(func(_ appcontext.AppContext, newMovingExpense *models.MovingExpense, originalMovingExpense *models.MovingExpense) error {
		verrs := validate.NewErrors()

		if newMovingExpense == nil || originalMovingExpense == nil {
			return verrs
		}

		if newMovingExpense.ID != originalMovingExpense.ID {
			verrs.Add("ID", "new MovingExpense ID must match original MovingExpense ID")
		}

		return verrs
	})
}

func checkBaseRequiredFields() movingExpenseValidator {
	return movingExpenseValidatorFunc(func(_ appcontext.AppContext, newMovingExpense *models.MovingExpense, originalMovingExpense *models.MovingExpense) error {
		verrs := validate.NewErrors()

		if newMovingExpense.PPMShipmentID.IsNil() {
			verrs.Add("PPMShipmentID", "PPMShipmentID must exist")
		}

		if newMovingExpense.Document.ServiceMemberID.IsNil() {
			verrs.Add("ServiceMemberID", "Document ServiceMemberID must exist")
		}

		return verrs
	})
}

func checkAdditionalRequiredFields() movingExpenseValidator {
	return movingExpenseValidatorFunc(func(_ appcontext.AppContext, newMovingExpense *models.MovingExpense, originalMovingExpense *models.MovingExpense) error {
		verrs := validate.NewErrors()

		// the model Validate should test for allowed values
		if newMovingExpense.MovingExpenseType == nil || *newMovingExpense.MovingExpenseType == "" {
			verrs.Add("MovingExpenseType", "MovingExpenseType must exist")
		} else if *newMovingExpense.MovingExpenseType == models.MovingExpenseReceiptTypeStorage {
			if newMovingExpense.SITStartDate == nil || newMovingExpense.SITStartDate.IsZero() {
				verrs.Add("SITStartDate", "SITStartDate is required for storage expenses")
			}

			if newMovingExpense.SITEndDate == nil || newMovingExpense.SITEndDate.IsZero() {
				verrs.Add("SITEndDate", "SITEndDate is required for storage expenses")
			}

			if newMovingExpense.SITStartDate != nil && newMovingExpense.SITEndDate != nil {
				if newMovingExpense.SITEndDate.Before(*newMovingExpense.SITStartDate) {
					verrs.Add("SITStartDate", "SITStartDate must be before SITEndDate")
				}
			}
		}

		if newMovingExpense.Description == nil || *newMovingExpense.Description == "" {
			verrs.Add("Description", "Description must have a value of at least 0")
		}

		if newMovingExpense.PaidWithGTCC == nil {
			verrs.Add("PaidWithGTCC", "PaidWithGTCC is required")
		}

		if newMovingExpense.Amount == nil || *newMovingExpense.Amount < 1 {
			verrs.Add("Amount", "Amount must have a value of at least 1")
		}

		if newMovingExpense.MissingReceipt == nil {
			verrs.Add("MissingReceipt", "MissingReceipt is required")
		}

		if len(originalMovingExpense.Document.UserUploads) < 1 {
			verrs.Add("Document", "At least 1 receipt file is required")
		}

		if newMovingExpense.Status != nil {
			if (*newMovingExpense.Status == models.PPMDocumentStatusExcluded || *newMovingExpense.Status == models.PPMDocumentStatusRejected) && (newMovingExpense.Reason == nil || *newMovingExpense.Reason == "") {
				verrs.Add("Reason", "A reason must be provided when the status is EXCLUDED or REJECTED")
			}
		}

		return verrs
	})
}

func createChecks() []movingExpenseValidator {
	return []movingExpenseValidator{
		checkID(),
		checkBaseRequiredFields(),
	}
}

func updateChecks() []movingExpenseValidator {
	return []movingExpenseValidator{
		checkID(),
		checkBaseRequiredFields(),
		checkAdditionalRequiredFields(),
	}
}
