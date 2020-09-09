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

	"github.com/mittwald/goharbor-client/apiv2/model/legacy"
)

// NewPostUsergroupsParams creates a new PostUsergroupsParams object
// with the default values initialized.
func NewPostUsergroupsParams() *PostUsergroupsParams {
	var ()
	return &PostUsergroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsergroupsParamsWithTimeout creates a new PostUsergroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsergroupsParamsWithTimeout(timeout time.Duration) *PostUsergroupsParams {
	var ()
	return &PostUsergroupsParams{

		timeout: timeout,
	}
}

// NewPostUsergroupsParamsWithContext creates a new PostUsergroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsergroupsParamsWithContext(ctx context.Context) *PostUsergroupsParams {
	var ()
	return &PostUsergroupsParams{

		Context: ctx,
	}
}

// NewPostUsergroupsParamsWithHTTPClient creates a new PostUsergroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsergroupsParamsWithHTTPClient(client *http.Client) *PostUsergroupsParams {
	var ()
	return &PostUsergroupsParams{
		HTTPClient: client,
	}
}

/*PostUsergroupsParams contains all the parameters to send to the API endpoint
for the post usergroups operation typically these are written to a http.Request
*/
type PostUsergroupsParams struct {

	/*Usergroup*/
	Usergroup *legacy.UserGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post usergroups params
func (o *PostUsergroupsParams) WithTimeout(timeout time.Duration) *PostUsergroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post usergroups params
func (o *PostUsergroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post usergroups params
func (o *PostUsergroupsParams) WithContext(ctx context.Context) *PostUsergroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post usergroups params
func (o *PostUsergroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post usergroups params
func (o *PostUsergroupsParams) WithHTTPClient(client *http.Client) *PostUsergroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post usergroups params
func (o *PostUsergroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsergroup adds the usergroup to the post usergroups params
func (o *PostUsergroupsParams) WithUsergroup(usergroup *legacy.UserGroup) *PostUsergroupsParams {
	o.SetUsergroup(usergroup)
	return o
}

// SetUsergroup adds the usergroup to the post usergroups params
func (o *PostUsergroupsParams) SetUsergroup(usergroup *legacy.UserGroup) {
	o.Usergroup = usergroup
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsergroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Usergroup != nil {
		if err := r.SetBodyParam(o.Usergroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}