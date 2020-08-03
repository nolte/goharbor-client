// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/model/v1_10_0"
)

// GetSystemGcScheduleReader is a Reader for the GetSystemGcSchedule structure.
type GetSystemGcScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemGcScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemGcScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSystemGcScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSystemGcScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSystemGcScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemGcScheduleOK creates a GetSystemGcScheduleOK with default headers values
func NewGetSystemGcScheduleOK() *GetSystemGcScheduleOK {
	return &GetSystemGcScheduleOK{}
}

/*GetSystemGcScheduleOK handles this case with default header values.

Get gc's schedule.
*/
type GetSystemGcScheduleOK struct {
	Payload *v1_10_0.AdminJobSchedule
}

func (o *GetSystemGcScheduleOK) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getSystemGcScheduleOK  %+v", 200, o.Payload)
}

func (o *GetSystemGcScheduleOK) GetPayload() *v1_10_0.AdminJobSchedule {
	return o.Payload
}

func (o *GetSystemGcScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1_10_0.AdminJobSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemGcScheduleUnauthorized creates a GetSystemGcScheduleUnauthorized with default headers values
func NewGetSystemGcScheduleUnauthorized() *GetSystemGcScheduleUnauthorized {
	return &GetSystemGcScheduleUnauthorized{}
}

/*GetSystemGcScheduleUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetSystemGcScheduleUnauthorized struct {
}

func (o *GetSystemGcScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getSystemGcScheduleUnauthorized ", 401)
}

func (o *GetSystemGcScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemGcScheduleForbidden creates a GetSystemGcScheduleForbidden with default headers values
func NewGetSystemGcScheduleForbidden() *GetSystemGcScheduleForbidden {
	return &GetSystemGcScheduleForbidden{}
}

/*GetSystemGcScheduleForbidden handles this case with default header values.

Only admin has this authority.
*/
type GetSystemGcScheduleForbidden struct {
}

func (o *GetSystemGcScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getSystemGcScheduleForbidden ", 403)
}

func (o *GetSystemGcScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemGcScheduleInternalServerError creates a GetSystemGcScheduleInternalServerError with default headers values
func NewGetSystemGcScheduleInternalServerError() *GetSystemGcScheduleInternalServerError {
	return &GetSystemGcScheduleInternalServerError{}
}

/*GetSystemGcScheduleInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetSystemGcScheduleInternalServerError struct {
}

func (o *GetSystemGcScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getSystemGcScheduleInternalServerError ", 500)
}

func (o *GetSystemGcScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
