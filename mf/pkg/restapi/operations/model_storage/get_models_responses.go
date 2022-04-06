// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// GetModelsOKCode is the HTTP code returned for type GetModelsOK
const GetModelsOKCode int = 200

/*GetModelsOK OK

swagger:response getModelsOK
*/
type GetModelsOK struct {

	/*
	  In: Body
	*/
	Payload models.Models `json:"body,omitempty"`
}

// NewGetModelsOK creates GetModelsOK with default headers values
func NewGetModelsOK() *GetModelsOK {

	return &GetModelsOK{}
}

// WithPayload adds the payload to the get models o k response
func (o *GetModelsOK) WithPayload(payload models.Models) *GetModelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get models o k response
func (o *GetModelsOK) SetPayload(payload models.Models) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Models{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetModelsUnauthorizedCode is the HTTP code returned for type GetModelsUnauthorized
const GetModelsUnauthorizedCode int = 401

/*GetModelsUnauthorized Unauthorized

swagger:response getModelsUnauthorized
*/
type GetModelsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelsUnauthorized creates GetModelsUnauthorized with default headers values
func NewGetModelsUnauthorized() *GetModelsUnauthorized {

	return &GetModelsUnauthorized{}
}

// WithPayload adds the payload to the get models unauthorized response
func (o *GetModelsUnauthorized) WithPayload(payload *models.Error) *GetModelsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get models unauthorized response
func (o *GetModelsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetModelsNotFoundCode is the HTTP code returned for type GetModelsNotFound
const GetModelsNotFoundCode int = 404

/*GetModelsNotFound The Models cannot be found

swagger:response getModelsNotFound
*/
type GetModelsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelsNotFound creates GetModelsNotFound with default headers values
func NewGetModelsNotFound() *GetModelsNotFound {

	return &GetModelsNotFound{}
}

// WithPayload adds the payload to the get models not found response
func (o *GetModelsNotFound) WithPayload(payload *models.Error) *GetModelsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get models not found response
func (o *GetModelsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}