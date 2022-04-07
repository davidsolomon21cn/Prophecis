// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// GetModelOKCode is the HTTP code returned for type GetModelOK
const GetModelOKCode int = 200

/*GetModelOK OK

swagger:response getModelOK
*/
type GetModelOK struct {

	/*
	  In: Body
	*/
	Payload *models.Model `json:"body,omitempty"`
}

// NewGetModelOK creates GetModelOK with default headers values
func NewGetModelOK() *GetModelOK {

	return &GetModelOK{}
}

// WithPayload adds the payload to the get model o k response
func (o *GetModelOK) WithPayload(payload *models.Model) *GetModelOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get model o k response
func (o *GetModelOK) SetPayload(payload *models.Model) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetModelUnauthorizedCode is the HTTP code returned for type GetModelUnauthorized
const GetModelUnauthorizedCode int = 401

/*GetModelUnauthorized Unauthorized

swagger:response getModelUnauthorized
*/
type GetModelUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelUnauthorized creates GetModelUnauthorized with default headers values
func NewGetModelUnauthorized() *GetModelUnauthorized {

	return &GetModelUnauthorized{}
}

// WithPayload adds the payload to the get model unauthorized response
func (o *GetModelUnauthorized) WithPayload(payload *models.Error) *GetModelUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get model unauthorized response
func (o *GetModelUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetModelNotFoundCode is the HTTP code returned for type GetModelNotFound
const GetModelNotFoundCode int = 404

/*GetModelNotFound Model create failed

swagger:response getModelNotFound
*/
type GetModelNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelNotFound creates GetModelNotFound with default headers values
func NewGetModelNotFound() *GetModelNotFound {

	return &GetModelNotFound{}
}

// WithPayload adds the payload to the get model not found response
func (o *GetModelNotFound) WithPayload(payload *models.Error) *GetModelNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get model not found response
func (o *GetModelNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
