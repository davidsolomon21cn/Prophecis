// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// PatchModelAcceptedCode is the HTTP code returned for type PatchModelAccepted
const PatchModelAcceptedCode int = 202

/*PatchModelAccepted Training successfully halted.

swagger:response patchModelAccepted
*/
type PatchModelAccepted struct {

	/*
	  In: Body
	*/
	Payload *restmodels.BasicModel `json:"body,omitempty"`
}

// NewPatchModelAccepted creates PatchModelAccepted with default headers values
func NewPatchModelAccepted() *PatchModelAccepted {

	return &PatchModelAccepted{}
}

// WithPayload adds the payload to the patch model accepted response
func (o *PatchModelAccepted) WithPayload(payload *restmodels.BasicModel) *PatchModelAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch model accepted response
func (o *PatchModelAccepted) SetPayload(payload *restmodels.BasicModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchModelAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchModelBadRequestCode is the HTTP code returned for type PatchModelBadRequest
const PatchModelBadRequestCode int = 400

/*PatchModelBadRequest Incorrect status specified.

swagger:response patchModelBadRequest
*/
type PatchModelBadRequest struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewPatchModelBadRequest creates PatchModelBadRequest with default headers values
func NewPatchModelBadRequest() *PatchModelBadRequest {

	return &PatchModelBadRequest{}
}

// WithPayload adds the payload to the patch model bad request response
func (o *PatchModelBadRequest) WithPayload(payload *restmodels.Error) *PatchModelBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch model bad request response
func (o *PatchModelBadRequest) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchModelBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchModelUnauthorizedCode is the HTTP code returned for type PatchModelUnauthorized
const PatchModelUnauthorizedCode int = 401

/*PatchModelUnauthorized Unauthorized

swagger:response patchModelUnauthorized
*/
type PatchModelUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewPatchModelUnauthorized creates PatchModelUnauthorized with default headers values
func NewPatchModelUnauthorized() *PatchModelUnauthorized {

	return &PatchModelUnauthorized{}
}

// WithPayload adds the payload to the patch model unauthorized response
func (o *PatchModelUnauthorized) WithPayload(payload *restmodels.Error) *PatchModelUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch model unauthorized response
func (o *PatchModelUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchModelUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchModelNotFoundCode is the HTTP code returned for type PatchModelNotFound
const PatchModelNotFoundCode int = 404

/*PatchModelNotFound Model with the given ID not found.

swagger:response patchModelNotFound
*/
type PatchModelNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewPatchModelNotFound creates PatchModelNotFound with default headers values
func NewPatchModelNotFound() *PatchModelNotFound {

	return &PatchModelNotFound{}
}

// WithPayload adds the payload to the patch model not found response
func (o *PatchModelNotFound) WithPayload(payload *restmodels.Error) *PatchModelNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch model not found response
func (o *PatchModelNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchModelNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
