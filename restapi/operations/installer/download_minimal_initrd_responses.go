// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadMinimalInitrdOKCode is the HTTP code returned for type DownloadMinimalInitrdOK
const DownloadMinimalInitrdOKCode int = 200

/*DownloadMinimalInitrdOK Success.

swagger:response downloadMinimalInitrdOK
*/
type DownloadMinimalInitrdOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdOK creates DownloadMinimalInitrdOK with default headers values
func NewDownloadMinimalInitrdOK() *DownloadMinimalInitrdOK {

	return &DownloadMinimalInitrdOK{}
}

// WithPayload adds the payload to the download minimal initrd o k response
func (o *DownloadMinimalInitrdOK) WithPayload(payload io.ReadCloser) *DownloadMinimalInitrdOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd o k response
func (o *DownloadMinimalInitrdOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadMinimalInitrdNoContentCode is the HTTP code returned for type DownloadMinimalInitrdNoContent
const DownloadMinimalInitrdNoContentCode int = 204

/*DownloadMinimalInitrdNoContent Empty Success.

swagger:response downloadMinimalInitrdNoContent
*/
type DownloadMinimalInitrdNoContent struct {
}

// NewDownloadMinimalInitrdNoContent creates DownloadMinimalInitrdNoContent with default headers values
func NewDownloadMinimalInitrdNoContent() *DownloadMinimalInitrdNoContent {

	return &DownloadMinimalInitrdNoContent{}
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DownloadMinimalInitrdUnauthorizedCode is the HTTP code returned for type DownloadMinimalInitrdUnauthorized
const DownloadMinimalInitrdUnauthorizedCode int = 401

/*DownloadMinimalInitrdUnauthorized Unauthorized.

swagger:response downloadMinimalInitrdUnauthorized
*/
type DownloadMinimalInitrdUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdUnauthorized creates DownloadMinimalInitrdUnauthorized with default headers values
func NewDownloadMinimalInitrdUnauthorized() *DownloadMinimalInitrdUnauthorized {

	return &DownloadMinimalInitrdUnauthorized{}
}

// WithPayload adds the payload to the download minimal initrd unauthorized response
func (o *DownloadMinimalInitrdUnauthorized) WithPayload(payload *models.InfraError) *DownloadMinimalInitrdUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd unauthorized response
func (o *DownloadMinimalInitrdUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadMinimalInitrdForbiddenCode is the HTTP code returned for type DownloadMinimalInitrdForbidden
const DownloadMinimalInitrdForbiddenCode int = 403

/*DownloadMinimalInitrdForbidden Forbidden.

swagger:response downloadMinimalInitrdForbidden
*/
type DownloadMinimalInitrdForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdForbidden creates DownloadMinimalInitrdForbidden with default headers values
func NewDownloadMinimalInitrdForbidden() *DownloadMinimalInitrdForbidden {

	return &DownloadMinimalInitrdForbidden{}
}

// WithPayload adds the payload to the download minimal initrd forbidden response
func (o *DownloadMinimalInitrdForbidden) WithPayload(payload *models.InfraError) *DownloadMinimalInitrdForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd forbidden response
func (o *DownloadMinimalInitrdForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadMinimalInitrdNotFoundCode is the HTTP code returned for type DownloadMinimalInitrdNotFound
const DownloadMinimalInitrdNotFoundCode int = 404

/*DownloadMinimalInitrdNotFound Error.

swagger:response downloadMinimalInitrdNotFound
*/
type DownloadMinimalInitrdNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdNotFound creates DownloadMinimalInitrdNotFound with default headers values
func NewDownloadMinimalInitrdNotFound() *DownloadMinimalInitrdNotFound {

	return &DownloadMinimalInitrdNotFound{}
}

// WithPayload adds the payload to the download minimal initrd not found response
func (o *DownloadMinimalInitrdNotFound) WithPayload(payload *models.Error) *DownloadMinimalInitrdNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd not found response
func (o *DownloadMinimalInitrdNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadMinimalInitrdMethodNotAllowedCode is the HTTP code returned for type DownloadMinimalInitrdMethodNotAllowed
const DownloadMinimalInitrdMethodNotAllowedCode int = 405

/*DownloadMinimalInitrdMethodNotAllowed Method Not Allowed.

swagger:response downloadMinimalInitrdMethodNotAllowed
*/
type DownloadMinimalInitrdMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdMethodNotAllowed creates DownloadMinimalInitrdMethodNotAllowed with default headers values
func NewDownloadMinimalInitrdMethodNotAllowed() *DownloadMinimalInitrdMethodNotAllowed {

	return &DownloadMinimalInitrdMethodNotAllowed{}
}

// WithPayload adds the payload to the download minimal initrd method not allowed response
func (o *DownloadMinimalInitrdMethodNotAllowed) WithPayload(payload *models.Error) *DownloadMinimalInitrdMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd method not allowed response
func (o *DownloadMinimalInitrdMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadMinimalInitrdConflictCode is the HTTP code returned for type DownloadMinimalInitrdConflict
const DownloadMinimalInitrdConflictCode int = 409

/*DownloadMinimalInitrdConflict Conflict.

swagger:response downloadMinimalInitrdConflict
*/
type DownloadMinimalInitrdConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdConflict creates DownloadMinimalInitrdConflict with default headers values
func NewDownloadMinimalInitrdConflict() *DownloadMinimalInitrdConflict {

	return &DownloadMinimalInitrdConflict{}
}

// WithPayload adds the payload to the download minimal initrd conflict response
func (o *DownloadMinimalInitrdConflict) WithPayload(payload *models.Error) *DownloadMinimalInitrdConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd conflict response
func (o *DownloadMinimalInitrdConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadMinimalInitrdInternalServerErrorCode is the HTTP code returned for type DownloadMinimalInitrdInternalServerError
const DownloadMinimalInitrdInternalServerErrorCode int = 500

/*DownloadMinimalInitrdInternalServerError Error.

swagger:response downloadMinimalInitrdInternalServerError
*/
type DownloadMinimalInitrdInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadMinimalInitrdInternalServerError creates DownloadMinimalInitrdInternalServerError with default headers values
func NewDownloadMinimalInitrdInternalServerError() *DownloadMinimalInitrdInternalServerError {

	return &DownloadMinimalInitrdInternalServerError{}
}

// WithPayload adds the payload to the download minimal initrd internal server error response
func (o *DownloadMinimalInitrdInternalServerError) WithPayload(payload *models.Error) *DownloadMinimalInitrdInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download minimal initrd internal server error response
func (o *DownloadMinimalInitrdInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadMinimalInitrdInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
