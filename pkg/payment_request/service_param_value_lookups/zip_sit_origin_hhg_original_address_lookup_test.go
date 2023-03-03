package serviceparamvaluelookups

import (
	"github.com/transcom/mymove/pkg/factory"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ServiceParamValueLookupsSuite) TestZipSITOriginHHGOriginalAddressLookup() {
	key := models.ServiceItemParamNameZipSITOriginHHGOriginalAddress

	originZip := "30901"
	actualOriginZipSameZip3 := "30907"

	var mtoServiceItemWithSITOriginZips models.MTOServiceItem
	var paymentRequest models.PaymentRequest
	var mtoServiceItemNoSITOriginZips models.MTOServiceItem

	setupTestData := func() {
		reService := factory.BuildReServiceByCode(suite.DB(), models.ReServiceCodeDOFSIT)

		originAddress := factory.BuildAddress(suite.DB(),
			[]factory.Customization{
				{
					Model: models.Address{
						PostalCode: originZip,
					},
				},
			}, nil)

		actualOriginSameZip3Address := factory.BuildAddress(suite.DB(),
			[]factory.Customization{
				{
					Model: models.Address{
						PostalCode: actualOriginZipSameZip3,
					},
				},
			}, nil)

		move := testdatagen.MakeDefaultMove(suite.DB())

		paymentRequest = testdatagen.MakePaymentRequest(suite.DB(),
			testdatagen.Assertions{
				Move: move,
			})

		mtoServiceItemWithSITOriginZips = testdatagen.MakeMTOServiceItem(suite.DB(),
			testdatagen.Assertions{
				ReService: reService,
				Move:      move,
				MTOServiceItem: models.MTOServiceItem{
					SITOriginHHGOriginalAddressID: &originAddress.ID,
					SITOriginHHGOriginalAddress:   &originAddress,
					SITOriginHHGActualAddressID:   &actualOriginSameZip3Address.ID,
					SITOriginHHGActualAddress:     &actualOriginSameZip3Address,
				},
			},
		)

		mtoServiceItemNoSITOriginZips = testdatagen.MakeMTOServiceItem(suite.DB(),
			testdatagen.Assertions{
				ReService:      reService,
				Move:           move,
				MTOServiceItem: models.MTOServiceItem{},
			},
		)
	}

	suite.Run("success SIT origin original zip lookup", func() {
		setupTestData()

		paramLookup, err := ServiceParamLookupInitialize(suite.AppContextForTest(), suite.planner, mtoServiceItemWithSITOriginZips, paymentRequest.ID, paymentRequest.MoveTaskOrderID, nil)
		suite.FatalNoError(err)

		sitOriginZipOriginal, err := paramLookup.ServiceParamValue(suite.AppContextForTest(), key)
		suite.FatalNoError(err)
		expected := mtoServiceItemWithSITOriginZips.SITOriginHHGOriginalAddress.PostalCode
		suite.Equal(expected, sitOriginZipOriginal)
	})

	suite.Run("fail to find SIT origin original zip lookup", func() {
		setupTestData()

		paramLookup, err := ServiceParamLookupInitialize(suite.AppContextForTest(), suite.planner, mtoServiceItemNoSITOriginZips, paymentRequest.ID, paymentRequest.MoveTaskOrderID, nil)
		suite.FatalNoError(err)

		sitOriginZipOriginal, err := paramLookup.ServiceParamValue(suite.AppContextForTest(), key)
		suite.Error(err)
		suite.Equal("", sitOriginZipOriginal)
		suite.Contains(err.Error(), "nil SITOriginHHGOriginalAddressID")
	})

}
