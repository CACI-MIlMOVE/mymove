package supportapi

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/appcontext"
	webhookops "github.com/transcom/mymove/pkg/gen/supportapi/supportoperations/webhook"
	"github.com/transcom/mymove/pkg/gen/supportmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/handlers/supportapi/internal/payloads"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/event"
)

// ReceiveWebhookNotificationHandler passes through a message
type ReceiveWebhookNotificationHandler struct {
	handlers.HandlerConfig
}

// Handle receipt of message
func (h ReceiveWebhookNotificationHandler) Handle(params webhookops.ReceiveWebhookNotificationParams) middleware.Responder {
	return h.AuditableAppContextFromRequestWithErrors(params.HTTPRequest,
		func(appCtx appcontext.AppContext) (middleware.Responder, error) {
			notif := params.Body

			objectID := "<empty>"
			if notif.ObjectID != nil {
				objectID = notif.ObjectID.String()
			}
			mtoID := "<empty>"
			if notif.MoveTaskOrderID != nil {
				mtoID = notif.MoveTaskOrderID.String()
			}

			appCtx.Logger().Info("Received Webhook Notification: ",
				zap.String("id", notif.ID.String()),
				zap.String("eventKey", notif.EventKey),
				zap.String("createdAt", notif.CreatedAt.String()),
				zap.String("traceID", notif.TraceID.String()),
				zap.String("moveID", mtoID),
				zap.String("objectID", objectID))
			return webhookops.NewReceiveWebhookNotificationOK().WithPayload(notif), nil
		})
}

// CreateWebhookNotificationHandler is the interface to handle the createWebhookNotification
type CreateWebhookNotificationHandler struct {
	handlers.HandlerConfig
}

// Handle handles the endpoint request to the createWebhookNotification handler
func (h CreateWebhookNotificationHandler) Handle(params webhookops.CreateWebhookNotificationParams) middleware.Responder {
	return h.AuditableAppContextFromRequestWithErrors(params.HTTPRequest,
		func(appCtx appcontext.AppContext) (middleware.Responder, error) {
			payload := params.Body

			var err error
			if payload == nil {
				// create a default notification payload
				message := "{ \"message\": \"This is a test notification\" }"
				payload = &supportmessages.WebhookNotification{
					EventKey: string(event.TestCreateEventKey),
					TraceID:  *handlers.FmtUUID(h.GetTraceIDFromRequest(params.HTTPRequest)),
					Object:   models.StringPointer(message),
					Status:   supportmessages.WebhookNotificationStatusPENDING,
				}
			}
			// Convert to model and create in DB
			notification, verrs := payloads.WebhookNotificationModel(payload, h.GetTraceIDFromRequest(params.HTTPRequest))
			if verrs == nil {
				verrs, err = appCtx.DB().ValidateAndCreate(notification)
			}
			if verrs != nil && verrs.HasAny() {
				appCtx.Logger().Error("Error validating WebhookNotification: ", zap.Error(verrs))

				return webhookops.NewCreateWebhookNotificationUnprocessableEntity().WithPayload(payloads.ValidationError(
					"The notification definition is invalid.", h.GetTraceIDFromRequest(params.HTTPRequest), verrs)), verrs
			}
			if err != nil {
				appCtx.Logger().Error("Error creating WebhookNotification: ", zap.Error(err))
				return webhookops.NewCreateWebhookNotificationInternalServerError().WithPayload(
					payloads.InternalServerError(models.StringPointer(err.Error()), h.GetTraceIDFromRequest(params.HTTPRequest))), err
			}

			payload = payloads.WebhookNotification(notification)
			return webhookops.NewCreateWebhookNotificationCreated().WithPayload(payload), nil
		})
}
