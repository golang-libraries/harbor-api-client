// Code generated by go-swagger; DO NOT EDIT.

package label

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
	"github.com/go-openapi/swag"

	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

// NewUpdateLabelParams creates a new UpdateLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLabelParams() *UpdateLabelParams {
	return &UpdateLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLabelParamsWithTimeout creates a new UpdateLabelParams object
// with the ability to set a timeout on a request.
func NewUpdateLabelParamsWithTimeout(timeout time.Duration) *UpdateLabelParams {
	return &UpdateLabelParams{
		timeout: timeout,
	}
}

// NewUpdateLabelParamsWithContext creates a new UpdateLabelParams object
// with the ability to set a context for a request.
func NewUpdateLabelParamsWithContext(ctx context.Context) *UpdateLabelParams {
	return &UpdateLabelParams{
		Context: ctx,
	}
}

// NewUpdateLabelParamsWithHTTPClient creates a new UpdateLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLabelParamsWithHTTPClient(client *http.Client) *UpdateLabelParams {
	return &UpdateLabelParams{
		HTTPClient: client,
	}
}

/* UpdateLabelParams contains all the parameters to send to the API endpoint
   for the update label operation.

   Typically these are written to a http.Request.
*/
type UpdateLabelParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Label.

	   The updated label json object.
	*/
	Label *models.Label

	/* LabelID.

	   Label ID

	   Format: int64
	*/
	LabelID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLabelParams) WithDefaults() *UpdateLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update label params
func (o *UpdateLabelParams) WithTimeout(timeout time.Duration) *UpdateLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update label params
func (o *UpdateLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update label params
func (o *UpdateLabelParams) WithContext(ctx context.Context) *UpdateLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update label params
func (o *UpdateLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update label params
func (o *UpdateLabelParams) WithHTTPClient(client *http.Client) *UpdateLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update label params
func (o *UpdateLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update label params
func (o *UpdateLabelParams) WithXRequestID(xRequestID *string) *UpdateLabelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update label params
func (o *UpdateLabelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithLabel adds the label to the update label params
func (o *UpdateLabelParams) WithLabel(label *models.Label) *UpdateLabelParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the update label params
func (o *UpdateLabelParams) SetLabel(label *models.Label) {
	o.Label = label
}

// WithLabelID adds the labelID to the update label params
func (o *UpdateLabelParams) WithLabelID(labelID int64) *UpdateLabelParams {
	o.SetLabelID(labelID)
	return o
}

// SetLabelID adds the labelId to the update label params
func (o *UpdateLabelParams) SetLabelID(labelID int64) {
	o.LabelID = labelID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	// path param label_id
	if err := r.SetPathParam("label_id", swag.FormatInt64(o.LabelID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
