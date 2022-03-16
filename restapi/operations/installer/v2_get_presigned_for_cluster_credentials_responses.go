// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2GetPresignedForClusterCredentialsOKCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsOK
const V2GetPresignedForClusterCredentialsOKCode int = 200

/*V2GetPresignedForClusterCredentialsOK Success.

swagger:response v2GetPresignedForClusterCredentialsOK
*/
type V2GetPresignedForClusterCredentialsOK struct {

	/*
	  In: Body
	*/
	Payload *models.PresignedURL `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsOK creates V2GetPresignedForClusterCredentialsOK with default headers values
func NewV2GetPresignedForClusterCredentialsOK() *V2GetPresignedForClusterCredentialsOK {

	return &V2GetPresignedForClusterCredentialsOK{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials o k response
func (o *V2GetPresignedForClusterCredentialsOK) WithPayload(payload *models.PresignedURL) *V2GetPresignedForClusterCredentialsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials o k response
func (o *V2GetPresignedForClusterCredentialsOK) SetPayload(payload *models.PresignedURL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsBadRequestCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsBadRequest
const V2GetPresignedForClusterCredentialsBadRequestCode int = 400

/*V2GetPresignedForClusterCredentialsBadRequest Error.

swagger:response v2GetPresignedForClusterCredentialsBadRequest
*/
type V2GetPresignedForClusterCredentialsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsBadRequest creates V2GetPresignedForClusterCredentialsBadRequest with default headers values
func NewV2GetPresignedForClusterCredentialsBadRequest() *V2GetPresignedForClusterCredentialsBadRequest {

	return &V2GetPresignedForClusterCredentialsBadRequest{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials bad request response
func (o *V2GetPresignedForClusterCredentialsBadRequest) WithPayload(payload *models.Error) *V2GetPresignedForClusterCredentialsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials bad request response
func (o *V2GetPresignedForClusterCredentialsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsUnauthorizedCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsUnauthorized
const V2GetPresignedForClusterCredentialsUnauthorizedCode int = 401

/*V2GetPresignedForClusterCredentialsUnauthorized Unauthorized.

swagger:response v2GetPresignedForClusterCredentialsUnauthorized
*/
type V2GetPresignedForClusterCredentialsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsUnauthorized creates V2GetPresignedForClusterCredentialsUnauthorized with default headers values
func NewV2GetPresignedForClusterCredentialsUnauthorized() *V2GetPresignedForClusterCredentialsUnauthorized {

	return &V2GetPresignedForClusterCredentialsUnauthorized{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials unauthorized response
func (o *V2GetPresignedForClusterCredentialsUnauthorized) WithPayload(payload *models.InfraError) *V2GetPresignedForClusterCredentialsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials unauthorized response
func (o *V2GetPresignedForClusterCredentialsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsForbiddenCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsForbidden
const V2GetPresignedForClusterCredentialsForbiddenCode int = 403

/*V2GetPresignedForClusterCredentialsForbidden Forbidden.

swagger:response v2GetPresignedForClusterCredentialsForbidden
*/
type V2GetPresignedForClusterCredentialsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsForbidden creates V2GetPresignedForClusterCredentialsForbidden with default headers values
func NewV2GetPresignedForClusterCredentialsForbidden() *V2GetPresignedForClusterCredentialsForbidden {

	return &V2GetPresignedForClusterCredentialsForbidden{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials forbidden response
func (o *V2GetPresignedForClusterCredentialsForbidden) WithPayload(payload *models.InfraError) *V2GetPresignedForClusterCredentialsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials forbidden response
func (o *V2GetPresignedForClusterCredentialsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsNotFoundCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsNotFound
const V2GetPresignedForClusterCredentialsNotFoundCode int = 404

/*V2GetPresignedForClusterCredentialsNotFound Error.

swagger:response v2GetPresignedForClusterCredentialsNotFound
*/
type V2GetPresignedForClusterCredentialsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsNotFound creates V2GetPresignedForClusterCredentialsNotFound with default headers values
func NewV2GetPresignedForClusterCredentialsNotFound() *V2GetPresignedForClusterCredentialsNotFound {

	return &V2GetPresignedForClusterCredentialsNotFound{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials not found response
func (o *V2GetPresignedForClusterCredentialsNotFound) WithPayload(payload *models.Error) *V2GetPresignedForClusterCredentialsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials not found response
func (o *V2GetPresignedForClusterCredentialsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsMethodNotAllowedCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsMethodNotAllowed
const V2GetPresignedForClusterCredentialsMethodNotAllowedCode int = 405

/*V2GetPresignedForClusterCredentialsMethodNotAllowed Method Not Allowed.

swagger:response v2GetPresignedForClusterCredentialsMethodNotAllowed
*/
type V2GetPresignedForClusterCredentialsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsMethodNotAllowed creates V2GetPresignedForClusterCredentialsMethodNotAllowed with default headers values
func NewV2GetPresignedForClusterCredentialsMethodNotAllowed() *V2GetPresignedForClusterCredentialsMethodNotAllowed {

	return &V2GetPresignedForClusterCredentialsMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials method not allowed response
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) WithPayload(payload *models.Error) *V2GetPresignedForClusterCredentialsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials method not allowed response
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsConflictCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsConflict
const V2GetPresignedForClusterCredentialsConflictCode int = 409

/*V2GetPresignedForClusterCredentialsConflict Error.

swagger:response v2GetPresignedForClusterCredentialsConflict
*/
type V2GetPresignedForClusterCredentialsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsConflict creates V2GetPresignedForClusterCredentialsConflict with default headers values
func NewV2GetPresignedForClusterCredentialsConflict() *V2GetPresignedForClusterCredentialsConflict {

	return &V2GetPresignedForClusterCredentialsConflict{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials conflict response
func (o *V2GetPresignedForClusterCredentialsConflict) WithPayload(payload *models.Error) *V2GetPresignedForClusterCredentialsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials conflict response
func (o *V2GetPresignedForClusterCredentialsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterCredentialsInternalServerErrorCode is the HTTP code returned for type V2GetPresignedForClusterCredentialsInternalServerError
const V2GetPresignedForClusterCredentialsInternalServerErrorCode int = 500

/*V2GetPresignedForClusterCredentialsInternalServerError Error.

swagger:response v2GetPresignedForClusterCredentialsInternalServerError
*/
type V2GetPresignedForClusterCredentialsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterCredentialsInternalServerError creates V2GetPresignedForClusterCredentialsInternalServerError with default headers values
func NewV2GetPresignedForClusterCredentialsInternalServerError() *V2GetPresignedForClusterCredentialsInternalServerError {

	return &V2GetPresignedForClusterCredentialsInternalServerError{}
}

// WithPayload adds the payload to the v2 get presigned for cluster credentials internal server error response
func (o *V2GetPresignedForClusterCredentialsInternalServerError) WithPayload(payload *models.Error) *V2GetPresignedForClusterCredentialsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster credentials internal server error response
func (o *V2GetPresignedForClusterCredentialsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterCredentialsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
