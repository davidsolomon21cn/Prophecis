// Code generated by go-swagger; DO NOT EDIT.

package experiment_run

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v2/restmodels"
)

// RetryExperimentRunOKCode is the HTTP code returned for type RetryExperimentRunOK
const RetryExperimentRunOKCode int = 200

/*RetryExperimentRunOK OK

swagger:response retryExperimentRunOK
*/
type RetryExperimentRunOK struct {

	/*
	  In: Body
	*/
	Payload *restmodels.CreateExperimentRunResponse `json:"body,omitempty"`
}

// NewRetryExperimentRunOK creates RetryExperimentRunOK with default headers values
func NewRetryExperimentRunOK() *RetryExperimentRunOK {

	return &RetryExperimentRunOK{}
}

// WithPayload adds the payload to the retry experiment run o k response
func (o *RetryExperimentRunOK) WithPayload(payload *restmodels.CreateExperimentRunResponse) *RetryExperimentRunOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retry experiment run o k response
func (o *RetryExperimentRunOK) SetPayload(payload *restmodels.CreateExperimentRunResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetryExperimentRunOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetryExperimentRunUnauthorizedCode is the HTTP code returned for type RetryExperimentRunUnauthorized
const RetryExperimentRunUnauthorizedCode int = 401

/*RetryExperimentRunUnauthorized Unauthorized

swagger:response retryExperimentRunUnauthorized
*/
type RetryExperimentRunUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewRetryExperimentRunUnauthorized creates RetryExperimentRunUnauthorized with default headers values
func NewRetryExperimentRunUnauthorized() *RetryExperimentRunUnauthorized {

	return &RetryExperimentRunUnauthorized{}
}

// WithPayload adds the payload to the retry experiment run unauthorized response
func (o *RetryExperimentRunUnauthorized) WithPayload(payload *restmodels.Error) *RetryExperimentRunUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retry experiment run unauthorized response
func (o *RetryExperimentRunUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetryExperimentRunUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetryExperimentRunNotFoundCode is the HTTP code returned for type RetryExperimentRunNotFound
const RetryExperimentRunNotFoundCode int = 404

/*RetryExperimentRunNotFound The Models cannot be found

swagger:response retryExperimentRunNotFound
*/
type RetryExperimentRunNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewRetryExperimentRunNotFound creates RetryExperimentRunNotFound with default headers values
func NewRetryExperimentRunNotFound() *RetryExperimentRunNotFound {

	return &RetryExperimentRunNotFound{}
}

// WithPayload adds the payload to the retry experiment run not found response
func (o *RetryExperimentRunNotFound) WithPayload(payload *restmodels.Error) *RetryExperimentRunNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retry experiment run not found response
func (o *RetryExperimentRunNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetryExperimentRunNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
