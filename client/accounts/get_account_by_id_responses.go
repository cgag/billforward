package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// GetAccountByIDReader is a Reader for the GetAccountByID structure.
type GetAccountByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAccountByIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAccountByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetAccountByIDOK creates a GetAccountByIDOK with default headers values
func NewGetAccountByIDOK() *GetAccountByIDOK {
	return &GetAccountByIDOK{}
}

/*GetAccountByIDOK handles this case with default header values.

success
*/
type GetAccountByIDOK struct {
	Payload *models.AccountPagedMetadata
}

func (o *GetAccountByIDOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-ID}][%d] getAccountByIdOK  %+v", 200, o.Payload)
}

func (o *GetAccountByIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountByIDDefault creates a GetAccountByIDDefault with default headers values
func NewGetAccountByIDDefault(code int) *GetAccountByIDDefault {
	return &GetAccountByIDDefault{
		_statusCode: code,
	}
}

/*GetAccountByIDDefault handles this case with default header values.

error
*/
type GetAccountByIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get account by ID default response
func (o *GetAccountByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountByIDDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account-ID}][%d] getAccountByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountByIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
