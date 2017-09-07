// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BraintreeAuthCaptureRequest braintree auth capture request
// swagger:model BraintreeAuthCaptureRequest
type BraintreeAuthCaptureRequest struct {
	accountIdField string

	companyNameField string

	defaultPaymentMethodField *bool

	emailField string

	firstNameField string

	gatewayField string

	lastNameField string

	mobileField string

	organizationIdField string

	// {"description":"(Required when vaulting a PayPal payment method; otherwise optional)<br>A JSON string providing information about the device your customer used to fill out the card capture form. This information is inserted into your form by <a href=\"https://developers.braintreepayments.com/javascript+node/guides/advanced-fraud-tools/client-side\">braintree-data.js</a> &mdash; if and only if you use Braintree's drop-in form integrations. You should ideally provide it if you have one (it aids with fraud detection), but it is only mandatory in the case of PayPal payment method vaulting.","verbs":["POST"]}
	DeviceData string `json:"deviceData,omitempty"`

	// {"description":"One-use cryptographic nonce <a href=\"https://developers.braintreepayments.com/javascript+node/start/overview\">provided by Braintree's client-side card capture SDK</a>, in response to your capturing a card into the Braintree vault. This nonce will be used by BillForward to find the tokenized card within the Braintree vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card","verbs":["POST"]}
	// Required: true
	PaymentMethodNonce *string `json:"paymentMethodNonce"`
}

func (m *BraintreeAuthCaptureRequest) AtType() string {
	return "BraintreeAuthCaptureRequest"
}
func (m *BraintreeAuthCaptureRequest) SetAtType(val string) {

}

func (m *BraintreeAuthCaptureRequest) AccountID() string {
	return m.accountIdField
}
func (m *BraintreeAuthCaptureRequest) SetAccountID(val string) {
	m.accountIdField = val
}

func (m *BraintreeAuthCaptureRequest) CompanyName() string {
	return m.companyNameField
}
func (m *BraintreeAuthCaptureRequest) SetCompanyName(val string) {
	m.companyNameField = val
}

func (m *BraintreeAuthCaptureRequest) DefaultPaymentMethod() *bool {
	return m.defaultPaymentMethodField
}
func (m *BraintreeAuthCaptureRequest) SetDefaultPaymentMethod(val *bool) {
	m.defaultPaymentMethodField = val
}

func (m *BraintreeAuthCaptureRequest) Email() string {
	return m.emailField
}
func (m *BraintreeAuthCaptureRequest) SetEmail(val string) {
	m.emailField = val
}

func (m *BraintreeAuthCaptureRequest) FirstName() string {
	return m.firstNameField
}
func (m *BraintreeAuthCaptureRequest) SetFirstName(val string) {
	m.firstNameField = val
}

func (m *BraintreeAuthCaptureRequest) Gateway() string {
	return m.gatewayField
}
func (m *BraintreeAuthCaptureRequest) SetGateway(val string) {
	m.gatewayField = val
}

func (m *BraintreeAuthCaptureRequest) LastName() string {
	return m.lastNameField
}
func (m *BraintreeAuthCaptureRequest) SetLastName(val string) {
	m.lastNameField = val
}

func (m *BraintreeAuthCaptureRequest) Mobile() string {
	return m.mobileField
}
func (m *BraintreeAuthCaptureRequest) SetMobile(val string) {
	m.mobileField = val
}

