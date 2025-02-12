// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra/api/swagger/server/models"
)

// MgmtWalletBalanceOKCode is the HTTP code returned for type MgmtWalletBalanceOK
const MgmtWalletBalanceOKCode int = 200

/*
MgmtWalletBalanceOK Balance retrieved

swagger:response mgmtWalletBalanceOK
*/
type MgmtWalletBalanceOK struct {

	/*
	  In: Body
	*/
	Payload *models.Data `json:"body,omitempty"`
}

// NewMgmtWalletBalanceOK creates MgmtWalletBalanceOK with default headers values
func NewMgmtWalletBalanceOK() *MgmtWalletBalanceOK {

	return &MgmtWalletBalanceOK{}
}

// WithPayload adds the payload to the mgmt wallet balance o k response
func (o *MgmtWalletBalanceOK) WithPayload(payload *models.Data) *MgmtWalletBalanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the mgmt wallet balance o k response
func (o *MgmtWalletBalanceOK) SetPayload(payload *models.Data) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MgmtWalletBalanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MgmtWalletBalanceBadRequestCode is the HTTP code returned for type MgmtWalletBalanceBadRequest
const MgmtWalletBalanceBadRequestCode int = 400

/*
MgmtWalletBalanceBadRequest Bad request.

swagger:response mgmtWalletBalanceBadRequest
*/
type MgmtWalletBalanceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMgmtWalletBalanceBadRequest creates MgmtWalletBalanceBadRequest with default headers values
func NewMgmtWalletBalanceBadRequest() *MgmtWalletBalanceBadRequest {

	return &MgmtWalletBalanceBadRequest{}
}

// WithPayload adds the payload to the mgmt wallet balance bad request response
func (o *MgmtWalletBalanceBadRequest) WithPayload(payload *models.Error) *MgmtWalletBalanceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the mgmt wallet balance bad request response
func (o *MgmtWalletBalanceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MgmtWalletBalanceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MgmtWalletBalanceInternalServerErrorCode is the HTTP code returned for type MgmtWalletBalanceInternalServerError
const MgmtWalletBalanceInternalServerErrorCode int = 500

/*
MgmtWalletBalanceInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response mgmtWalletBalanceInternalServerError
*/
type MgmtWalletBalanceInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMgmtWalletBalanceInternalServerError creates MgmtWalletBalanceInternalServerError with default headers values
func NewMgmtWalletBalanceInternalServerError() *MgmtWalletBalanceInternalServerError {

	return &MgmtWalletBalanceInternalServerError{}
}

// WithPayload adds the payload to the mgmt wallet balance internal server error response
func (o *MgmtWalletBalanceInternalServerError) WithPayload(payload *models.Error) *MgmtWalletBalanceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the mgmt wallet balance internal server error response
func (o *MgmtWalletBalanceInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MgmtWalletBalanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
