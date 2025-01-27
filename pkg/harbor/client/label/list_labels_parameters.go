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
)

// NewListLabelsParams creates a new ListLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLabelsParams() *ListLabelsParams {
	return &ListLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLabelsParamsWithTimeout creates a new ListLabelsParams object
// with the ability to set a timeout on a request.
func NewListLabelsParamsWithTimeout(timeout time.Duration) *ListLabelsParams {
	return &ListLabelsParams{
		timeout: timeout,
	}
}

// NewListLabelsParamsWithContext creates a new ListLabelsParams object
// with the ability to set a context for a request.
func NewListLabelsParamsWithContext(ctx context.Context) *ListLabelsParams {
	return &ListLabelsParams{
		Context: ctx,
	}
}

// NewListLabelsParamsWithHTTPClient creates a new ListLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLabelsParamsWithHTTPClient(client *http.Client) *ListLabelsParams {
	return &ListLabelsParams{
		HTTPClient: client,
	}
}

/* ListLabelsParams contains all the parameters to send to the API endpoint
   for the list labels operation.

   Typically these are written to a http.Request.
*/
type ListLabelsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Name.

	   The label name.
	*/
	Name *string `js:"name"`

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64 `js:"page"`

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64 `js:"pageSize"`

	/* ProjectID.

	   Relevant project ID, required when scope is p.

	   Format: int64
	*/
	ProjectID *int64 `js:"projectID"`

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string `js:"q"`

	/* Scope.

	   The label scope. Valid values are g and p. g for global labels and p for project labels.
	*/
	Scope *string `js:"scope"`

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string `js:"sort"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the list labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelsParams) WithDefaults() *ListLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLabelsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListLabelsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list labels params
func (o *ListLabelsParams) WithTimeout(timeout time.Duration) *ListLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list labels params
func (o *ListLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list labels params
func (o *ListLabelsParams) WithContext(ctx context.Context) *ListLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list labels params
func (o *ListLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list labels params
func (o *ListLabelsParams) WithHTTPClient(client *http.Client) *ListLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list labels params
func (o *ListLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list labels params
func (o *ListLabelsParams) WithXRequestID(xRequestID *string) *ListLabelsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list labels params
func (o *ListLabelsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the list labels params
func (o *ListLabelsParams) WithName(name *string) *ListLabelsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list labels params
func (o *ListLabelsParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the list labels params
func (o *ListLabelsParams) WithPage(page *int64) *ListLabelsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list labels params
func (o *ListLabelsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list labels params
func (o *ListLabelsParams) WithPageSize(pageSize *int64) *ListLabelsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list labels params
func (o *ListLabelsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the list labels params
func (o *ListLabelsParams) WithProjectID(projectID *int64) *ListLabelsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list labels params
func (o *ListLabelsParams) SetProjectID(projectID *int64) {
	o.ProjectID = projectID
}

// WithQ adds the q to the list labels params
func (o *ListLabelsParams) WithQ(q *string) *ListLabelsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list labels params
func (o *ListLabelsParams) SetQ(q *string) {
	o.Q = q
}

// WithScope adds the scope to the list labels params
func (o *ListLabelsParams) WithScope(scope *string) *ListLabelsParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the list labels params
func (o *ListLabelsParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSort adds the sort to the list labels params
func (o *ListLabelsParams) WithSort(sort *string) *ListLabelsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list labels params
func (o *ListLabelsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
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

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID int64

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt64(qrProjectID)
		if qProjectID != "" {

			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
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

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
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
