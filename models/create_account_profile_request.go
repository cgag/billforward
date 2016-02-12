package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*CreateAccountProfileRequest CreateAccountProfileRequest create account profile request

swagger:model CreateAccountProfileRequest
*/
type CreateAccountProfileRequest struct {

	/* { "description" : "", "verbs":["GET"] }

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* { "description" : "Any additional information", "verbs":["POST","PUT","GET"] }
	 */
	AdditionalInformation *string `json:"additionalInformation,omitempty"`

	/* { "description" : "Address associated with the profile", "verbs":["POST","PUT","GET"] }
	 */
	Addresses []*CreateProfileAddressRequest `json:"addresses,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	CompanyName *string `json:"companyName,omitempty"`

	/* { "description" : "Date of birth in YYYY-MM-DD format", "verbs":["POST","PUT","GET"] }
	 */
	Dob strfmt.DateTime `json:"dob,omitempty"`

	/* { "description" : "E-mail address", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Email string `json:"email,omitempty"`

	/* { "description" : "Fax number", "verbs":["POST","PUT","GET"] }
	 */
	Fax *string `json:"fax,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	FirstName string `json:"firstName,omitempty"`

	/* { "description" : "Home telephone number", "verbs":["POST","PUT","GET"] }
	 */
	Landline *string `json:"landline,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	LastName string `json:"lastName,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	LogoURL *string `json:"logoURL,omitempty"`

	/* { "description" : "Mobile telephone number", "verbs":["POST","PUT","GET"] }
	 */
	Mobile *string `json:"mobile,omitempty"`

	/* { "description" : "VAT number", "verbs":["POST","PUT","GET"] }
	 */
	VatNumber *string `json:"vatNumber,omitempty"`
}

// Validate validates this create account profile request
func (m *CreateAccountProfileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAccountProfileRequest) validateAccountID(formats strfmt.Registry) error {

	if err := validate.RequiredString("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

func (m *CreateAccountProfileRequest) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {

			if err := m.Addresses[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CreateAccountProfileRequest) validateEmail(formats strfmt.Registry) error {

	if err := validate.RequiredString("email", "body", string(m.Email)); err != nil {
		return err
	}

	return nil
}

func (m *CreateAccountProfileRequest) validateFirstName(formats strfmt.Registry) error {

	if err := validate.RequiredString("firstName", "body", string(m.FirstName)); err != nil {
		return err
	}

	return nil
}

func (m *CreateAccountProfileRequest) validateLastName(formats strfmt.Registry) error {

	if err := validate.RequiredString("lastName", "body", string(m.LastName)); err != nil {
		return err
	}

	return nil
}
