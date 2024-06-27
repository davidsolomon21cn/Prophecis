// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"webank/DI/restapi/api_v2/restmodels"
)

// ExportExperimentOKCode is the HTTP code returned for type ExportExperimentOK
const ExportExperimentOKCode int = 200

/*ExportExperimentOK 导出的实验定义

swagger:response exportExperimentOK
*/
type ExportExperimentOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewExportExperimentOK creates ExportExperimentOK with default headers values
func NewExportExperimentOK() *ExportExperimentOK {

	return &ExportExperimentOK{}
}

// WithPayload adds the payload to the export experiment o k response
func (o *ExportExperimentOK) WithPayload(payload io.ReadCloser) *ExportExperimentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export experiment o k response
func (o *ExportExperimentOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportExperimentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ExportExperimentUnauthorizedCode is the HTTP code returned for type ExportExperimentUnauthorized
const ExportExperimentUnauthorizedCode int = 401

/*ExportExperimentUnauthorized Unauthorized

swagger:response exportExperimentUnauthorized
*/
type ExportExperimentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewExportExperimentUnauthorized creates ExportExperimentUnauthorized with default headers values
func NewExportExperimentUnauthorized() *ExportExperimentUnauthorized {

	return &ExportExperimentUnauthorized{}
}

// WithPayload adds the payload to the export experiment unauthorized response
func (o *ExportExperimentUnauthorized) WithPayload(payload *restmodels.Error) *ExportExperimentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export experiment unauthorized response
func (o *ExportExperimentUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportExperimentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ExportExperimentNotFoundCode is the HTTP code returned for type ExportExperimentNotFound
const ExportExperimentNotFoundCode int = 404

/*ExportExperimentNotFound The Experiment cannot be found

swagger:response exportExperimentNotFound
*/
type ExportExperimentNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewExportExperimentNotFound creates ExportExperimentNotFound with default headers values
func NewExportExperimentNotFound() *ExportExperimentNotFound {

	return &ExportExperimentNotFound{}
}

// WithPayload adds the payload to the export experiment not found response
func (o *ExportExperimentNotFound) WithPayload(payload *restmodels.Error) *ExportExperimentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export experiment not found response
func (o *ExportExperimentNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportExperimentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}