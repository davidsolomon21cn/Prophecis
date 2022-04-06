// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// GetModelsByClusterOKCode is the HTTP code returned for type GetModelsByClusterOK
const GetModelsByClusterOKCode int = 200

/*GetModelsByClusterOK OK

swagger:response getModelsByClusterOK
*/
type GetModelsByClusterOK struct {

	/*
	  In: Body
	*/
	Payload models.Models `json:"body,omitempty"`
}

// NewGetModelsByClusterOK creates GetModelsByClusterOK with default headers values
func NewGetModelsByClusterOK() *GetModelsByClusterOK {

	return &GetModelsByClusterOK{}
}

// WithPayload adds the payload to the get models by cluster o k response
func (o *GetModelsByClusterOK) WithPayload(payload models.Models) *GetModelsByClusterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get models by cluster o k response
func (o *GetModelsByClusterOK) SetPayload(payload models.Models) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelsByClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetModelsByClusterUnauthorizedCode is the HTTP code returned for type GetModelsByClusterUnauthorized
const GetModelsByClusterUnauthorizedCode int = 401

/*GetModelsByClusterUnauthorized Unauthorized

swagger:response getModelsByClusterUnauthorized
*/
type GetModelsByClusterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelsByClusterUnauthorized creates GetModelsByClusterUnauthorized with default headers values
func NewGetModelsByClusterUnauthorized() *GetModelsByClusterUnauthorized {

	return &GetModelsByClusterUnauthorized{}
}

// WithPayload adds the payload to the get models by cluster unauthorized response
func (o *GetModelsByClusterUnauthorized) WithPayload(payload *models.Error) *GetModelsByClusterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get models by cluster unauthorized response
func (o *GetModelsByClusterUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelsByClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetModelsByClusterNotFoundCode is the HTTP code returned for type GetModelsByClusterNotFound
const GetModelsByClusterNotFoundCode int = 404

/*GetModelsByClusterNotFound The Models cannot be found

swagger:response getModelsByClusterNotFound
*/
type GetModelsByClusterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelsByClusterNotFound creates GetModelsByClusterNotFound with default headers values
func NewGetModelsByClusterNotFound() *GetModelsByClusterNotFound {

	return &GetModelsByClusterNotFound{}
}

// WithPayload adds the payload to the get models by cluster not found response
func (o *GetModelsByClusterNotFound) WithPayload(payload *models.Error) *GetModelsByClusterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get models by cluster not found response
func (o *GetModelsByClusterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelsByClusterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