func (m *BraintreeAuthCaptureRequest) OrganizationID() string {
	return m.organizationIdField
}
func (m *BraintreeAuthCaptureRequest) SetOrganizationID(val string) {
	m.organizationIdField = val
}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *BraintreeAuthCaptureRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		AtType string `json:"@type"`

		AccountID string `json:"accountID,omitempty"`

		CompanyName string `json:"companyName,omitempty"`

		DefaultPaymentMethod *bool `json:"defaultPaymentMethod,omitempty"`

		Email string `json:"email,omitempty"`

		FirstName string `json:"firstName,omitempty"`

		Gateway string `json:"gateway,omitempty"`

		LastName string `json:"lastName,omitempty"`

		Mobile string `json:"mobile,omitempty"`

		OrganizationID string `json:"organizationID,omitempty"`

		// {"description":"(Required when vaulting a PayPal payment method; otherwise optional)<br>A JSON string providing information about the device your customer used to fill out the card capture form. This information is inserted into your form by <a href=\"https://developers.braintreepayments.com/javascript+node/guides/advanced-fraud-tools/client-side\">braintree-data.js</a> &mdash; if and only if you use Braintree's drop-in form integrations. You should ideally provide it if you have one (it aids with fraud detection), but it is only mandatory in the case of PayPal payment method vaulting.","verbs":["POST"]}
		DeviceData string `json:"deviceData,omitempty"`

		// {"description":"One-use cryptographic nonce <a href=\"https://developers.braintreepayments.com/javascript+node/start/overview\">provided by Braintree's client-side card capture SDK</a>, in response to your capturing a card into the Braintree vault. This nonce will be used by BillForward to find the tokenized card within the Braintree vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card","verbs":["POST"]}
		// Required: true
		PaymentMethodNonce *string `json:"paymentMethodNonce"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
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
	m.DeviceData = data.DeviceData
	m.PaymentMethodNonce = data.PaymentMethodNonce

	return nil
}

// MarshalJSON marshals this polymorphic type to a JSON structure
func (m BraintreeAuthCaptureRequest) MarshalJSON() ([]byte, error) {
	var data struct {
		AtType string `json:"@type"`

		AccountID string `json:"accountID,omitempty"`

		CompanyName string `json:"companyName,omitempty"`

		DefaultPaymentMethod *bool `json:"defaultPaymentMethod,omitempty"`

		Email string `json:"email,omitempty"`

		FirstName string `json:"firstName,omitempty"`

		Gateway string `json:"gateway,omitempty"`

		LastName string `json:"lastName,omitempty"`

		Mobile string `json:"mobile,omitempty"`

		OrganizationID string `json:"organizationID,omitempty"`

		// {"description":"(Required when vaulting a PayPal payment method; otherwise optional)<br>A JSON string providing information about the device your customer used to fill out the card capture form. This information is inserted into your form by <a href=\"https://developers.braintreepayments.com/javascript+node/guides/advanced-fraud-tools/client-side\">braintree-data.js</a> &mdash; if and only if you use Braintree's drop-in form integrations. You should ideally provide it if you have one (it aids with fraud detection), but it is only mandatory in the case of PayPal payment method vaulting.","verbs":["POST"]}
		DeviceData string `json:"deviceData,omitempty"`

		// {"description":"One-use cryptographic nonce <a href=\"https://developers.braintreepayments.com/javascript+node/start/overview\">provided by Braintree's client-side card capture SDK</a>, in response to your capturing a card into the Braintree vault. This nonce will be used by BillForward to find the tokenized card within the Braintree vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card","verbs":["POST"]}
		// Required: true
		PaymentMethodNonce *string `json:"paymentMethodNonce"`
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
	data.DeviceData = m.DeviceData
	data.PaymentMethodNonce = m.PaymentMethodNonce
	data.AtType = "BraintreeAuthCaptureRequest"
	return json.Marshal(data)
}

// Validate validates this braintree auth capture request
func (m *BraintreeAuthCaptureRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethodNonce(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var braintreeAuthCaptureRequestTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Balanced","Braintree","Cybersource","Paypal","Stripe","AuthorizeNet","Spreedly","SagePay","TrustCommerce","Payvision","Kash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		braintreeAuthCaptureRequestTypeGatewayPropEnum = append(braintreeAuthCaptureRequestTypeGatewayPropEnum, v)
	}
}

// property enum
func (m *BraintreeAuthCaptureRequest) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, braintreeAuthCaptureRequestTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BraintreeAuthCaptureRequest) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway()) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway()); err != nil {
		return err
	}

	return nil
}

func (m *BraintreeAuthCaptureRequest) validatePaymentMethodNonce(formats strfmt.Registry) error {

	if err := validate.Required("paymentMethodNonce", "body", m.PaymentMethodNonce); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BraintreeAuthCaptureRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BraintreeAuthCaptureRequest) UnmarshalBinary(b []byte) error {
	var res BraintreeAuthCaptureRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
