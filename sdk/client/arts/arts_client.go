// Code generated by go-swagger; DO NOT EDIT.

package arts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new arts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for arts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddArtPiece(params *AddArtPieceParams, opts ...ClientOption) (*AddArtPieceOK, error)

	DeleteArtPiece(params *DeleteArtPieceParams, opts ...ClientOption) (*DeleteArtPieceCreated, error)

	UpdateArtPiece(params *UpdateArtPieceParams, opts ...ClientOption) (*UpdateArtPieceCreated, error)

	ListArtPieces(params *ListArtPiecesParams, opts ...ClientOption) (*ListArtPiecesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddArtPiece Create a new art piece
*/
func (a *Client) AddArtPiece(params *AddArtPieceParams, opts ...ClientOption) (*AddArtPieceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddArtPieceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddArtPiece",
		Method:             "POST",
		PathPattern:        "/arts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddArtPieceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddArtPieceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddArtPiece: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteArtPiece Delete an art piece
*/
func (a *Client) DeleteArtPiece(params *DeleteArtPieceParams, opts ...ClientOption) (*DeleteArtPieceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteArtPieceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteArtPiece",
		Method:             "DELETE",
		PathPattern:        "/art/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteArtPieceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteArtPieceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteArtPiece: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateArtPiece Update an art piece details
*/
func (a *Client) UpdateArtPiece(params *UpdateArtPieceParams, opts ...ClientOption) (*UpdateArtPieceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateArtPieceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateArtPiece",
		Method:             "PUT",
		PathPattern:        "/arts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateArtPieceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateArtPieceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateArtPiece: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListArtPieces Returns a list of art pieces
*/
func (a *Client) ListArtPieces(params *ListArtPiecesParams, opts ...ClientOption) (*ListArtPiecesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListArtPiecesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listArtPieces",
		Method:             "GET",
		PathPattern:        "/arts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListArtPiecesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListArtPiecesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listArtPieces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
