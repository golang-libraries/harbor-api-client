// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/* GetUserOK describes a response with status code 200, with default header values.

Get user's info successfully.
*/
type GetUserOK struct {
	Payload *models.UserResp
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserOK  %+v", 200, o.Payload)
}
func (o *GetUserOK) GetPayload() *models.UserResp {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUnauthorized creates a GetUserUnauthorized with default headers values
func NewGetUserUnauthorized() *GetUserUnauthorized {
	return &GetUserUnauthorized{}
}

/* GetUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetUserUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/* GetUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUserForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserForbidden  %+v", 403, o.Payload)
}
func (o *GetUserForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/* GetUserNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetUserNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserNotFound  %+v", 404, o.Payload)
}
func (o *GetUserNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUserInternalServerError creates a GetUserInternalServerError with default headers values
func NewGetUserInternalServerError() *GetUserInternalServerError {
	return &GetUserInternalServerError{}
}

/* GetUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUserInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
