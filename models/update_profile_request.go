package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
UpdateProfileRequest

swagger:model UpdateProfileRequest
*/
type UpdateProfileRequest struct {

	/* { "description" : "", "verbs":["GET"] }

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* { "description" : "Any additional information", "verbs":["POST","PUT","GET"] }
	 */
	AdditionalInformation *string `json:"additionalInformation,omitempty"`

	/* { "description" : "Address associated with the profile", "verbs":["POST","PUT","GET"] }
	 */
	Addresses []*Address `json:"addresses,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	CompanyName *string `json:"companyName,omitempty"`

	/* { "description" : "Date of birth in YYYY-MM-DD format", "verbs":["POST","PUT","GET"] }
	 */
	Dob *strfmt.DateTime `json:"dob,omitempty"`

	/* { "description" : "E-mail address", "verbs":["POST","PUT","GET"] }
	 */
	Email *string `json:"email,omitempty"`

	/* { "description" : "Fax number", "verbs":["POST","PUT","GET"] }
	 */
	Fax *string `json:"fax,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	FirstName *string `json:"firstName,omitempty"`

	/* { "description" : "Home telephone number", "verbs":["POST","PUT","GET"] }
	 */
	Landline *string `json:"landline,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	LastName *string `json:"lastName,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	LogoURL *string `json:"logoURL,omitempty"`

	/* { "description" : "Mobile telephone number", "verbs":["POST","PUT","GET"] }
	 */
	Mobile *string `json:"mobile,omitempty"`

	/* { "description" : "", "verbs":[] }
	 */
	OrganizationID *string `json:"organizationID,omitempty"`

	/* { "description" : "VAT number", "verbs":["POST","PUT","GET"] }
	 */
	VatNumber *string `json:"vatNumber,omitempty"`
}

// Validate validates this update profile request
func (m *UpdateProfileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProfileRequest) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}