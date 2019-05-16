// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gopad/gopad-api/pkg/api/v1/models"
)

// TeamUserPermOKCode is the HTTP code returned for type TeamUserPermOK
const TeamUserPermOKCode int = 200

/*TeamUserPermOK Plain success message

swagger:response teamUserPermOK
*/
type TeamUserPermOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserPermOK creates TeamUserPermOK with default headers values
func NewTeamUserPermOK() *TeamUserPermOK {

	return &TeamUserPermOK{}
}

// WithPayload adds the payload to the team user perm o k response
func (o *TeamUserPermOK) WithPayload(payload *models.GeneralError) *TeamUserPermOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user perm o k response
func (o *TeamUserPermOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserPermOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserPermForbiddenCode is the HTTP code returned for type TeamUserPermForbidden
const TeamUserPermForbiddenCode int = 403

/*TeamUserPermForbidden User is not authorized

swagger:response teamUserPermForbidden
*/
type TeamUserPermForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserPermForbidden creates TeamUserPermForbidden with default headers values
func NewTeamUserPermForbidden() *TeamUserPermForbidden {

	return &TeamUserPermForbidden{}
}

// WithPayload adds the payload to the team user perm forbidden response
func (o *TeamUserPermForbidden) WithPayload(payload *models.GeneralError) *TeamUserPermForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user perm forbidden response
func (o *TeamUserPermForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserPermForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserPermPreconditionFailedCode is the HTTP code returned for type TeamUserPermPreconditionFailed
const TeamUserPermPreconditionFailedCode int = 412

/*TeamUserPermPreconditionFailed Failed to parse request body

swagger:response teamUserPermPreconditionFailed
*/
type TeamUserPermPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserPermPreconditionFailed creates TeamUserPermPreconditionFailed with default headers values
func NewTeamUserPermPreconditionFailed() *TeamUserPermPreconditionFailed {

	return &TeamUserPermPreconditionFailed{}
}

// WithPayload adds the payload to the team user perm precondition failed response
func (o *TeamUserPermPreconditionFailed) WithPayload(payload *models.GeneralError) *TeamUserPermPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user perm precondition failed response
func (o *TeamUserPermPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserPermPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserPermUnprocessableEntityCode is the HTTP code returned for type TeamUserPermUnprocessableEntity
const TeamUserPermUnprocessableEntityCode int = 422

/*TeamUserPermUnprocessableEntity User is not assigned

swagger:response teamUserPermUnprocessableEntity
*/
type TeamUserPermUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserPermUnprocessableEntity creates TeamUserPermUnprocessableEntity with default headers values
func NewTeamUserPermUnprocessableEntity() *TeamUserPermUnprocessableEntity {

	return &TeamUserPermUnprocessableEntity{}
}

// WithPayload adds the payload to the team user perm unprocessable entity response
func (o *TeamUserPermUnprocessableEntity) WithPayload(payload *models.GeneralError) *TeamUserPermUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user perm unprocessable entity response
func (o *TeamUserPermUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserPermUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamUserPermDefault Some error unrelated to the handler

swagger:response teamUserPermDefault
*/
type TeamUserPermDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserPermDefault creates TeamUserPermDefault with default headers values
func NewTeamUserPermDefault(code int) *TeamUserPermDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamUserPermDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team user perm default response
func (o *TeamUserPermDefault) WithStatusCode(code int) *TeamUserPermDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team user perm default response
func (o *TeamUserPermDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team user perm default response
func (o *TeamUserPermDefault) WithPayload(payload *models.GeneralError) *TeamUserPermDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user perm default response
func (o *TeamUserPermDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserPermDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
