import Swagger from 'swagger-client';

import { makeSwaggerRequest, requestInterceptor, responseInterceptor } from './swaggerRequest';

let internalClient = null;

// setting up the same config from Swagger/api.js
export async function getInternalClient() {
  if (!internalClient) {
    internalClient = await Swagger({
      url: '/internal/swagger.yaml',
      requestInterceptor,
      responseInterceptor,
    });
  }

  return internalClient;
}

// Attempt at catch-all error handling
// TODO improve this function when we have better standardized errors
export function getResponseError(response, defaultErrorMessage) {
  return response?.body?.detail || response?.statusText || defaultErrorMessage;
}

export async function makeInternalRequest(operationPath, params = {}, options = {}) {
  const client = await getInternalClient();
  return makeSwaggerRequest(client, operationPath, params, options);
}

export async function getLoggedInUser(normalize = true) {
  return makeInternalRequest('users.showLoggedInUser', {}, { normalize });
}

export async function getLoggedInUserQueries(key, normalize = false) {
  return makeInternalRequest('users.showLoggedInUser', {}, { normalize });
}

export async function getFeatureFlagForUser(key, flagContext) {
  const normalize = false;
  return makeInternalRequest('featureFlags.featureFlagForUser', { key, flagContext }, { normalize });
}

export async function getMTOShipmentsForMove(moveTaskOrderID, normalize = true) {
  return makeInternalRequest(
    'mtoShipment.listMTOShipments',
    { moveTaskOrderID },
    { normalize, label: 'mtoShipment.listMTOShipments', schemaKey: 'mtoShipments' },
  );
}

/** BELOW API CALLS ARE NOT NORMALIZED BY DEFAULT */

/** SERVICE MEMBERS */
export async function createServiceMember(serviceMember = {}) {
  return makeInternalRequest(
    'service_members.createServiceMember',
    { createServiceMemberPayload: serviceMember },
    { normalize: false },
  );
}

export async function getServiceMember(serviceMemberId) {
  return makeInternalRequest(
    'service_members.showServiceMember',
    {
      serviceMemberId,
    },
    {
      normalize: false,
    },
  );
}

export async function patchServiceMember(serviceMember) {
  return makeInternalRequest(
    'service_members.patchServiceMember',
    {
      serviceMemberId: serviceMember.id,
      patchServiceMemberPayload: serviceMember,
    },
    {
      normalize: false,
    },
  );
}

/** BACKUP CONTACTS */
export async function createBackupContactForServiceMember(serviceMemberId, backupContact) {
  return makeInternalRequest(
    'backup_contacts.createServiceMemberBackupContact',
    {
      serviceMemberId,
      createBackupContactPayload: backupContact,
    },
    {
      normalize: false,
    },
  );
}

export async function patchBackupContact(backupContact) {
  return makeInternalRequest(
    'backup_contacts.updateServiceMemberBackupContact',
    {
      backupContactId: backupContact.id,
      updateServiceMemberBackupContactPayload: backupContact,
    },
    {
      normalize: false,
    },
  );
}

/** ORDERS */
export async function getOrdersForServiceMember(serviceMemberId) {
  return makeInternalRequest(
    'service_members.showServiceMemberOrders',
    {
      serviceMemberId,
    },
    {
      normalize: false,
    },
  );
}

export async function createOrders(orders) {
  return makeInternalRequest(
    'orders.createOrders',
    {
      createOrders: orders,
    },
    {
      normalize: false,
    },
  );
}

export async function patchOrders(orders) {
  return makeInternalRequest(
    'orders.updateOrders',
    {
      ordersId: orders.id,
      updateOrders: orders,
    },
    {
      normalize: false,
    },
  );
}

/** UPLOADS */
export async function createUpload(file) {
  return makeInternalRequest(
    'uploads.createUpload',
    {
      file,
    },
    {
      normalize: false,
    },
  );
}

