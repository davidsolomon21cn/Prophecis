// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// ListImageOKCode is the HTTP code returned for type ListImageOK
const ListImageOKCode int = 200

/*ListImageOK OK

swagger:response listImageOK
*/
type ListImageOK struct {

	/*
	  In: Body
	*/
	Payload models.Images `json:"body,omitempty"`
}

// NewListImageOK creates ListImageOK with default headers values
func NewListImageOK() *ListImageOK {

	return &ListImageOK{}
}

// WithPayload adds the payload to the list image o k response
func (o *ListImageOK) WithPayload(payload models.Images) *ListImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list image o k response
func (o *ListImageOK) SetPayload(payload models.Images) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Images{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListImageUnauthorizedCode is the HTTP code returned for type ListImageUnauthorized
const ListImageUnauthorizedCode int = 401

/*ListImageUnauthorized Unauthorized

swagger:response listImageUnauthorized
*/
type ListImageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListImageUnauthorized creates ListImageUnauthorized with default headers values
func NewListImageUnauthorized() *ListImageUnauthorized {

	return &ListImageUnauthorized{}
}

// WithPayload adds the payload to the list image unauthorized response
func (o *ListImageUnauthorized) WithPayload(payload *models.Error) *ListImageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list image unauthorized response
func (o *ListImageUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListImageNotFoundCode is the HTTP code returned for type ListImageNotFound
const ListImageNotFoundCode int = 404

/*ListImageNotFound The Images cannot be found

swagger:response listImageNotFound
*/
type ListImageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListImageNotFound creates ListImageNotFound with default headers values
func NewListImageNotFound() *ListImageNotFound {

	return &ListImageNotFound{}
}

// WithPayload adds the payload to the list image not found response
func (o *ListImageNotFound) WithPayload(payload *models.Error) *ListImageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list image not found response
func (o *ListImageNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
