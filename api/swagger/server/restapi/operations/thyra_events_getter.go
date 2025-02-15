// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ThyraEventsGetterHandlerFunc turns a function with the right signature into a thyra events getter handler
type ThyraEventsGetterHandlerFunc func(ThyraEventsGetterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ThyraEventsGetterHandlerFunc) Handle(params ThyraEventsGetterParams) middleware.Responder {
	return fn(params)
}

// ThyraEventsGetterHandler interface for that can handle valid thyra events getter params
type ThyraEventsGetterHandler interface {
	Handle(ThyraEventsGetterParams) middleware.Responder
}

// NewThyraEventsGetter creates a new http.Handler for the thyra events getter operation
func NewThyraEventsGetter(ctx *middleware.Context, handler ThyraEventsGetterHandler) *ThyraEventsGetter {
	return &ThyraEventsGetter{Context: ctx, Handler: handler}
}

/*
	ThyraEventsGetter swagger:route GET /thyra/events/{str}/{caller} thyraEventsGetter

ThyraEventsGetter thyra events getter API
*/
type ThyraEventsGetter struct {
	Context *middleware.Context
	Handler ThyraEventsGetterHandler
}

func (o *ThyraEventsGetter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewThyraEventsGetterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
