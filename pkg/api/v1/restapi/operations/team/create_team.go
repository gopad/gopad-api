// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
)

// CreateTeamHandlerFunc turns a function with the right signature into a create team handler
type CreateTeamHandlerFunc func(CreateTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTeamHandlerFunc) Handle(params CreateTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CreateTeamHandler interface for that can handle valid create team params
type CreateTeamHandler interface {
	Handle(CreateTeamParams, *models.User) middleware.Responder
}

// NewCreateTeam creates a new http.Handler for the create team operation
func NewCreateTeam(ctx *middleware.Context, handler CreateTeamHandler) *CreateTeam {
	return &CreateTeam{Context: ctx, Handler: handler}
}

/* CreateTeam swagger:route POST /teams team createTeam

Create a new team

*/
type CreateTeam struct {
	Context *middleware.Context
	Handler CreateTeamHandler
}

func (o *CreateTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateTeamParams()
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
