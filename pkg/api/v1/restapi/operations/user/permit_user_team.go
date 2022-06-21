// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
)

// PermitUserTeamHandlerFunc turns a function with the right signature into a permit user team handler
type PermitUserTeamHandlerFunc func(PermitUserTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PermitUserTeamHandlerFunc) Handle(params PermitUserTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PermitUserTeamHandler interface for that can handle valid permit user team params
type PermitUserTeamHandler interface {
	Handle(PermitUserTeamParams, *models.User) middleware.Responder
}

// NewPermitUserTeam creates a new http.Handler for the permit user team operation
func NewPermitUserTeam(ctx *middleware.Context, handler PermitUserTeamHandler) *PermitUserTeam {
	return &PermitUserTeam{Context: ctx, Handler: handler}
}

/* PermitUserTeam swagger:route PUT /users/{user_id}/teams user permitUserTeam

Update team perms for user

*/
type PermitUserTeam struct {
	Context *middleware.Context
	Handler PermitUserTeamHandler
}

func (o *PermitUserTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPermitUserTeamParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
