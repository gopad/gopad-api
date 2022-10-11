// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
)

// ShowProfileOKCode is the HTTP code returned for type ShowProfileOK
const ShowProfileOKCode int = 200

/*
ShowProfileOK The current profile data

swagger:response showProfileOK
*/
type ShowProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewShowProfileOK creates ShowProfileOK with default headers values
func NewShowProfileOK() *ShowProfileOK {

	return &ShowProfileOK{}
}

// WithPayload adds the payload to the show profile o k response
func (o *ShowProfileOK) WithPayload(payload *models.Profile) *ShowProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show profile o k response
func (o *ShowProfileOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowProfileForbiddenCode is the HTTP code returned for type ShowProfileForbidden
const ShowProfileForbiddenCode int = 403

/*
ShowProfileForbidden User is not authorized

swagger:response showProfileForbidden
*/
type ShowProfileForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewShowProfileForbidden creates ShowProfileForbidden with default headers values
func NewShowProfileForbidden() *ShowProfileForbidden {

	return &ShowProfileForbidden{}
}

// WithPayload adds the payload to the show profile forbidden response
func (o *ShowProfileForbidden) WithPayload(payload *models.GeneralError) *ShowProfileForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show profile forbidden response
func (o *ShowProfileForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowProfileForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ShowProfileDefault Some error unrelated to the handler

swagger:response showProfileDefault
*/
type ShowProfileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewShowProfileDefault creates ShowProfileDefault with default headers values
func NewShowProfileDefault(code int) *ShowProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &ShowProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the show profile default response
func (o *ShowProfileDefault) WithStatusCode(code int) *ShowProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the show profile default response
func (o *ShowProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the show profile default response
func (o *ShowProfileDefault) WithPayload(payload *models.GeneralError) *ShowProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show profile default response
func (o *ShowProfileDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
