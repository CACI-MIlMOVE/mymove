package handlers

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/auth"
	uploadop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/uploads"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/models"
	uploaderpkg "github.com/transcom/mymove/pkg/uploader"
)

func payloadForUploadModel(upload models.Upload, url string) *internalmessages.UploadPayload {
	return &internalmessages.UploadPayload{
		ID:          fmtUUID(upload.ID),
		Filename:    swag.String(upload.Filename),
		ContentType: swag.String(upload.ContentType),
		URL:         fmtURI(url),
		Bytes:       &upload.Bytes,
		CreatedAt:   fmtDateTime(upload.CreatedAt),
		UpdatedAt:   fmtDateTime(upload.UpdatedAt),
	}
}

// CreateUploadHandler creates a new upload via POST /documents/{documentID}/uploads
type CreateUploadHandler HandlerContext

// Handle creates a new Upload from a request payload
func (h CreateUploadHandler) Handle(params uploadop.CreateUploadParams) middleware.Responder {

	file, ok := params.File.(*runtime.File)
	if !ok {
		h.logger.Error("This should always be a runtime.File, something has changed in go-swagger.")
		return uploadop.NewCreateUploadInternalServerError()
	}
	h.logger.Info("File name and size: ", zap.String("name", file.Header.Filename), zap.Int64("size", file.Header.Size))

	// User should always be populated by middleware
	session := auth.SessionFromRequestContext(params.HTTPRequest)
	documentID, err := uuid.FromString(params.DocumentID.String())
	if err != nil {
		h.logger.Info("Badly formed UUID for document", zap.String("document_id", params.DocumentID.String()), zap.Error(err))
		return uploadop.NewCreateUploadBadRequest()
	}

	// Fetch document to ensure user has access to it
	document, docErr := models.FetchDocument(h.db, session, documentID)
	if docErr != nil {
		return responseForError(h.logger, docErr)
	}

	uploader := uploaderpkg.NewUploader(h.db, h.logger, h.storage)
	newUpload, verrs, err := uploader.CreateUpload(document.ID, session.UserID, file)
	if err != nil {
		switch err := errors.Cause(err).(type) {
		case *uploaderpkg.ZeroLengthFile:
			return uploadop.NewCreateUploadBadRequest()
		default:
			h.logger.Error("Failed to create upload", zap.Error(err))
			return uploadop.NewCreateUploadInternalServerError()
		}
	} else if verrs != nil {
		payload := createFailedValidationPayload(verrs)
		return uploadop.NewCreateUploadBadRequest().WithPayload(payload)
	}

	url, err := uploader.PresignedURL(newUpload)
	if err != nil {
		h.logger.Error("failed to get presigned url", zap.Error(err))
		return uploadop.NewCreateUploadInternalServerError()
	}
	uploadPayload := payloadForUploadModel(*newUpload, url)
	return uploadop.NewCreateUploadCreated().WithPayload(uploadPayload)
}

// DeleteUploadHandler deletes an upload
type DeleteUploadHandler HandlerContext

// Handle deletes an upload
func (h DeleteUploadHandler) Handle(params uploadop.DeleteUploadParams) middleware.Responder {
	session := auth.SessionFromRequestContext(params.HTTPRequest)

	uploadID, _ := uuid.FromString(params.UploadID.String())
	upload, err := models.FetchUpload(h.db, session, uploadID)
	if err != nil {
		return responseForError(h.logger, err)
	}

	key := h.storage.Key("documents", upload.DocumentID.String(), "uploads", upload.ID.String())
	err = h.storage.Delete(key)
	if err != nil {
		return responseForError(h.logger, err)
	}

	err = models.DeleteUpload(h.db, &upload)
	if err != nil {
		return responseForError(h.logger, err)
	}

	return uploadop.NewDeleteUploadCreated()
}

// DeleteUploadsHandler deletes a collection of uploads
type DeleteUploadsHandler HandlerContext

// Handle deletes uploads
func (h DeleteUploadsHandler) Handle(params uploadop.DeleteUploadsParams) middleware.Responder {
	// User should always be populated by middleware
	session := auth.SessionFromRequestContext(params.HTTPRequest)

	for _, uploadID := range params.UploadIds {
		uuid, _ := uuid.FromString(uploadID.String())
		upload, err := models.FetchUpload(h.db, session, uuid)
		if err != nil {
			return responseForError(h.logger, err)
		}

		key := h.storage.Key("documents", upload.DocumentID.String(), "uploads", upload.ID.String())
		err = h.storage.Delete(key)
		if err != nil {
			return responseForError(h.logger, err)
		}

		err = models.DeleteUpload(h.db, &upload)
		if err != nil {
			return responseForError(h.logger, err)
		}
	}

	return uploadop.NewDeleteUploadsCreated()
}
