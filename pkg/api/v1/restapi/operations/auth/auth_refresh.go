// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AuthRefreshHandlerFunc turns a function with the right signature into a auth refresh handler
type AuthRefreshHandlerFunc func(AuthRefreshParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthRefreshHandlerFunc) Handle(params AuthRefreshParams) middleware.Responder {
	return fn(params)
}

// AuthRefreshHandler interface for that can handle valid auth refresh params
type AuthRefreshHandler interface {
	Handle(AuthRefreshParams) middleware.Responder
}

// NewAuthRefresh creates a new http.Handler for the auth refresh operation
func NewAuthRefresh(ctx *middleware.Context, handler AuthRefreshHandler) *AuthRefresh {
	return &AuthRefresh{Context: ctx, Handler: handler}
}

/*AuthRefresh swagger:route GET /auth/refresh auth authRefresh

Refresh an auth token before it expires

*/
type AuthRefresh struct {
	Context *middleware.Context
	Handler AuthRefreshHandler
}

func (o *AuthRefresh) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAuthRefreshParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
