package ghcapi

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/factory"
	documentop "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/ghc_documents"
	"github.com/transcom/mymove/pkg/models"
	storageTest "github.com/transcom/mymove/pkg/storage/test"
	"github.com/transcom/mymove/pkg/uploader"
)

func (suite *HandlerSuite) TestGetDocumentHandler() {
	t := suite.T()

	userUpload := factory.BuildUserUpload(suite.DB(), nil, nil)

	documentID := userUpload.DocumentID
	var document models.Document

	err := suite.DB().Eager("ServiceMember.User").Find(&document, documentID)
	if err != nil {
		suite.Fail("could not load document: %s", err)
	}

	params := documentop.NewGetDocumentParams()
	params.DocumentID = strfmt.UUID(documentID.String())

	req := &http.Request{}
	req = suite.AuthenticateRequest(req, document.ServiceMember)
	params.HTTPRequest = req

	handlerConfig := suite.HandlerConfig()
	fakeS3 := storageTest.NewFakeS3Storage(true)
	handlerConfig.SetFileStorer(fakeS3)
	handler := GetDocumentHandler{handlerConfig}

	// Validate incoming payload: no body to validate

	response := handler.Handle(params)

	showResponse, ok := response.(*documentop.GetDocumentOK)
	if !ok {
		suite.Fail("Request failed: %#v", response)
	}
	documentPayload := showResponse.Payload

	// Validate outgoing payload
	suite.NoError(documentPayload.Validate(strfmt.Default))

	responseDocumentUUID := documentPayload.ID.String()
	if responseDocumentUUID != documentID.String() {
		t.Errorf("wrong document uuid, expected %v, got %v", documentID, responseDocumentUUID)
	}

	if len(documentPayload.Uploads) != 1 {
		t.Errorf("wrong number of uploads, expected 1, got %d", len(documentPayload.Uploads))
	}

	uploadPayload := documentPayload.Uploads[0]
	expectedURL := fmt.Sprintf("https://example.com/dir/%s?contentType=%s&signed=test", userUpload.Upload.StorageKey, uploader.FileTypePDF)
	if (uploadPayload.URL).String() != expectedURL {
		t.Errorf("wrong URL for upload, expected %s, got %s", expectedURL, uploadPayload.URL)
	}
}
