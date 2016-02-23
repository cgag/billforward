package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*StripeAuthCaptureRequest StripeAuthCaptureRequest stripe auth capture request

swagger:model StripeAuthCaptureRequest
*/
type StripeAuthCaptureRequest struct {
	accountIdField *string

	companyNameField *string

	defaultPaymentMethodField *bool

	emailField *string

	firstNameField *string

	gatewayField *string

	lastNameField *string

	mobileField *string

	organizationIdField *string

	/* {"description":"ID of the captured Card in Stripe. This can be provided as well as &mdash; or instead of &mdash; the one-use `stripeToken`, to lead BillForward to the card tokenized within the Stripe vault.","verbs":["POST"]}
	 */
	CardID *string `json:"cardID,omitempty"`

	/* CustomerID customer ID
	 */
	CustomerID *string `json:"customerID,omitempty"`

	/* {"description":"Single-use token <a href=\"https://stripe.com/docs/stripe.js\">provided by Stripe's client-side card capture SDK</a>, in response to your capturing a card into the Stripe vault. This token will be used by BillForward to find the tokenized card within the Stripe vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card.","verbs":["POST"]}

	Required: true
	*/
	StripeToken string `json:"stripeToken,omitempty"`
}

func (m *StripeAuthCaptureRequest) AtType() string {
	return "StripeAuthCaptureRequest"
}
func (m *StripeAuthCaptureRequest) SetAtType(val string) {

}

func (m *StripeAuthCaptureRequest) AccountID() *string {
	return m.accountIdField
}
func (m *StripeAuthCaptureRequest) SetAccountID(val *string) {
	m.accountIdField = val
}

func (m *StripeAuthCaptureRequest) CompanyName() *string {
	return m.companyNameField
}
func (m *StripeAuthCaptureRequest) SetCompanyName(val *string) {
	m.companyNameField = val
}

func (m *StripeAuthCaptureRequest) DefaultPaymentMethod() *bool {
	return m.defaultPaymentMethodField
}
func (m *StripeAuthCaptureRequest) SetDefaultPaymentMethod(val *bool) {
	m.defaultPaymentMethodField = val
}

func (m *StripeAuthCaptureRequest) Email() *string {
	return m.emailField
}
func (m *StripeAuthCaptureRequest) SetEmail(val *string) {
	m.emailField = val
}

func (m *StripeAuthCaptureRequest) FirstName() *string {
	return m.firstNameField
}
func (m *StripeAuthCaptureRequest) SetFirstName(val *string) {
	m.firstNameField = val
}

func (m *StripeAuthCaptureRequest) Gateway() *string {
	return m.gatewayField
}
func (m *StripeAuthCaptureRequest) SetGateway(val *string) {
	m.gatewayField = val
}

func (m *StripeAuthCaptureRequest) LastName() *string {
	return m.lastNameField
}
func (m *StripeAuthCaptureRequest) SetLastName(val *string) {
	m.lastNameField = val
}

func (m *StripeAuthCaptureRequest) Mobile() *string {
	return m.mobileField
}
func (m *StripeAuthCaptureRequest) SetMobile(val *string) {
	m.mobileField = val
}

func (m *StripeAuthCaptureRequest) OrganizationID() *string {
	return m.organizationIdField
}
func (m *StripeAuthCaptureRequest) SetOrganizationID(val *string) {
	m.organizationIdField = val
}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *StripeAuthCaptureRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		AtType string `json:"@type,omitempty"`

		AccountID *string `json:"accountID,omitempty"`

		CompanyName *string `json:"companyName,omitempty"`

		DefaultPaymentMethod *bool `json:"defaultPaymentMethod,omitempty"`

		Email *string `json:"email,omitempty"`

		FirstName *string `json:"firstName,omitempty"`

		Gateway *string `json:"gateway,omitempty"`

		LastName *string `json:"lastName,omitempty"`

		Mobile *string `json:"mobile,omitempty"`

		OrganizationID *string `json:"organizationID,omitempty"`

		/* {"description":"ID of the captured Card in Stripe. This can be provided as well as &mdash; or instead of &mdash; the one-use `stripeToken`, to lead BillForward to the card tokenized within the Stripe vault.","verbs":["POST"]}
		 */
		CardID *string `json:"cardID,omitempty"`

		/* CustomerID customer ID
		 */
		CustomerID *string `json:"customerID,omitempty"`

		/* {"description":"Single-use token <a href=\"https://stripe.com/docs/stripe.js\">provided by Stripe's client-side card capture SDK</a>, in response to your capturing a card into the Stripe vault. This token will be used by BillForward to find the tokenized card within the Stripe vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card.","verbs":["POST"]}

		Required: true
		*/
		StripeToken string `json:"stripeToken,omitempty"`
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		return err
	}

	m.accountIdField = data.AccountID
	m.companyNameField = data.CompanyName
	m.defaultPaymentMethodField = data.DefaultPaymentMethod
	m.emailField = data.Email
	m.firstNameField = data.FirstName
	m.gatewayField = data.Gateway
	m.lastNameField = data.LastName
	m.mobileField = data.Mobile
	m.organizationIdField = data.OrganizationID
	m.CardID = data.CardID
	m.CustomerID = data.CustomerID
	m.StripeToken = data.StripeToken

	return nil
}

