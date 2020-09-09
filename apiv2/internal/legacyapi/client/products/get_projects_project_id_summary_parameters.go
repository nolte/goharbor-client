// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewGetProjectsProjectIDSummaryParams creates a new GetProjectsProjectIDSummaryParams object
// with the default values initialized.
func NewGetProjectsProjectIDSummaryParams() *GetProjectsProjectIDSummaryParams {
	var ()
	return &GetProjectsProjectIDSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsProjectIDSummaryParamsWithTimeout creates a new GetProjectsProjectIDSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsProjectIDSummaryParamsWithTimeout(timeout time.Duration) *GetProjectsProjectIDSummaryParams {
	var ()
	return &GetProjectsProjectIDSummaryParams{

		timeout: timeout,
	}
}

// NewGetProjectsProjectIDSummaryParamsWithContext creates a new GetProjectsProjectIDSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsProjectIDSummaryParamsWithContext(ctx context.Context) *GetProjectsProjectIDSummaryParams {
	var ()
	return &GetProjectsProjectIDSummaryParams{

		Context: ctx,
	}
}

// NewGetProjectsProjectIDSummaryParamsWithHTTPClient creates a new GetProjectsProjectIDSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectsProjectIDSummaryParamsWithHTTPClient(client *http.Client) *GetProjectsProjectIDSummaryParams {
	var ()
	return &GetProjectsProjectIDSummaryParams{
		HTTPClient: client,
	}
}

/*GetProjectsProjectIDSummaryParams contains all the parameters to send to the API endpoint
for the get projects project ID summary operation typically these are written to a http.Request
*/
type GetProjectsProjectIDSummaryParams struct {

	/*ProjectID
	  Relevant project ID

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) WithTimeout(timeout time.Duration) *GetProjectsProjectIDSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) WithContext(ctx context.Context) *GetProjectsProjectIDSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) WithHTTPClient(client *http.Client) *GetProjectsProjectIDSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) WithProjectID(projectID int64) *GetProjectsProjectIDSummaryParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get projects project ID summary params
func (o *GetProjectsProjectIDSummaryParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsProjectIDSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}