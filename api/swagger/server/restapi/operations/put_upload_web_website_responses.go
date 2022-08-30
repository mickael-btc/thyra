// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra/api/swagger/server/models"
)

// PutUploadWebWebsiteOKCode is the HTTP code returned for type PutUploadWebWebsiteOK
const PutUploadWebWebsiteOKCode int = 200

/*PutUploadWebWebsiteOK First website chunk deployed.

swagger:response putUploadWebWebsiteOK
*/
type PutUploadWebWebsiteOK struct {

	/*
	  In: Body
	*/
	Payload *models.Websites `json:"body,omitempty"`
}

// NewPutUploadWebWebsiteOK creates PutUploadWebWebsiteOK with default headers values
func NewPutUploadWebWebsiteOK() *PutUploadWebWebsiteOK {

	return &PutUploadWebWebsiteOK{}
}

// WithPayload adds the payload to the put upload web website o k response
func (o *PutUploadWebWebsiteOK) WithPayload(payload *models.Websites) *PutUploadWebWebsiteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put upload web website o k response
func (o *PutUploadWebWebsiteOK) SetPayload(payload *models.Websites) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUploadWebWebsiteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUploadWebWebsiteBadRequestCode is the HTTP code returned for type PutUploadWebWebsiteBadRequest
const PutUploadWebWebsiteBadRequestCode int = 400

/*PutUploadWebWebsiteBadRequest Bad request.

swagger:response putUploadWebWebsiteBadRequest
*/
type PutUploadWebWebsiteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutUploadWebWebsiteBadRequest creates PutUploadWebWebsiteBadRequest with default headers values
func NewPutUploadWebWebsiteBadRequest() *PutUploadWebWebsiteBadRequest {

	return &PutUploadWebWebsiteBadRequest{}
}

// WithPayload adds the payload to the put upload web website bad request response
func (o *PutUploadWebWebsiteBadRequest) WithPayload(payload *models.Error) *PutUploadWebWebsiteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put upload web website bad request response
func (o *PutUploadWebWebsiteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUploadWebWebsiteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUploadWebWebsiteUnprocessableEntityCode is the HTTP code returned for type PutUploadWebWebsiteUnprocessableEntity
const PutUploadWebWebsiteUnprocessableEntityCode int = 422

/*PutUploadWebWebsiteUnprocessableEntity Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.

swagger:response putUploadWebWebsiteUnprocessableEntity
*/
type PutUploadWebWebsiteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutUploadWebWebsiteUnprocessableEntity creates PutUploadWebWebsiteUnprocessableEntity with default headers values
func NewPutUploadWebWebsiteUnprocessableEntity() *PutUploadWebWebsiteUnprocessableEntity {

	return &PutUploadWebWebsiteUnprocessableEntity{}
}

// WithPayload adds the payload to the put upload web website unprocessable entity response
func (o *PutUploadWebWebsiteUnprocessableEntity) WithPayload(payload *models.Error) *PutUploadWebWebsiteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put upload web website unprocessable entity response
func (o *PutUploadWebWebsiteUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUploadWebWebsiteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUploadWebWebsiteInternalServerErrorCode is the HTTP code returned for type PutUploadWebWebsiteInternalServerError
const PutUploadWebWebsiteInternalServerErrorCode int = 500

/*PutUploadWebWebsiteInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response putUploadWebWebsiteInternalServerError
*/
type PutUploadWebWebsiteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutUploadWebWebsiteInternalServerError creates PutUploadWebWebsiteInternalServerError with default headers values
func NewPutUploadWebWebsiteInternalServerError() *PutUploadWebWebsiteInternalServerError {

	return &PutUploadWebWebsiteInternalServerError{}
}

// WithPayload adds the payload to the put upload web website internal server error response
func (o *PutUploadWebWebsiteInternalServerError) WithPayload(payload *models.Error) *PutUploadWebWebsiteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put upload web website internal server error response
func (o *PutUploadWebWebsiteInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUploadWebWebsiteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
