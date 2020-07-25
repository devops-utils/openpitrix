// Code generated by go-swagger; DO NOT EDIT.

package release_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new release manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for release manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRelease creates release
*/
func (a *Client) CreateRelease(params *CreateReleaseParams, authInfo runtime.ClientAuthInfoWriter) (*CreateReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRelease",
		Method:             "PUT",
		PathPattern:        "/v1/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateReleaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateReleaseOK), nil

}

/*
DeleteRelease deletes release
*/
func (a *Client) DeleteRelease(params *DeleteReleaseParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRelease",
		Method:             "DELETE",
		PathPattern:        "/v1/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReleaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteReleaseOK), nil

}

/*
ListReleases lists release
*/
func (a *Client) ListReleases(params *ListReleasesParams, authInfo runtime.ClientAuthInfoWriter) (*ListReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListReleases",
		Method:             "GET",
		PathPattern:        "/v1/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReleasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReleasesOK), nil

}

/*
RollbackRelease rollbacks release
*/
func (a *Client) RollbackRelease(params *RollbackReleaseParams, authInfo runtime.ClientAuthInfoWriter) (*RollbackReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRollbackReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RollbackRelease",
		Method:             "POST",
		PathPattern:        "/v1/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RollbackReleaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RollbackReleaseOK), nil

}

/*
UpgradeRelease upgrades release
*/
func (a *Client) UpgradeRelease(params *UpgradeReleaseParams, authInfo runtime.ClientAuthInfoWriter) (*UpgradeReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpgradeRelease",
		Method:             "PATCH",
		PathPattern:        "/v1/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeReleaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeReleaseOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}