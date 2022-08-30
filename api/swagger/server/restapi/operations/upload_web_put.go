// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadWebPutHandlerFunc turns a function with the right signature into a upload web put handler
type UploadWebPutHandlerFunc func(UploadWebPutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadWebPutHandlerFunc) Handle(params UploadWebPutParams) middleware.Responder {
	return fn(params)
}

// UploadWebPutHandler interface for that can handle valid upload web put params
type UploadWebPutHandler interface {
	Handle(UploadWebPutParams) middleware.Responder
}

// NewUploadWebPut creates a new http.Handler for the upload web put operation
func NewUploadWebPut(ctx *middleware.Context, handler UploadWebPutHandler) *UploadWebPut {
	return &UploadWebPut{Context: ctx, Handler: handler}
}

/* UploadWebPut swagger:route PUT /uploadWeb/{website} uploadWebPut

UploadWebPut upload web put API

*/
type UploadWebPut struct {
	Context *middleware.Context
	Handler UploadWebPutHandler
}

func (o *UploadWebPut) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUploadWebPutParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
