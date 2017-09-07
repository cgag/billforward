// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Alias The alias used to encrypt card details
// swagger:model Alias
type Alias struct {

	// { "description" : "A string representation of the alias.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Alias *string `json:"alias"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "Has this alias been deleted?", "verbs":["POST","PUT","GET"] }
	// Required: true
	Deleted bool `json:"deleted"`

	// { "description" : "ID of the alias.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ID *string `json:"id"`

	// { "description" : "ID of the organization associated with the alias.", "verbs":["POST","PUT","GET"] }
	// Required: true
	OrganizationID *string `json:"organizationID"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this alias
func (m *Alias) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlias(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alias) validateAlias(formats strfmt.Registry) error {

	if err := validate.Required("alias", "body", m.Alias); err != nil {
		return err
	}

	return nil
}

func (m *Alias) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

func (m *Alias) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Alias) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alias) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alias) UnmarshalBinary(b []byte) error {
	var res Alias
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
