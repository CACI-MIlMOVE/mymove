import { getClient as getInternalClient, checkResponse, getPublicClient } from 'shared/Swagger/api';

export async function CreateUpload(fileUpload, documentId, isPublic = false) {
  const clientToUse = isPublic ? getPublicClient : getInternalClient;
  const client = await clientToUse();
  const response = await client.apis.uploads.createUpload({
    file: fileUpload,
    documentId,
  });
  checkResponse(response, 'failed to upload file due to server error');
  return response.body;
}

export async function DeleteUpload(uploadId, isPublic = false) {
  const clientToUse = isPublic ? getPublicClient : getInternalClient;
  const client = await clientToUse();
  const response = await client.apis.uploads.deleteUpload({
    uploadId,
  });
  checkResponse(response, 'failed to delete file due to server error');
  return response.body;
}

export async function CreateDocument(name, serviceMemberId) {
  const client = await getInternalClient();
  const response = await client.apis.documents.createDocument({
    documentPayload: {
      name,
      service_member_id: serviceMemberId,
    },
  });
  checkResponse(response, 'failed to create document due to server error');
  return response.body;
}

export async function ValidateZipRateData(zipCode, zipType) {
  const client = await getInternalClient();
  const response = await client.apis.postal_codes.validatePostalCodeWithRateData({
    postal_code: zipCode,
    postal_code_type: zipType,
  });
  checkResponse(response, 'failed to validate ppm data due to server error');
  return response.body;
}