export async function createUploadForAmendedOrdersDocument(file, ordersId) {
  return makeInternalRequest(
    'orders.uploadAmendedOrders',
    {
      ordersId,
      file,
    },
    {
      normalize: false,
    },
  );
}

export async function createUploadForDocument(file, documentId) {
  return makeInternalRequest(
    'uploads.createUpload',
    {
      documentId,
      file,
    },
    {
      normalize: false,
    },
  );
}

export async function createUploadForPPMDocument(ppmShipmentId, documentId, file) {
  return makeInternalRequest(
    'ppm.createPPMUpload',
    {
      ppmShipmentId,
      documentId,
      file,
    },
    {
      normalize: false,
    },
  );
}

export async function deleteUpload(uploadId) {
  return makeInternalRequest(
    'uploads.deleteUpload',
    {
      uploadId,
    },
    {
      normalize: false,
    },
  );
}

/** MOVES */
export async function getMove(moveId) {
  return makeInternalRequest(
    'moves.showMove',
    {
      moveId,
    },
    {
      normalize: false,
    },
  );
}

export async function patchMove(moveId, move, ifMatchETag) {
  return makeInternalRequest(
    'moves.patchMove',
    {
      moveId,
      'If-Match': ifMatchETag,
      patchMovePayload: move,
    },
    {
      normalize: false,
    },
  );
}

export async function submitMoveForApproval(moveId, certificate) {
  return makeInternalRequest(
    'moves.submitMoveForApproval',
    {
      moveId,
      submitMoveForApprovalPayload: {
        certificate,
      },
    },
    {
      normalize: false,
    },
  );
}

export async function submitAmendedOrders(moveId) {
  return makeInternalRequest(
    'moves.submitAmendedOrders',
    {
      moveId,
    },
    {
      normalize: false,
    },
  );
}

/** MTO SHIPMENTS */
export async function createMTOShipment(mtoShipment) {
  return makeInternalRequest(
    'mtoShipment.createMTOShipment',
    {
      body: mtoShipment,
    },
    {
      normalize: false,
    },
  );
}

export async function patchMTOShipment(mtoShipmentId, mtoShipment, ifMatchETag) {
  return makeInternalRequest(
    'mtoShipment.updateMTOShipment',
    {
      mtoShipmentId,
      'If-Match': ifMatchETag,
      body: mtoShipment,
    },
    {
      normalize: false,
    },
  );
}

export async function deleteMTOShipment(mtoShipmentId) {
  return makeInternalRequest(
    'mtoShipment.deleteShipment',
    {
      mtoShipmentId,
    },
    {
      normalize: false,
    },
  );
}

export async function createWeightTicket(ppmShipmentId) {
  return makeInternalRequest(
    'ppm.createWeightTicket',
    {
      ppmShipmentId,
    },
    {
      normalize: false,
    },
  );
}

export async function patchWeightTicket(ppmShipmentId, weightTicketId, payload, eTag) {
  return makeInternalRequest(
    'ppm.updateWeightTicket',
    {
      ppmShipmentId,
      weightTicketId,
      'If-Match': eTag,
      updateWeightTicketPayload: payload,
    },
    {
      normalize: false,
    },
  );
}

export async function deleteWeightTicket(ppmShipmentId, weightTicketId) {
  return makeInternalRequest(
    'ppm.deleteWeightTicket',
    {
      ppmShipmentId,
      weightTicketId,
    },
    {
      normalize: false,
    },
  );
}

export async function createProGearWeightTicket(ppmShipmentId) {
  return makeInternalRequest(
    'ppm.createProGearWeightTicket',
    {
      ppmShipmentId,
    },
    {
      normalize: false,
    },
  );
}

export async function patchProGearWeightTicket(ppmShipmentId, proGearWeightTicketId, payload, eTag) {
  return makeInternalRequest(
    'ppm.updateProGearWeightTicket',
    {
      ppmShipmentId,
      proGearWeightTicketId,
      'If-Match': eTag,
      updateProGearWeightTicket: payload,
    },
    {
      normalize: false,
    },
  );
}

