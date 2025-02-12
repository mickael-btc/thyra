// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ThyraWalletHandlerFunc turns a function with the right signature into a thyra wallet handler
type ThyraWalletHandlerFunc func(ThyraWalletParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ThyraWalletHandlerFunc) Handle(params ThyraWalletParams) middleware.Responder {
	return fn(params)
}

// ThyraWalletHandler interface for that can handle valid thyra wallet params
type ThyraWalletHandler interface {
	Handle(ThyraWalletParams) middleware.Responder
}

// NewThyraWallet creates a new http.Handler for the thyra wallet operation
func NewThyraWallet(ctx *middleware.Context, handler ThyraWalletHandler) *ThyraWallet {
	return &ThyraWallet{Context: ctx, Handler: handler}
}

/*
	ThyraWallet swagger:route GET /thyra/wallet/{resource} thyraWallet

ThyraWallet thyra wallet API
*/
type ThyraWallet struct {
	Context *middleware.Context
	Handler ThyraWalletHandler
}

func (o *ThyraWallet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewThyraWalletParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
