// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewGetWorkflowVariablesParams creates a new GetWorkflowVariablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowVariablesParams() *GetWorkflowVariablesParams {
	return &GetWorkflowVariablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowVariablesParamsWithTimeout creates a new GetWorkflowVariablesParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowVariablesParamsWithTimeout(timeout time.Duration) *GetWorkflowVariablesParams {
	return &GetWorkflowVariablesParams{
		timeout: timeout,
	}
}

// NewGetWorkflowVariablesParamsWithContext creates a new GetWorkflowVariablesParams object
// with the ability to set a context for a request.
func NewGetWorkflowVariablesParamsWithContext(ctx context.Context) *GetWorkflowVariablesParams {
	return &GetWorkflowVariablesParams{
		Context: ctx,
	}
}

// NewGetWorkflowVariablesParamsWithHTTPClient creates a new GetWorkflowVariablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowVariablesParamsWithHTTPClient(client *http.Client) *GetWorkflowVariablesParams {
	return &GetWorkflowVariablesParams{
		HTTPClient: client,
	}
}

/* GetWorkflowVariablesParams contains all the parameters to send to the API endpoint
   for the get workflow variables operation.

   Typically these are written to a http.Request.
*/
type GetWorkflowVariablesParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Workflow.

	   path to target workflow
	*/
	Workflow string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowVariablesParams) WithDefaults() *GetWorkflowVariablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowVariablesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow variables params
func (o *GetWorkflowVariablesParams) WithTimeout(timeout time.Duration) *GetWorkflowVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow variables params
func (o *GetWorkflowVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow variables params
func (o *GetWorkflowVariablesParams) WithContext(ctx context.Context) *GetWorkflowVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow variables params
func (o *GetWorkflowVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow variables params
func (o *GetWorkflowVariablesParams) WithHTTPClient(client *http.Client) *GetWorkflowVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow variables params
func (o *GetWorkflowVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get workflow variables params
func (o *GetWorkflowVariablesParams) WithNamespace(namespace string) *GetWorkflowVariablesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get workflow variables params
func (o *GetWorkflowVariablesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithWorkflow adds the workflow to the get workflow variables params
func (o *GetWorkflowVariablesParams) WithWorkflow(workflow string) *GetWorkflowVariablesParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the get workflow variables params
func (o *GetWorkflowVariablesParams) SetWorkflow(workflow string) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param workflow
	if err := r.SetPathParam("workflow", o.Workflow); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}