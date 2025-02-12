// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra/api/swagger/server/models"
)

// MgmtPluginsListOKCode is the HTTP code returned for type MgmtPluginsListOK
const MgmtPluginsListOKCode int = 200

/*
MgmtPluginsListOK Plugins list

swagger:response mgmtPluginsListOK
*/
type MgmtPluginsListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Plugin `json:"body,omitempty"`
}

// NewMgmtPluginsListOK creates MgmtPluginsListOK with default headers values
func NewMgmtPluginsListOK() *MgmtPluginsListOK {

	return &MgmtPluginsListOK{}
}

// WithPayload adds the payload to the mgmt plugins list o k response
func (o *MgmtPluginsListOK) WithPayload(payload []*models.Plugin) *MgmtPluginsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the mgmt plugins list o k response
func (o *MgmtPluginsListOK) SetPayload(payload []*models.Plugin) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MgmtPluginsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Plugin, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// MgmtPluginsListBadRequestCode is the HTTP code returned for type MgmtPluginsListBadRequest
const MgmtPluginsListBadRequestCode int = 400

/*
MgmtPluginsListBadRequest Bad request.

swagger:response mgmtPluginsListBadRequest
*/
type MgmtPluginsListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMgmtPluginsListBadRequest creates MgmtPluginsListBadRequest with default headers values
func NewMgmtPluginsListBadRequest() *MgmtPluginsListBadRequest {

	return &MgmtPluginsListBadRequest{}
}

// WithPayload adds the payload to the mgmt plugins list bad request response
func (o *MgmtPluginsListBadRequest) WithPayload(payload *models.Error) *MgmtPluginsListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the mgmt plugins list bad request response
func (o *MgmtPluginsListBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MgmtPluginsListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MgmtPluginsListInternalServerErrorCode is the HTTP code returned for type MgmtPluginsListInternalServerError
const MgmtPluginsListInternalServerErrorCode int = 500

/*
MgmtPluginsListInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response mgmtPluginsListInternalServerError
*/
type MgmtPluginsListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMgmtPluginsListInternalServerError creates MgmtPluginsListInternalServerError with default headers values
func NewMgmtPluginsListInternalServerError() *MgmtPluginsListInternalServerError {

	return &MgmtPluginsListInternalServerError{}
}

// WithPayload adds the payload to the mgmt plugins list internal server error response
func (o *MgmtPluginsListInternalServerError) WithPayload(payload *models.Error) *MgmtPluginsListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the mgmt plugins list internal server error response
func (o *MgmtPluginsListInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MgmtPluginsListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
