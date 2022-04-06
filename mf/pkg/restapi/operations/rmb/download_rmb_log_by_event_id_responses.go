// Code generated by go-swagger; DO NOT EDIT.

package rmb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// DownloadRmbLogByEventIDOKCode is the HTTP code returned for type DownloadRmbLogByEventIDOK
const DownloadRmbLogByEventIDOKCode int = 200

/*DownloadRmbLogByEventIDOK OK

swagger:response downloadRmbLogByEventIdOK
*/
type DownloadRmbLogByEventIDOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadRmbLogByEventIDOK creates DownloadRmbLogByEventIDOK with default headers values
func NewDownloadRmbLogByEventIDOK() *DownloadRmbLogByEventIDOK {

	return &DownloadRmbLogByEventIDOK{}
}

// WithPayload adds the payload to the download rmb log by event Id o k response
func (o *DownloadRmbLogByEventIDOK) WithPayload(payload io.ReadCloser) *DownloadRmbLogByEventIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download rmb log by event Id o k response
func (o *DownloadRmbLogByEventIDOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadRmbLogByEventIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadRmbLogByEventIDUnauthorizedCode is the HTTP code returned for type DownloadRmbLogByEventIDUnauthorized
const DownloadRmbLogByEventIDUnauthorizedCode int = 401

/*DownloadRmbLogByEventIDUnauthorized Unauthorized

swagger:response downloadRmbLogByEventIdUnauthorized
*/
type DownloadRmbLogByEventIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadRmbLogByEventIDUnauthorized creates DownloadRmbLogByEventIDUnauthorized with default headers values
func NewDownloadRmbLogByEventIDUnauthorized() *DownloadRmbLogByEventIDUnauthorized {

	return &DownloadRmbLogByEventIDUnauthorized{}
}

// WithPayload adds the payload to the download rmb log by event Id unauthorized response
func (o *DownloadRmbLogByEventIDUnauthorized) WithPayload(payload *models.Error) *DownloadRmbLogByEventIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download rmb log by event Id unauthorized response
func (o *DownloadRmbLogByEventIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadRmbLogByEventIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadRmbLogByEventIDNotFoundCode is the HTTP code returned for type DownloadRmbLogByEventIDNotFound
const DownloadRmbLogByEventIDNotFoundCode int = 404

/*DownloadRmbLogByEventIDNotFound report download fail

swagger:response downloadRmbLogByEventIdNotFound
*/
type DownloadRmbLogByEventIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadRmbLogByEventIDNotFound creates DownloadRmbLogByEventIDNotFound with default headers values
func NewDownloadRmbLogByEventIDNotFound() *DownloadRmbLogByEventIDNotFound {

	return &DownloadRmbLogByEventIDNotFound{}
}

// WithPayload adds the payload to the download rmb log by event Id not found response
func (o *DownloadRmbLogByEventIDNotFound) WithPayload(payload *models.Error) *DownloadRmbLogByEventIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download rmb log by event Id not found response
func (o *DownloadRmbLogByEventIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadRmbLogByEventIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
