// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PluginRouterHandlerFunc turns a function with the right signature into a plugin router handler
type PluginRouterHandlerFunc func(PluginRouterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PluginRouterHandlerFunc) Handle(params PluginRouterParams) middleware.Responder {
	return fn(params)
}

// PluginRouterHandler interface for that can handle valid plugin router params
type PluginRouterHandler interface {
	Handle(PluginRouterParams) middleware.Responder
}

// NewPluginRouter creates a new http.Handler for the plugin router operation
func NewPluginRouter(ctx *middleware.Context, handler PluginRouterHandler) *PluginRouter {
	return &PluginRouter{Context: ctx, Handler: handler}
}

/*
	PluginRouter swagger:route GET /thyra/plugin/{author-name}/{plugin-name} pluginRouter

virtual endpoint handling requests for third party plugin. The actual handler is defined as an HTTP handler middleware.
*/
type PluginRouter struct {
	Context *middleware.Context
	Handler PluginRouterHandler
}

func (o *PluginRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPluginRouterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
