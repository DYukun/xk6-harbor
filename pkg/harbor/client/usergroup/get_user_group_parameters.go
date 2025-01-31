// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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

// NewGetUserGroupParams creates a new GetUserGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserGroupParams() *GetUserGroupParams {
	return &GetUserGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserGroupParamsWithTimeout creates a new GetUserGroupParams object
// with the ability to set a timeout on a request.
func NewGetUserGroupParamsWithTimeout(timeout time.Duration) *GetUserGroupParams {
	return &GetUserGroupParams{
		timeout: timeout,
	}
}

// NewGetUserGroupParamsWithContext creates a new GetUserGroupParams object
// with the ability to set a context for a request.
func NewGetUserGroupParamsWithContext(ctx context.Context) *GetUserGroupParams {
	return &GetUserGroupParams{
		Context: ctx,
	}
}

// NewGetUserGroupParamsWithHTTPClient creates a new GetUserGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserGroupParamsWithHTTPClient(client *http.Client) *GetUserGroupParams {
	return &GetUserGroupParams{
		HTTPClient: client,
	}
}

/* GetUserGroupParams contains all the parameters to send to the API endpoint
   for the get user group operation.

   Typically these are written to a http.Request.
*/
type GetUserGroupParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* GroupID.

	   Group ID

	   Format: int64
	*/
	GroupID int64 `js:"groupID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserGroupParams) WithDefaults() *GetUserGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user group params
func (o *GetUserGroupParams) WithTimeout(timeout time.Duration) *GetUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user group params
func (o *GetUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user group params
func (o *GetUserGroupParams) WithContext(ctx context.Context) *GetUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user group params
func (o *GetUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user group params
func (o *GetUserGroupParams) WithHTTPClient(client *http.Client) *GetUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user group params
func (o *GetUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get user group params
func (o *GetUserGroupParams) WithXRequestID(xRequestID *string) *GetUserGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get user group params
func (o *GetUserGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGroupID adds the groupID to the get user group params
func (o *GetUserGroupParams) WithGroupID(groupID int64) *GetUserGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get user group params
func (o *GetUserGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
