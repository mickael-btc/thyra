// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MgmtWalletDeleteHandlerFunc turns a function with the right signature into a mgmt wallet delete handler
type MgmtWalletDeleteHandlerFunc func(MgmtWalletDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MgmtWalletDeleteHandlerFunc) Handle(params MgmtWalletDeleteParams) middleware.Responder {
	return fn(params)
}

// MgmtWalletDeleteHandler interface for that can handle valid mgmt wallet delete params
type MgmtWalletDeleteHandler interface {
	Handle(MgmtWalletDeleteParams) middleware.Responder
}

// NewMgmtWalletDelete creates a new http.Handler for the mgmt wallet delete operation
func NewMgmtWalletDelete(ctx *middleware.Context, handler MgmtWalletDeleteHandler) *MgmtWalletDelete {
	return &MgmtWalletDelete{Context: ctx, Handler: handler}
}

/*
	MgmtWalletDelete swagger:route DELETE /mgmt/wallet/{nickname} mgmtWalletDelete

MgmtWalletDelete mgmt wallet delete API
*/
type MgmtWalletDelete struct {
	Context *middleware.Context
	Handler MgmtWalletDeleteHandler
}

func (o *MgmtWalletDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMgmtWalletDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
