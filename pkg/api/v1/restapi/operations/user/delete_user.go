// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
)

// DeleteUserHandlerFunc turns a function with the right signature into a delete user handler
type DeleteUserHandlerFunc func(DeleteUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserHandlerFunc) Handle(params DeleteUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteUserHandler interface for that can handle valid delete user params
type DeleteUserHandler interface {
	Handle(DeleteUserParams, *models.User) middleware.Responder
}

// NewDeleteUser creates a new http.Handler for the delete user operation
func NewDeleteUser(ctx *middleware.Context, handler DeleteUserHandler) *DeleteUser {
	return &DeleteUser{Context: ctx, Handler: handler}
}

/*
	DeleteUser swagger:route DELETE /users/{user_id} user deleteUser

Delete a specific user
*/
type DeleteUser struct {
	Context *middleware.Context
	Handler DeleteUserHandler
}

func (o *DeleteUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteUserParams()
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
