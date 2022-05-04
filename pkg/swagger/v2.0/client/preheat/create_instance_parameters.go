// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

// NewCreateInstanceParams creates a new CreateInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInstanceParams() *CreateInstanceParams {
	return &CreateInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstanceParamsWithTimeout creates a new CreateInstanceParams object
// with the ability to set a timeout on a request.
func NewCreateInstanceParamsWithTimeout(timeout time.Duration) *CreateInstanceParams {
	return &CreateInstanceParams{
		timeout: timeout,
	}
}

// NewCreateInstanceParamsWithContext creates a new CreateInstanceParams object
// with the ability to set a context for a request.
func NewCreateInstanceParamsWithContext(ctx context.Context) *CreateInstanceParams {
	return &CreateInstanceParams{
		Context: ctx,
	}
}

// NewCreateInstanceParamsWithHTTPClient creates a new CreateInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInstanceParamsWithHTTPClient(client *http.Client) *CreateInstanceParams {
	return &CreateInstanceParams{
		HTTPClient: client,
	}
}

/* CreateInstanceParams contains all the parameters to send to the API endpoint
   for the create instance operation.

   Typically these are written to a http.Request.
*/
type CreateInstanceParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Instance.

	   The JSON object of instance.
	*/
	Instance *models.Instance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstanceParams) WithDefaults() *CreateInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create instance params
func (o *CreateInstanceParams) WithTimeout(timeout time.Duration) *CreateInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create instance params
func (o *CreateInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create instance params
func (o *CreateInstanceParams) WithContext(ctx context.Context) *CreateInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create instance params
func (o *CreateInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create instance params
func (o *CreateInstanceParams) WithHTTPClient(client *http.Client) *CreateInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create instance params
func (o *CreateInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create instance params
func (o *CreateInstanceParams) WithXRequestID(xRequestID *string) *CreateInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create instance params
func (o *CreateInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithInstance adds the instance to the create instance params
func (o *CreateInstanceParams) WithInstance(instance *models.Instance) *CreateInstanceParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the create instance params
func (o *CreateInstanceParams) SetInstance(instance *models.Instance) {
	o.Instance = instance
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Instance != nil {
		if err := r.SetBodyParam(o.Instance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
