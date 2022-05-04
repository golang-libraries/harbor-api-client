// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

// GetQuotaReader is a Reader for the GetQuota structure.
type GetQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQuotaOK creates a GetQuotaOK with default headers values
func NewGetQuotaOK() *GetQuotaOK {
	return &GetQuotaOK{}
}

/* GetQuotaOK describes a response with status code 200, with default header values.

Successfully retrieved the quota.
*/
type GetQuotaOK struct {
	Payload *models.Quota
}

func (o *GetQuotaOK) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotaOK  %+v", 200, o.Payload)
}
func (o *GetQuotaOK) GetPayload() *models.Quota {
	return o.Payload
}

func (o *GetQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Quota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaUnauthorized creates a GetQuotaUnauthorized with default headers values
func NewGetQuotaUnauthorized() *GetQuotaUnauthorized {
	return &GetQuotaUnauthorized{}
}

/* GetQuotaUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetQuotaUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotaUnauthorized  %+v", 401, o.Payload)
}
func (o *GetQuotaUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaForbidden creates a GetQuotaForbidden with default headers values
func NewGetQuotaForbidden() *GetQuotaForbidden {
	return &GetQuotaForbidden{}
}

/* GetQuotaForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetQuotaForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetQuotaForbidden) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotaForbidden  %+v", 403, o.Payload)
}
func (o *GetQuotaForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaNotFound creates a GetQuotaNotFound with default headers values
func NewGetQuotaNotFound() *GetQuotaNotFound {
	return &GetQuotaNotFound{}
}

/* GetQuotaNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetQuotaNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetQuotaNotFound) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotaNotFound  %+v", 404, o.Payload)
}
func (o *GetQuotaNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaInternalServerError creates a GetQuotaInternalServerError with default headers values
func NewGetQuotaInternalServerError() *GetQuotaInternalServerError {
	return &GetQuotaInternalServerError{}
}

/* GetQuotaInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetQuotaInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotaInternalServerError  %+v", 500, o.Payload)
}
func (o *GetQuotaInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
