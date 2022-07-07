// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// UpdateExperimentInfoOKCode is the HTTP code returned for type UpdateExperimentInfoOK
const UpdateExperimentInfoOKCode int = 200

/*UpdateExperimentInfoOK OK

swagger:response updateExperimentInfoOK
*/
type UpdateExperimentInfoOK struct {
}

// NewUpdateExperimentInfoOK creates UpdateExperimentInfoOK with default headers values
func NewUpdateExperimentInfoOK() *UpdateExperimentInfoOK {

	return &UpdateExperimentInfoOK{}
}

// WriteResponse to the client
func (o *UpdateExperimentInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateExperimentInfoUnauthorizedCode is the HTTP code returned for type UpdateExperimentInfoUnauthorized
const UpdateExperimentInfoUnauthorizedCode int = 401

/*UpdateExperimentInfoUnauthorized Unauthorized

swagger:response updateExperimentInfoUnauthorized
*/
type UpdateExperimentInfoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewUpdateExperimentInfoUnauthorized creates UpdateExperimentInfoUnauthorized with default headers values
func NewUpdateExperimentInfoUnauthorized() *UpdateExperimentInfoUnauthorized {

	return &UpdateExperimentInfoUnauthorized{}
}

// WithPayload adds the payload to the update experiment info unauthorized response
func (o *UpdateExperimentInfoUnauthorized) WithPayload(payload *restmodels.Error) *UpdateExperimentInfoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update experiment info unauthorized response
func (o *UpdateExperimentInfoUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExperimentInfoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateExperimentInfoNotFoundCode is the HTTP code returned for type UpdateExperimentInfoNotFound
const UpdateExperimentInfoNotFoundCode int = 404

/*UpdateExperimentInfoNotFound Model create failed

swagger:response updateExperimentInfoNotFound
*/
type UpdateExperimentInfoNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewUpdateExperimentInfoNotFound creates UpdateExperimentInfoNotFound with default headers values
func NewUpdateExperimentInfoNotFound() *UpdateExperimentInfoNotFound {

	return &UpdateExperimentInfoNotFound{}
}

// WithPayload adds the payload to the update experiment info not found response
func (o *UpdateExperimentInfoNotFound) WithPayload(payload *restmodels.Error) *UpdateExperimentInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update experiment info not found response
func (o *UpdateExperimentInfoNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateExperimentInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}