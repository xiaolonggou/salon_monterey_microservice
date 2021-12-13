// Code generated by go-swagger; DO NOT EDIT.

package arts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xiaolonggou/microservice/v1/sdk/models"
)

// AddArtPieceReader is a Reader for the AddArtPiece structure.
type AddArtPieceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddArtPieceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddArtPieceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewAddArtPieceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewAddArtPieceNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddArtPieceOK creates a AddArtPieceOK with default headers values
func NewAddArtPieceOK() *AddArtPieceOK {
	return &AddArtPieceOK{}
}

/* AddArtPieceOK describes a response with status code 200, with default header values.

A list of art pieces
*/
type AddArtPieceOK struct {
	Payload []*models.ArtPiece
}

func (o *AddArtPieceOK) Error() string {
	return fmt.Sprintf("[POST /arts][%d] addArtPieceOK  %+v", 200, o.Payload)
}
func (o *AddArtPieceOK) GetPayload() []*models.ArtPiece {
	return o.Payload
}

func (o *AddArtPieceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddArtPieceUnprocessableEntity creates a AddArtPieceUnprocessableEntity with default headers values
func NewAddArtPieceUnprocessableEntity() *AddArtPieceUnprocessableEntity {
	return &AddArtPieceUnprocessableEntity{}
}

/* AddArtPieceUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors defined as an array of strings
*/
type AddArtPieceUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *AddArtPieceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /arts][%d] addArtPieceUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *AddArtPieceUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AddArtPieceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddArtPieceNotImplemented creates a AddArtPieceNotImplemented with default headers values
func NewAddArtPieceNotImplemented() *AddArtPieceNotImplemented {
	return &AddArtPieceNotImplemented{}
}

/* AddArtPieceNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type AddArtPieceNotImplemented struct {
	Payload *models.GenericError
}

func (o *AddArtPieceNotImplemented) Error() string {
	return fmt.Sprintf("[POST /arts][%d] addArtPieceNotImplemented  %+v", 501, o.Payload)
}
func (o *AddArtPieceNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *AddArtPieceNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}