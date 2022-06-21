// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
)

// PermitTeamUserOKCode is the HTTP code returned for type PermitTeamUserOK
const PermitTeamUserOKCode int = 200

/*PermitTeamUserOK Plain success message

swagger:response permitTeamUserOK
*/
type PermitTeamUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamUserOK creates PermitTeamUserOK with default headers values
func NewPermitTeamUserOK() *PermitTeamUserOK {

	return &PermitTeamUserOK{}
}

// WithPayload adds the payload to the permit team user o k response
func (o *PermitTeamUserOK) WithPayload(payload *models.GeneralError) *PermitTeamUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team user o k response
func (o *PermitTeamUserOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamUserForbiddenCode is the HTTP code returned for type PermitTeamUserForbidden
const PermitTeamUserForbiddenCode int = 403

/*PermitTeamUserForbidden User is not authorized

swagger:response permitTeamUserForbidden
*/
type PermitTeamUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamUserForbidden creates PermitTeamUserForbidden with default headers values
func NewPermitTeamUserForbidden() *PermitTeamUserForbidden {

	return &PermitTeamUserForbidden{}
}

// WithPayload adds the payload to the permit team user forbidden response
func (o *PermitTeamUserForbidden) WithPayload(payload *models.GeneralError) *PermitTeamUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team user forbidden response
func (o *PermitTeamUserForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamUserNotFoundCode is the HTTP code returned for type PermitTeamUserNotFound
const PermitTeamUserNotFoundCode int = 404

/*PermitTeamUserNotFound Team or user not found

swagger:response permitTeamUserNotFound
*/
type PermitTeamUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamUserNotFound creates PermitTeamUserNotFound with default headers values
func NewPermitTeamUserNotFound() *PermitTeamUserNotFound {

	return &PermitTeamUserNotFound{}
}

// WithPayload adds the payload to the permit team user not found response
func (o *PermitTeamUserNotFound) WithPayload(payload *models.GeneralError) *PermitTeamUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team user not found response
func (o *PermitTeamUserNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamUserPreconditionFailedCode is the HTTP code returned for type PermitTeamUserPreconditionFailed
const PermitTeamUserPreconditionFailedCode int = 412

/*PermitTeamUserPreconditionFailed User is not assigned

swagger:response permitTeamUserPreconditionFailed
*/
type PermitTeamUserPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamUserPreconditionFailed creates PermitTeamUserPreconditionFailed with default headers values
func NewPermitTeamUserPreconditionFailed() *PermitTeamUserPreconditionFailed {

	return &PermitTeamUserPreconditionFailed{}
}

// WithPayload adds the payload to the permit team user precondition failed response
func (o *PermitTeamUserPreconditionFailed) WithPayload(payload *models.GeneralError) *PermitTeamUserPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team user precondition failed response
func (o *PermitTeamUserPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamUserPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamUserUnprocessableEntityCode is the HTTP code returned for type PermitTeamUserUnprocessableEntity
const PermitTeamUserUnprocessableEntityCode int = 422

/*PermitTeamUserUnprocessableEntity Failed to validate request

swagger:response permitTeamUserUnprocessableEntity
*/
type PermitTeamUserUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewPermitTeamUserUnprocessableEntity creates PermitTeamUserUnprocessableEntity with default headers values
func NewPermitTeamUserUnprocessableEntity() *PermitTeamUserUnprocessableEntity {

	return &PermitTeamUserUnprocessableEntity{}
}

// WithPayload adds the payload to the permit team user unprocessable entity response
func (o *PermitTeamUserUnprocessableEntity) WithPayload(payload *models.ValidationError) *PermitTeamUserUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team user unprocessable entity response
func (o *PermitTeamUserUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamUserUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PermitTeamUserDefault Some error unrelated to the handler

swagger:response permitTeamUserDefault
*/
type PermitTeamUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamUserDefault creates PermitTeamUserDefault with default headers values
func NewPermitTeamUserDefault(code int) *PermitTeamUserDefault {
	if code <= 0 {
		code = 500
	}

	return &PermitTeamUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the permit team user default response
func (o *PermitTeamUserDefault) WithStatusCode(code int) *PermitTeamUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the permit team user default response
func (o *PermitTeamUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the permit team user default response
func (o *PermitTeamUserDefault) WithPayload(payload *models.GeneralError) *PermitTeamUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team user default response
func (o *PermitTeamUserDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
