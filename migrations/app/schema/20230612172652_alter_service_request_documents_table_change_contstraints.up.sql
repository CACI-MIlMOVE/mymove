ALTER TABLE service_request_documents
	DROP CONSTRAINT service_request_documents_unique_key;

ALTER TABLE service_request_document_uploads
	ADD CONSTRAINT service_request_documents_uploads_unique_key UNIQUE (upload_id);

