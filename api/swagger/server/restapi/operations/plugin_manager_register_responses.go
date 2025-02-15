// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra/api/swagger/server/models"
)

// PluginManagerRegisterNoContentCode is the HTTP code returned for type PluginManagerRegisterNoContent
const PluginManagerRegisterNoContentCode int = 204

/*
PluginManagerRegisterNoContent Plugin successfully installed

swagger:response pluginManagerRegisterNoContent
*/
type PluginManagerRegisterNoContent struct {
}

// NewPluginManagerRegisterNoContent creates PluginManagerRegisterNoContent with default headers values
func NewPluginManagerRegisterNoContent() *PluginManagerRegisterNoContent {

	return &PluginManagerRegisterNoContent{}
}

// WriteResponse to the client
func (o *PluginManagerRegisterNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PluginManagerRegisterBadRequestCode is the HTTP code returned for type PluginManagerRegisterBadRequest
const PluginManagerRegisterBadRequestCode int = 400

/*
PluginManagerRegisterBadRequest Bad request.

swagger:response pluginManagerRegisterBadRequest
*/
type PluginManagerRegisterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerRegisterBadRequest creates PluginManagerRegisterBadRequest with default headers values
func NewPluginManagerRegisterBadRequest() *PluginManagerRegisterBadRequest {

	return &PluginManagerRegisterBadRequest{}
}

// WithPayload adds the payload to the plugin manager register bad request response
func (o *PluginManagerRegisterBadRequest) WithPayload(payload *models.Error) *PluginManagerRegisterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager register bad request response
func (o *PluginManagerRegisterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerRegisterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerRegisterNotFoundCode is the HTTP code returned for type PluginManagerRegisterNotFound
const PluginManagerRegisterNotFoundCode int = 404

/*
PluginManagerRegisterNotFound Not found.

swagger:response pluginManagerRegisterNotFound
*/
type PluginManagerRegisterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerRegisterNotFound creates PluginManagerRegisterNotFound with default headers values
func NewPluginManagerRegisterNotFound() *PluginManagerRegisterNotFound {

	return &PluginManagerRegisterNotFound{}
}

// WithPayload adds the payload to the plugin manager register not found response
func (o *PluginManagerRegisterNotFound) WithPayload(payload *models.Error) *PluginManagerRegisterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager register not found response
func (o *PluginManagerRegisterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerRegisterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerRegisterUnprocessableEntityCode is the HTTP code returned for type PluginManagerRegisterUnprocessableEntity
const PluginManagerRegisterUnprocessableEntityCode int = 422

/*
PluginManagerRegisterUnprocessableEntity Unprocessable Entity - The syntax is correct, but the server was unable to process the contained instructions.

swagger:response pluginManagerRegisterUnprocessableEntity
*/
type PluginManagerRegisterUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerRegisterUnprocessableEntity creates PluginManagerRegisterUnprocessableEntity with default headers values
func NewPluginManagerRegisterUnprocessableEntity() *PluginManagerRegisterUnprocessableEntity {

	return &PluginManagerRegisterUnprocessableEntity{}
}

// WithPayload adds the payload to the plugin manager register unprocessable entity response
func (o *PluginManagerRegisterUnprocessableEntity) WithPayload(payload *models.Error) *PluginManagerRegisterUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager register unprocessable entity response
func (o *PluginManagerRegisterUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerRegisterUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerRegisterInternalServerErrorCode is the HTTP code returned for type PluginManagerRegisterInternalServerError
const PluginManagerRegisterInternalServerErrorCode int = 500

/*
PluginManagerRegisterInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response pluginManagerRegisterInternalServerError
*/
type PluginManagerRegisterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerRegisterInternalServerError creates PluginManagerRegisterInternalServerError with default headers values
func NewPluginManagerRegisterInternalServerError() *PluginManagerRegisterInternalServerError {

	return &PluginManagerRegisterInternalServerError{}
}

// WithPayload adds the payload to the plugin manager register internal server error response
func (o *PluginManagerRegisterInternalServerError) WithPayload(payload *models.Error) *PluginManagerRegisterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager register internal server error response
func (o *PluginManagerRegisterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerRegisterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
