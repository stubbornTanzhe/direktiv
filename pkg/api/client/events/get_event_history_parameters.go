// Code generated by go-swagger; DO NOT EDIT.

package events

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
)

// NewGetEventHistoryParams creates a new GetEventHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEventHistoryParams() *GetEventHistoryParams {
	return &GetEventHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventHistoryParamsWithTimeout creates a new GetEventHistoryParams object
// with the ability to set a timeout on a request.
func NewGetEventHistoryParamsWithTimeout(timeout time.Duration) *GetEventHistoryParams {
	return &GetEventHistoryParams{
		timeout: timeout,
	}
}

// NewGetEventHistoryParamsWithContext creates a new GetEventHistoryParams object
// with the ability to set a context for a request.
func NewGetEventHistoryParamsWithContext(ctx context.Context) *GetEventHistoryParams {
	return &GetEventHistoryParams{
		Context: ctx,
	}
}

// NewGetEventHistoryParamsWithHTTPClient creates a new GetEventHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEventHistoryParamsWithHTTPClient(client *http.Client) *GetEventHistoryParams {
	return &GetEventHistoryParams{
		HTTPClient: client,
	}
}

/* GetEventHistoryParams contains all the parameters to send to the API endpoint
   for the get event history operation.

   Typically these are written to a http.Request.
*/
type GetEventHistoryParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get event history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventHistoryParams) WithDefaults() *GetEventHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get event history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get event history params
func (o *GetEventHistoryParams) WithTimeout(timeout time.Duration) *GetEventHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event history params
func (o *GetEventHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event history params
func (o *GetEventHistoryParams) WithContext(ctx context.Context) *GetEventHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event history params
func (o *GetEventHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event history params
func (o *GetEventHistoryParams) WithHTTPClient(client *http.Client) *GetEventHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event history params
func (o *GetEventHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get event history params
func (o *GetEventHistoryParams) WithNamespace(namespace string) *GetEventHistoryParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get event history params
func (o *GetEventHistoryParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}