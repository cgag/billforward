// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PendingComponentValueChange pending component value change
// swagger:model PendingComponentValueChange
type PendingComponentValueChange struct {

	// at
	At strfmt.DateTime `json:"at,omitempty"`

	// discard Http verb
	DiscardHTTPVerb string `json:"discardHttpVerb,omitempty"`

	// discard Url
	DiscardURL string `json:"discardUrl,omitempty"`

	// value
	Value int32 `json:"value,omitempty"`
}

// Validate validates this pending component value change
func (m *PendingComponentValueChange) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PendingComponentValueChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PendingComponentValueChange) UnmarshalBinary(b []byte) error {
	var res PendingComponentValueChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
