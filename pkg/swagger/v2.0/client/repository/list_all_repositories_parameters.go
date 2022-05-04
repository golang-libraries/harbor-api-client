// Code generated by go-swagger; DO NOT EDIT.

package repository

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
)

// NewListAllRepositoriesParams creates a new ListAllRepositoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAllRepositoriesParams() *ListAllRepositoriesParams {
	return &ListAllRepositoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAllRepositoriesParamsWithTimeout creates a new ListAllRepositoriesParams object
// with the ability to set a timeout on a request.
func NewListAllRepositoriesParamsWithTimeout(timeout time.Duration) *ListAllRepositoriesParams {
	return &ListAllRepositoriesParams{
		timeout: timeout,
	}
}

// NewListAllRepositoriesParamsWithContext creates a new ListAllRepositoriesParams object
// with the ability to set a context for a request.
func NewListAllRepositoriesParamsWithContext(ctx context.Context) *ListAllRepositoriesParams {
	return &ListAllRepositoriesParams{
		Context: ctx,
	}
}

// NewListAllRepositoriesParamsWithHTTPClient creates a new ListAllRepositoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAllRepositoriesParamsWithHTTPClient(client *http.Client) *ListAllRepositoriesParams {
	return &ListAllRepositoriesParams{
		HTTPClient: client,
	}
}

/* ListAllRepositoriesParams contains all the parameters to send to the API endpoint
   for the list all repositories operation.

   Typically these are written to a http.Request.
*/
type ListAllRepositoriesParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list all repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAllRepositoriesParams) WithDefaults() *ListAllRepositoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list all repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAllRepositoriesParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListAllRepositoriesParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list all repositories params
func (o *ListAllRepositoriesParams) WithTimeout(timeout time.Duration) *ListAllRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all repositories params
func (o *ListAllRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all repositories params
func (o *ListAllRepositoriesParams) WithContext(ctx context.Context) *ListAllRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all repositories params
func (o *ListAllRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all repositories params
func (o *ListAllRepositoriesParams) WithHTTPClient(client *http.Client) *ListAllRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all repositories params
func (o *ListAllRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list all repositories params
func (o *ListAllRepositoriesParams) WithXRequestID(xRequestID *string) *ListAllRepositoriesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list all repositories params
func (o *ListAllRepositoriesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list all repositories params
func (o *ListAllRepositoriesParams) WithPage(page *int64) *ListAllRepositoriesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list all repositories params
func (o *ListAllRepositoriesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list all repositories params
func (o *ListAllRepositoriesParams) WithPageSize(pageSize *int64) *ListAllRepositoriesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list all repositories params
func (o *ListAllRepositoriesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithQ adds the q to the list all repositories params
func (o *ListAllRepositoriesParams) WithQ(q *string) *ListAllRepositoriesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list all repositories params
func (o *ListAllRepositoriesParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list all repositories params
func (o *ListAllRepositoriesParams) WithSort(sort *string) *ListAllRepositoriesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list all repositories params
func (o *ListAllRepositoriesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
