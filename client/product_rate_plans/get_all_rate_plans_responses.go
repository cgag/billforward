package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// GetAllRatePlansReader is a Reader for the GetAllRatePlans structure.
type GetAllRatePlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAllRatePlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllRatePlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAllRatePlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetAllRatePlansOK creates a GetAllRatePlansOK with default headers values
func NewGetAllRatePlansOK() *GetAllRatePlansOK {
	return &GetAllRatePlansOK{}
}

/*GetAllRatePlansOK handles this case with default header values.

success
*/
type GetAllRatePlansOK struct {
	Payload *models.ProductRatePlanPagedMetadata
}

func (o *GetAllRatePlansOK) Error() string {
	return fmt.Sprintf("[GET /product-rate-plans][%d] getAllRatePlansOK  %+v", 200, o.Payload)
}

func (o *GetAllRatePlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductRatePlanPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllRatePlansDefault creates a GetAllRatePlansDefault with default headers values
func NewGetAllRatePlansDefault(code int) *GetAllRatePlansDefault {
	return &GetAllRatePlansDefault{
		_statusCode: code,
	}
}

/*GetAllRatePlansDefault handles this case with default header values.

error
*/
type GetAllRatePlansDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get all rate plans default response
func (o *GetAllRatePlansDefault) Code() int {
	return o._statusCode
}

func (o *GetAllRatePlansDefault) Error() string {
	return fmt.Sprintf("[GET /product-rate-plans][%d] getAllRatePlans default  %+v", o._statusCode, o.Payload)
}

func (o *GetAllRatePlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
