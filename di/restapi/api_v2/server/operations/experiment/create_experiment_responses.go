// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v2/restmodels"
)

// CreateExperimentOKCode is the HTTP code returned for type CreateExperimentOK
const CreateExperimentOKCode int = 200

/*CreateExperimentOK OK

swagger:response createExperimentOK
*/
type CreateExperimentOK struct {

	/*
	  In: Body
	*/
	Payload *restmodels.CreateExperimentResponse `json:"body,omitempty"`
}

// NewCreateExperimentOK creates CreateExperimentOK with default headers values
func NewCreateExperimentOK() *CreateExperimentOK {

	return &CreateExperimentOK{}
}

// WithPayload adds the payload to the create experiment o k response
func (o *CreateExperimentOK) WithPayload(payload *restmodels.CreateExperimentResponse) *CreateExperimentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create experiment o k response
func (o *CreateExperimentOK) SetPayload(payload *restmodels.CreateExperimentResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExperimentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExperimentUnauthorizedCode is the HTTP code returned for type CreateExperimentUnauthorized
const CreateExperimentUnauthorizedCode int = 401

/*CreateExperimentUnauthorized Unauthorized

swagger:response createExperimentUnauthorized
*/
type CreateExperimentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewCreateExperimentUnauthorized creates CreateExperimentUnauthorized with default headers values
func NewCreateExperimentUnauthorized() *CreateExperimentUnauthorized {

	return &CreateExperimentUnauthorized{}
}

// WithPayload adds the payload to the create experiment unauthorized response
func (o *CreateExperimentUnauthorized) WithPayload(payload *restmodels.Error) *CreateExperimentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create experiment unauthorized response
func (o *CreateExperimentUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExperimentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExperimentNotFoundCode is the HTTP code returned for type CreateExperimentNotFound
const CreateExperimentNotFoundCode int = 404

/*CreateExperimentNotFound The Models cannot be found

swagger:response createExperimentNotFound
*/
type CreateExperimentNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewCreateExperimentNotFound creates CreateExperimentNotFound with default headers values
func NewCreateExperimentNotFound() *CreateExperimentNotFound {

	return &CreateExperimentNotFound{}
}

// WithPayload adds the payload to the create experiment not found response
func (o *CreateExperimentNotFound) WithPayload(payload *restmodels.Error) *CreateExperimentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create experiment not found response
func (o *CreateExperimentNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExperimentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
