// Code generated by go-swagger; DO NOT EDIT.

package experiment_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// GetExperimentRunsHistoryReader is a Reader for the GetExperimentRunsHistory structure.
type GetExperimentRunsHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExperimentRunsHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExperimentRunsHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExperimentRunsHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExperimentRunsHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExperimentRunsHistoryOK creates a GetExperimentRunsHistoryOK with default headers values
func NewGetExperimentRunsHistoryOK() *GetExperimentRunsHistoryOK {
	return &GetExperimentRunsHistoryOK{}
}

/*GetExperimentRunsHistoryOK handles this case with default header values.

OK
*/
type GetExperimentRunsHistoryOK struct {
	Payload *restmodels.ProphecisExperiments
}

func (o *GetExperimentRunsHistoryOK) Error() string {
	return fmt.Sprintf("[GET /di/v1/experimentRunsHistory/{exp_id}][%d] getExperimentRunsHistoryOK  %+v", 200, o.Payload)
}

func (o *GetExperimentRunsHistoryOK) GetPayload() *restmodels.ProphecisExperiments {
	return o.Payload
}

func (o *GetExperimentRunsHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.ProphecisExperiments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExperimentRunsHistoryUnauthorized creates a GetExperimentRunsHistoryUnauthorized with default headers values
func NewGetExperimentRunsHistoryUnauthorized() *GetExperimentRunsHistoryUnauthorized {
	return &GetExperimentRunsHistoryUnauthorized{}
}

/*GetExperimentRunsHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetExperimentRunsHistoryUnauthorized struct {
	Payload *restmodels.Error
}

func (o *GetExperimentRunsHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /di/v1/experimentRunsHistory/{exp_id}][%d] getExperimentRunsHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExperimentRunsHistoryUnauthorized) GetPayload() *restmodels.Error {
	return o.Payload
}

func (o *GetExperimentRunsHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExperimentRunsHistoryNotFound creates a GetExperimentRunsHistoryNotFound with default headers values
func NewGetExperimentRunsHistoryNotFound() *GetExperimentRunsHistoryNotFound {
	return &GetExperimentRunsHistoryNotFound{}
}

/*GetExperimentRunsHistoryNotFound handles this case with default header values.

Get ExperimentRun's History Failed
*/
type GetExperimentRunsHistoryNotFound struct {
	Payload *restmodels.Error
}

func (o *GetExperimentRunsHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /di/v1/experimentRunsHistory/{exp_id}][%d] getExperimentRunsHistoryNotFound  %+v", 404, o.Payload)
}

func (o *GetExperimentRunsHistoryNotFound) GetPayload() *restmodels.Error {
	return o.Payload
}

func (o *GetExperimentRunsHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