// MarshalJSON marshals this polymorphic type to a JSON structure
func (m StripeAuthCaptureRequest) MarshalJSON() ([]byte, error) {
	var data struct {
		AtType string `json:"@type,omitempty"`

		AccountID *string `json:"accountID,omitempty"`

		CompanyName *string `json:"companyName,omitempty"`

		DefaultPaymentMethod *bool `json:"defaultPaymentMethod,omitempty"`

		Email *string `json:"email,omitempty"`

		FirstName *string `json:"firstName,omitempty"`

		Gateway *string `json:"gateway,omitempty"`

		LastName *string `json:"lastName,omitempty"`

		Mobile *string `json:"mobile,omitempty"`

		OrganizationID *string `json:"organizationID,omitempty"`

		/* {"description":"ID of the captured Card in Stripe. This can be provided as well as &mdash; or instead of &mdash; the one-use `stripeToken`, to lead BillForward to the card tokenized within the Stripe vault.","verbs":["POST"]}
		 */
		CardID *string `json:"cardID,omitempty"`

		/* CustomerID customer ID
		 */
		CustomerID *string `json:"customerID,omitempty"`

		/* {"description":"Single-use token <a href=\"https://stripe.com/docs/stripe.js\">provided by Stripe's client-side card capture SDK</a>, in response to your capturing a card into the Stripe vault. This token will be used by BillForward to find the tokenized card within the Stripe vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card.","verbs":["POST"]}

		Required: true
		*/
		StripeToken string `json:"stripeToken,omitempty"`
	}

	data.AccountID = m.accountIdField
	data.CompanyName = m.companyNameField
	data.DefaultPaymentMethod = m.defaultPaymentMethodField
	data.Email = m.emailField
	data.FirstName = m.firstNameField
	data.Gateway = m.gatewayField
	data.LastName = m.lastNameField
	data.Mobile = m.mobileField
	data.OrganizationID = m.organizationIdField
	data.CardID = m.CardID
	data.CustomerID = m.CustomerID
	data.StripeToken = m.StripeToken
	data.AtType = "StripeAuthCaptureRequest"
	return json.Marshal(data)
}

// Validate validates this stripe auth capture request
func (m *StripeAuthCaptureRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var stripeAuthCaptureRequestGatewayEnum []interface{}

func (m *StripeAuthCaptureRequest) validateGatewayEnum(path, location string, value string) error {
	if stripeAuthCaptureRequestGatewayEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Balanced","Braintree","Cybersource","Paypal","Stripe","AuthorizeNet","Spreedly","SagePay","TrustCommerce","Payvision","Kash"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			stripeAuthCaptureRequestGatewayEnum = append(stripeAuthCaptureRequestGatewayEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, stripeAuthCaptureRequestGatewayEnum); err != nil {
		return err
	}
	return nil
}

func (m *StripeAuthCaptureRequest) validateGateway(formats strfmt.Registry) error {

	if err := m.validateGatewayEnum("gateway", "body", *m.Gateway()); err != nil {
		return err
	}

	return nil
}

func (m *StripeAuthCaptureRequest) validateStripeToken(formats strfmt.Registry) error {

	if err := validate.RequiredString("stripeToken", "body", string(m.StripeToken)); err != nil {
		return err
	}

	return nil
}