export async function deleteProGearWeightTicket(ppmShipmentId, proGearWeightTicketId) {
  return makeInternalRequest(
    'ppm.deleteProGearWeightTicket',
    {
      ppmShipmentId,
      proGearWeightTicketId,
    },
    {
      normalize: false,
    },
  );
}

/** PPMS */
export async function getPPMsForMove(moveId) {
  return makeInternalRequest(
    'ppm.indexPersonallyProcuredMoves',
    {
      moveId,
    },
    {
      normalize: false,
    },
  );
}

export async function createPPMForMove(moveId, ppm) {
  return makeInternalRequest(
    'ppm.createPersonallyProcuredMove',
    {
      moveId,
      createPersonallyProcuredMovePayload: ppm,
    },
    {
      normalize: false,
    },
  );
}

export async function patchPPM(moveId, ppm) {
  return makeInternalRequest(
    'ppm.patchPersonallyProcuredMove',
    {
      moveId,
      personallyProcuredMoveId: ppm.id,
      patchPersonallyProcuredMovePayload: ppm,
    },
    {
      normalize: false,
    },
  );
}

export async function calculatePPMEstimate(moveDate, originZip, originDutyLocationZip, ordersId, weightEstimate) {
  return makeInternalRequest(
    'ppm.showPPMEstimate',
    {
      original_move_date: moveDate,
      origin_zip: originZip,
      origin_duty_location_zip: originDutyLocationZip,
      orders_id: ordersId,
      weight_estimate: weightEstimate,
    },
    {
      normalize: false,
    },
  );
}

export async function persistPPMEstimate(moveId, ppmId) {
  return makeInternalRequest(
    'ppm.updatePersonallyProcuredMoveEstimate',
    {
      moveId,
      personallyProcuredMoveId: ppmId,
    },
    {
      normalize: false,
    },
  );
}

export async function calculatePPMSITEstimate(ppmId, moveDate, sitDays, originZip, ordersId, weightEstimate) {
  return makeInternalRequest(
    'ppm.showPPMSitEstimate',
    {
      personally_procured_move_id: ppmId,
      original_move_date: moveDate,
      days_in_storage: sitDays,
      origin_zip: originZip,
      orders_id: ordersId,
      weight_estimate: weightEstimate,
    },
    {
      normalize: false,
    },
  );
}

export async function requestPayment(ppmId) {
  return makeInternalRequest(
    'ppm.requestPPMPayment',
    {
      personallyProcuredMoveId: ppmId,
    },
    {
      normalize: false,
    },
  );
}

export async function createMovingExpense(ppmShipmentId) {
  return makeInternalRequest(
    'ppm.createMovingExpense',
    {
      ppmShipmentId,
    },
    {
      normalize: false,
    },
  );
}

export async function patchMovingExpense(ppmShipmentId, movingExpenseId, payload, eTag) {
  return makeInternalRequest(
    'ppm.updateMovingExpense',
    {
      ppmShipmentId,
      movingExpenseId,
      'If-Match': eTag,
      updateMovingExpense: payload,
    },
    {
      normalize: false,
    },
  );
}

export async function deleteMovingExpense(ppmShipmentId, movingExpenseId) {
  return makeInternalRequest(
    'ppm.deleteMovingExpense',
    {
      ppmShipmentId,
      movingExpenseId,
    },
    {
      normalize: false,
    },
  );
}

export async function submitPPMShipmentSignedCertification(ppmShipmentId, payload) {
  return makeInternalRequest(
    'ppm.submitPPMShipmentDocumentation',
    {
      ppmShipmentId,
      savePPMShipmentSignedCertificationPayload: payload,
    },
    {
      normalize: false,
    },
  );
}

export async function searchTransportationOffices(search) {
  return makeInternalRequest('transportation_offices.getTransportationOffices', { search }, { normalize: false });
}
