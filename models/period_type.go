package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
PeriodType period type

swagger:model PeriodType
*/
type PeriodType struct {

	/* Name name
	 */
	Name string `json:"name,omitempty"`
}

// Validate validates this period type
func (m *PeriodType) Validate(formats strfmt.Registry) error {
	return nil
}