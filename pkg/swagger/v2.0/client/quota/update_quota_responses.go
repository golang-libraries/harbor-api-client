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

// UpdateQuotaReader is a Reader for the UpdateQuota structure.
type UpdateQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateQuotaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateQuotaOK creates a UpdateQuotaOK with default headers values
func NewUpdateQuotaOK() *UpdateQuotaOK {
	return &UpdateQuotaOK{}
}

/* UpdateQuotaOK describes a response with status code 200, with default header values.

Success
*/
type UpdateQuotaOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateQuotaOK) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] updateQuotaOK ", 200)
}

func (o *UpdateQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateQuotaBadRequest creates a UpdateQuotaBadRequest with default headers values
func NewUpdateQuotaBadRequest() *UpdateQuotaBadRequest {
	return &UpdateQuotaBadRequest{}
}

/* UpdateQuotaBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateQuotaBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateQuotaBadRequest) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] updateQuotaBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateQuotaBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateQuotaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaUnauthorized creates a UpdateQuotaUnauthorized with default headers values
func NewUpdateQuotaUnauthorized() *UpdateQuotaUnauthorized {
	return &UpdateQuotaUnauthorized{}
}

/* UpdateQuotaUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateQuotaUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] updateQuotaUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateQuotaUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaForbidden creates a UpdateQuotaForbidden with default headers values
func NewUpdateQuotaForbidden() *UpdateQuotaForbidden {
	return &UpdateQuotaForbidden{}
}

/* UpdateQuotaForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateQuotaForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateQuotaForbidden) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] updateQuotaForbidden  %+v", 403, o.Payload)
}
func (o *UpdateQuotaForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaNotFound creates a UpdateQuotaNotFound with default headers values
func NewUpdateQuotaNotFound() *UpdateQuotaNotFound {
	return &UpdateQuotaNotFound{}
}

/* UpdateQuotaNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateQuotaNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateQuotaNotFound) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] updateQuotaNotFound  %+v", 404, o.Payload)
}
func (o *UpdateQuotaNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaInternalServerError creates a UpdateQuotaInternalServerError with default headers values
func NewUpdateQuotaInternalServerError() *UpdateQuotaInternalServerError {
	return &UpdateQuotaInternalServerError{}
}

/* UpdateQuotaInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateQuotaInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] updateQuotaInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateQuotaInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
