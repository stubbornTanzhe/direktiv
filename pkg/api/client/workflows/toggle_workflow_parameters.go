// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// NewToggleWorkflowParams creates a new ToggleWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewToggleWorkflowParams() *ToggleWorkflowParams {
	return &ToggleWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewToggleWorkflowParamsWithTimeout creates a new ToggleWorkflowParams object
// with the ability to set a timeout on a request.
func NewToggleWorkflowParamsWithTimeout(timeout time.Duration) *ToggleWorkflowParams {
	return &ToggleWorkflowParams{
		timeout: timeout,
	}
}

// NewToggleWorkflowParamsWithContext creates a new ToggleWorkflowParams object
// with the ability to set a context for a request.
func NewToggleWorkflowParamsWithContext(ctx context.Context) *ToggleWorkflowParams {
	return &ToggleWorkflowParams{
		Context: ctx,
	}
}

// NewToggleWorkflowParamsWithHTTPClient creates a new ToggleWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewToggleWorkflowParamsWithHTTPClient(client *http.Client) *ToggleWorkflowParams {
	return &ToggleWorkflowParams{
		HTTPClient: client,
	}
}

/* ToggleWorkflowParams contains all the parameters to send to the API endpoint
   for the toggle workflow operation.

   Typically these are written to a http.Request.
*/
type ToggleWorkflowParams struct {

	/* WorkflowLiveStatus.

	   Whether or not the workflow is alive or disabled
	*/
	WorkflowLiveStatus ToggleWorkflowBody

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

// WithDefaults hydrates default values in the toggle workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleWorkflowParams) WithDefaults() *ToggleWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the toggle workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleWorkflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the toggle workflow params
func (o *ToggleWorkflowParams) WithTimeout(timeout time.Duration) *ToggleWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the toggle workflow params
func (o *ToggleWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the toggle workflow params
func (o *ToggleWorkflowParams) WithContext(ctx context.Context) *ToggleWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the toggle workflow params
func (o *ToggleWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the toggle workflow params
func (o *ToggleWorkflowParams) WithHTTPClient(client *http.Client) *ToggleWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the toggle workflow params
func (o *ToggleWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowLiveStatus adds the workflowLiveStatus to the toggle workflow params
func (o *ToggleWorkflowParams) WithWorkflowLiveStatus(workflowLiveStatus ToggleWorkflowBody) *ToggleWorkflowParams {
	o.SetWorkflowLiveStatus(workflowLiveStatus)
	return o
}

// SetWorkflowLiveStatus adds the workflowLiveStatus to the toggle workflow params
func (o *ToggleWorkflowParams) SetWorkflowLiveStatus(workflowLiveStatus ToggleWorkflowBody) {
	o.WorkflowLiveStatus = workflowLiveStatus
}

// WithNamespace adds the namespace to the toggle workflow params
func (o *ToggleWorkflowParams) WithNamespace(namespace string) *ToggleWorkflowParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the toggle workflow params
func (o *ToggleWorkflowParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithWorkflow adds the workflow to the toggle workflow params
func (o *ToggleWorkflowParams) WithWorkflow(workflow string) *ToggleWorkflowParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the toggle workflow params
func (o *ToggleWorkflowParams) SetWorkflow(workflow string) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *ToggleWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.WorkflowLiveStatus); err != nil {
		return err
	}

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