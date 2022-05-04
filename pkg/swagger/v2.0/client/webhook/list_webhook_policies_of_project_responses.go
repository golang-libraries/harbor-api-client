// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

// ListWebhookPoliciesOfProjectReader is a Reader for the ListWebhookPoliciesOfProject structure.
type ListWebhookPoliciesOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWebhookPoliciesOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWebhookPoliciesOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListWebhookPoliciesOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListWebhookPoliciesOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListWebhookPoliciesOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListWebhookPoliciesOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListWebhookPoliciesOfProjectOK creates a ListWebhookPoliciesOfProjectOK with default headers values
func NewListWebhookPoliciesOfProjectOK() *ListWebhookPoliciesOfProjectOK {
	return &ListWebhookPoliciesOfProjectOK{}
}

/* ListWebhookPoliciesOfProjectOK describes a response with status code 200, with default header values.

Success
*/
type ListWebhookPoliciesOfProjectOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of webhook policies.
	 */
	XTotalCount int64

	Payload []*models.WebhookPolicy
}

func (o *ListWebhookPoliciesOfProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies][%d] listWebhookPoliciesOfProjectOK  %+v", 200, o.Payload)
}
func (o *ListWebhookPoliciesOfProjectOK) GetPayload() []*models.WebhookPolicy {
	return o.Payload
}

func (o *ListWebhookPoliciesOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListWebhookPoliciesOfProjectBadRequest creates a ListWebhookPoliciesOfProjectBadRequest with default headers values
func NewListWebhookPoliciesOfProjectBadRequest() *ListWebhookPoliciesOfProjectBadRequest {
	return &ListWebhookPoliciesOfProjectBadRequest{}
}

/* ListWebhookPoliciesOfProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListWebhookPoliciesOfProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListWebhookPoliciesOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies][%d] listWebhookPoliciesOfProjectBadRequest  %+v", 400, o.Payload)
}
func (o *ListWebhookPoliciesOfProjectBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookPoliciesOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookPoliciesOfProjectUnauthorized creates a ListWebhookPoliciesOfProjectUnauthorized with default headers values
func NewListWebhookPoliciesOfProjectUnauthorized() *ListWebhookPoliciesOfProjectUnauthorized {
	return &ListWebhookPoliciesOfProjectUnauthorized{}
}

/* ListWebhookPoliciesOfProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListWebhookPoliciesOfProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListWebhookPoliciesOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies][%d] listWebhookPoliciesOfProjectUnauthorized  %+v", 401, o.Payload)
}
func (o *ListWebhookPoliciesOfProjectUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookPoliciesOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookPoliciesOfProjectForbidden creates a ListWebhookPoliciesOfProjectForbidden with default headers values
func NewListWebhookPoliciesOfProjectForbidden() *ListWebhookPoliciesOfProjectForbidden {
	return &ListWebhookPoliciesOfProjectForbidden{}
}

/* ListWebhookPoliciesOfProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListWebhookPoliciesOfProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListWebhookPoliciesOfProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies][%d] listWebhookPoliciesOfProjectForbidden  %+v", 403, o.Payload)
}
func (o *ListWebhookPoliciesOfProjectForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookPoliciesOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListWebhookPoliciesOfProjectInternalServerError creates a ListWebhookPoliciesOfProjectInternalServerError with default headers values
func NewListWebhookPoliciesOfProjectInternalServerError() *ListWebhookPoliciesOfProjectInternalServerError {
	return &ListWebhookPoliciesOfProjectInternalServerError{}
}

/* ListWebhookPoliciesOfProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListWebhookPoliciesOfProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListWebhookPoliciesOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies][%d] listWebhookPoliciesOfProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *ListWebhookPoliciesOfProjectInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListWebhookPoliciesOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
