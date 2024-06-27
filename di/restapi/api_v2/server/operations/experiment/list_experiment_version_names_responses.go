// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v2/restmodels"
)

// ListExperimentVersionNamesOKCode is the HTTP code returned for type ListExperimentVersionNamesOK
const ListExperimentVersionNamesOKCode int = 200

/*ListExperimentVersionNamesOK OK

swagger:response listExperimentVersionNamesOK
*/
type ListExperimentVersionNamesOK struct {

	/*实验所有可用版本名称
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewListExperimentVersionNamesOK creates ListExperimentVersionNamesOK with default headers values
func NewListExperimentVersionNamesOK() *ListExperimentVersionNamesOK {

	return &ListExperimentVersionNamesOK{}
}

// WithPayload adds the payload to the list experiment version names o k response
func (o *ListExperimentVersionNamesOK) WithPayload(payload []string) *ListExperimentVersionNamesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list experiment version names o k response
func (o *ListExperimentVersionNamesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExperimentVersionNamesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListExperimentVersionNamesUnauthorizedCode is the HTTP code returned for type ListExperimentVersionNamesUnauthorized
const ListExperimentVersionNamesUnauthorizedCode int = 401

/*ListExperimentVersionNamesUnauthorized Unauthorized

swagger:response listExperimentVersionNamesUnauthorized
*/
type ListExperimentVersionNamesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewListExperimentVersionNamesUnauthorized creates ListExperimentVersionNamesUnauthorized with default headers values
func NewListExperimentVersionNamesUnauthorized() *ListExperimentVersionNamesUnauthorized {

	return &ListExperimentVersionNamesUnauthorized{}
}

// WithPayload adds the payload to the list experiment version names unauthorized response
func (o *ListExperimentVersionNamesUnauthorized) WithPayload(payload *restmodels.Error) *ListExperimentVersionNamesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list experiment version names unauthorized response
func (o *ListExperimentVersionNamesUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExperimentVersionNamesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListExperimentVersionNamesNotFoundCode is the HTTP code returned for type ListExperimentVersionNamesNotFound
const ListExperimentVersionNamesNotFoundCode int = 404

/*ListExperimentVersionNamesNotFound Model create failed

swagger:response listExperimentVersionNamesNotFound
*/
type ListExperimentVersionNamesNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewListExperimentVersionNamesNotFound creates ListExperimentVersionNamesNotFound with default headers values
func NewListExperimentVersionNamesNotFound() *ListExperimentVersionNamesNotFound {

	return &ListExperimentVersionNamesNotFound{}
}

// WithPayload adds the payload to the list experiment version names not found response
func (o *ListExperimentVersionNamesNotFound) WithPayload(payload *restmodels.Error) *ListExperimentVersionNamesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list experiment version names not found response
func (o *ListExperimentVersionNamesNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExperimentVersionNamesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
