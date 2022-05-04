// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

// PingLdapReader is a Reader for the PingLdap structure.
type PingLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingLdapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPingLdapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPingLdapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPingLdapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPingLdapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPingLdapOK creates a PingLdapOK with default headers values
func NewPingLdapOK() *PingLdapOK {
	return &PingLdapOK{}
}

/* PingLdapOK describes a response with status code 200, with default header values.

Ping ldap service successfully.
*/
type PingLdapOK struct {
	Payload *models.LdapPingResult
}

func (o *PingLdapOK) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] pingLdapOK  %+v", 200, o.Payload)
}
func (o *PingLdapOK) GetPayload() *models.LdapPingResult {
	return o.Payload
}

func (o *PingLdapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LdapPingResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingLdapBadRequest creates a PingLdapBadRequest with default headers values
func NewPingLdapBadRequest() *PingLdapBadRequest {
	return &PingLdapBadRequest{}
}

/* PingLdapBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PingLdapBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *PingLdapBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] pingLdapBadRequest  %+v", 400, o.Payload)
}
func (o *PingLdapBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingLdapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingLdapUnauthorized creates a PingLdapUnauthorized with default headers values
func NewPingLdapUnauthorized() *PingLdapUnauthorized {
	return &PingLdapUnauthorized{}
}

/* PingLdapUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PingLdapUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *PingLdapUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] pingLdapUnauthorized  %+v", 401, o.Payload)
}
func (o *PingLdapUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingLdapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingLdapForbidden creates a PingLdapForbidden with default headers values
func NewPingLdapForbidden() *PingLdapForbidden {
	return &PingLdapForbidden{}
}

/* PingLdapForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PingLdapForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *PingLdapForbidden) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] pingLdapForbidden  %+v", 403, o.Payload)
}
func (o *PingLdapForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingLdapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingLdapInternalServerError creates a PingLdapInternalServerError with default headers values
func NewPingLdapInternalServerError() *PingLdapInternalServerError {
	return &PingLdapInternalServerError{}
}

/* PingLdapInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PingLdapInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *PingLdapInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] pingLdapInternalServerError  %+v", 500, o.Payload)
}
func (o *PingLdapInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingLdapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
