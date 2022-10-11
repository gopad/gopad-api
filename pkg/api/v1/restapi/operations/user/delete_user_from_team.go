// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
)

// DeleteUserFromTeamHandlerFunc turns a function with the right signature into a delete user from team handler
type DeleteUserFromTeamHandlerFunc func(DeleteUserFromTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserFromTeamHandlerFunc) Handle(params DeleteUserFromTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteUserFromTeamHandler interface for that can handle valid delete user from team params
type DeleteUserFromTeamHandler interface {
	Handle(DeleteUserFromTeamParams, *models.User) middleware.Responder
}

// NewDeleteUserFromTeam creates a new http.Handler for the delete user from team operation
func NewDeleteUserFromTeam(ctx *middleware.Context, handler DeleteUserFromTeamHandler) *DeleteUserFromTeam {
	return &DeleteUserFromTeam{Context: ctx, Handler: handler}
}

/*
	DeleteUserFromTeam swagger:route DELETE /users/{user_id}/teams user deleteUserFromTeam

Remove a team from user
*/
type DeleteUserFromTeam struct {
	Context *middleware.Context
	Handler DeleteUserFromTeamHandler
}

func (o *DeleteUserFromTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteUserFromTeamParams()
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
