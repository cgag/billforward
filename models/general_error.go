package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
GeneralError general error

swagger:model GeneralError
*/
type GeneralError struct {

	/* ErrorCode error code
	 */
	ErrorCode int64 `json:"errorCode,omitempty"`

	/* ErrorMessage error message
	 */
	ErrorMessage string `json:"errorMessage,omitempty"`

	/* ErrorType error type
	 */
	ErrorType string `json:"errorType,omitempty"`
}

// Validate validates this general error
func (m *GeneralError) Validate(formats strfmt.Registry) error {
	return nil
}
