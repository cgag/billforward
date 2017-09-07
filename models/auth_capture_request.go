// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthCaptureRequest [Note: this request can be built automatically by our client-side card capture library, <a href="https://github.com/billforward/billforward-js">BillForward.js</a>; you should not need to interact with this API manually unless you have particularly bespoke requirements] This entity is used for requesting that BillForward produce a PaymentMethod, linked to a funding instrument you have vaulted in some payment gateway. The BillForward PaymentMethod will be associated with a BillForward Account of your choosing (or a newly-created Account, if preferred).
// swagger:discriminator AuthCaptureRequest @type
type AuthCaptureRequest interface {
	runtime.Validatable

	// at type
	// Required: true
	AtType() string
	SetAtType(string)

	// {"description":"ID of the BillForward Account with which you would like to associate the created payment method.<br>If omitted, BillForward will associate the created PaymentMethod with a newly-created Account, whose Profile details will be populated using billing information from the funding instrument.","verbs":["POST"]}
	AccountID() string
	SetAccountID(string)

	// {"description":"The name of the company of the customer from whose card a PaymentMethod is being produced. If provided: this metadata will be used to populate a Profile &mdash; should a BillForward Account be created by this request.","verbs":["POST"]}
	CompanyName() string
	SetCompanyName(string)

	// {"default":false,"description":"Whether the PaymentMethod produced by this request should be elected as the 'default' payment method for the concerned BillForward Account. Whichever PaymentMethod is elected as an Account's default payment method, will be consulted whenever payment is demanded of that Account (i.e. upon the execution of any of the Account's invoices).","verbs":["POST"]}
	DefaultPaymentMethod() *bool
	SetDefaultPaymentMethod(*bool)

	// {"description":"The email address of the customer from whose card a PaymentMethod is being produced. If provided: this metadata will be used to populate a Profile &mdash; should a BillForward Account be created by this request.","verbs":["POST"]}
	Email() string
	SetEmail(string)

	// {"description":"The first name of the customer from whose card a PaymentMethod is being produced. If provided: this metadata will be used to populate a Profile &mdash; should a BillForward Account be created by this request.","verbs":["POST"]}
	FirstName() string
	SetFirstName(string)

	// {"description":"The gateway with which your funding instrument has been vaulted.","verbs":["POST"]}
	Gateway() string
	SetGateway(string)

	// {"description":"The last name of the customer from whose card a PaymentMethod is being produced. If provided: this metadata will be used to populate a Profile &mdash; should a BillForward Account be created by this request.","verbs":["POST"]}
	LastName() string
	SetLastName(string)

	// {"description":"The mobile phone number of the customer from whose card a PaymentMethod is being produced. If provided: this metadata will be used to populate a Profile &mdash; should a BillForward Account be created by this request.","verbs":["POST"]}
	Mobile() string
	SetMobile(string)

	// {"description":"ID of the BillForward Organization within which the requested PaymentMethod should be created. If omitted, this will be auto-populated using your authentication credentials.","verbs":["POST"]}
	OrganizationID() string
	SetOrganizationID(string)
}

type authCaptureRequest struct {
	atTypeField string

	accountIdField string

	companyNameField string

	defaultPaymentMethodField *bool

	emailField string

	firstNameField string

	gatewayField string

	lastNameField string

	mobileField string

	organizationIdField string
}

func (m *authCaptureRequest) AtType() string {
	return "AuthCaptureRequest"
}
func (m *authCaptureRequest) SetAtType(val string) {

}

func (m *authCaptureRequest) AccountID() string {
	return m.accountIdField
}
func (m *authCaptureRequest) SetAccountID(val string) {
	m.accountIdField = val
}

func (m *authCaptureRequest) CompanyName() string {
	return m.companyNameField
}
func (m *authCaptureRequest) SetCompanyName(val string) {
	m.companyNameField = val
}

func (m *authCaptureRequest) DefaultPaymentMethod() *bool {
	return m.defaultPaymentMethodField
}
func (m *authCaptureRequest) SetDefaultPaymentMethod(val *bool) {
	m.defaultPaymentMethodField = val
}

func (m *authCaptureRequest) Email() string {
	return m.emailField
}
func (m *authCaptureRequest) SetEmail(val string) {
	m.emailField = val
}

func (m *authCaptureRequest) FirstName() string {
	return m.firstNameField
}
func (m *authCaptureRequest) SetFirstName(val string) {
	m.firstNameField = val
}

func (m *authCaptureRequest) Gateway() string {
	return m.gatewayField
}
func (m *authCaptureRequest) SetGateway(val string) {
	m.gatewayField = val
}

func (m *authCaptureRequest) LastName() string {
	return m.lastNameField
}
func (m *authCaptureRequest) SetLastName(val string) {
	m.lastNameField = val
}

func (m *authCaptureRequest) Mobile() string {
	return m.mobileField
}
func (m *authCaptureRequest) SetMobile(val string) {
	m.mobileField = val
}

func (m *authCaptureRequest) OrganizationID() string {
	return m.organizationIdField
}
func (m *authCaptureRequest) SetOrganizationID(val string) {
	m.organizationIdField = val
}

// Validate validates this auth capture request
func (m *authCaptureRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var authCaptureRequestTypeAtTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["StripeAuthCaptureRequest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authCaptureRequestTypeAtTypePropEnum = append(authCaptureRequestTypeAtTypePropEnum, v)
	}
}

const (
	// AuthCaptureRequestAtTypeStripeAuthCaptureRequest captures enum value "StripeAuthCaptureRequest"
	AuthCaptureRequestAtTypeStripeAuthCaptureRequest string = "StripeAuthCaptureRequest"
)

// prop value enum
func (m *authCaptureRequest) validateAtTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, authCaptureRequestTypeAtTypePropEnum); err != nil {
		return err
	}
	return nil
}

var authCaptureRequestTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Balanced","Braintree","Cybersource","Paypal","Stripe","AuthorizeNet","Spreedly","SagePay","TrustCommerce","Payvision","Kash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authCaptureRequestTypeGatewayPropEnum = append(authCaptureRequestTypeGatewayPropEnum, v)
	}
}

const (
	// AuthCaptureRequestGatewayBalanced captures enum value "Balanced"
	AuthCaptureRequestGatewayBalanced string = "Balanced"
	// AuthCaptureRequestGatewayBraintree captures enum value "Braintree"
	AuthCaptureRequestGatewayBraintree string = "Braintree"
	// AuthCaptureRequestGatewayCybersource captures enum value "Cybersource"
	AuthCaptureRequestGatewayCybersource string = "Cybersource"
	// AuthCaptureRequestGatewayPaypal captures enum value "Paypal"
	AuthCaptureRequestGatewayPaypal string = "Paypal"
	// AuthCaptureRequestGatewayStripe captures enum value "Stripe"
	AuthCaptureRequestGatewayStripe string = "Stripe"
	// AuthCaptureRequestGatewayAuthorizeNet captures enum value "AuthorizeNet"
	AuthCaptureRequestGatewayAuthorizeNet string = "AuthorizeNet"
	// AuthCaptureRequestGatewaySpreedly captures enum value "Spreedly"
	AuthCaptureRequestGatewaySpreedly string = "Spreedly"
	// AuthCaptureRequestGatewaySagePay captures enum value "SagePay"
	AuthCaptureRequestGatewaySagePay string = "SagePay"
	// AuthCaptureRequestGatewayTrustCommerce captures enum value "TrustCommerce"
	AuthCaptureRequestGatewayTrustCommerce string = "TrustCommerce"
	// AuthCaptureRequestGatewayPayvision captures enum value "Payvision"
	AuthCaptureRequestGatewayPayvision string = "Payvision"
	// AuthCaptureRequestGatewayKash captures enum value "Kash"
	AuthCaptureRequestGatewayKash string = "Kash"
)

// prop value enum
func (m *authCaptureRequest) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, authCaptureRequestTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *authCaptureRequest) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway()) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway()); err != nil {
		return err
	}

	return nil
}

// UnmarshalAuthCaptureRequestSlice unmarshals polymorphic slices of AuthCaptureRequest
func UnmarshalAuthCaptureRequestSlice(reader io.Reader, consumer runtime.Consumer) ([]AuthCaptureRequest, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AuthCaptureRequest
	for _, element := range elements {
		obj, err := unmarshalAuthCaptureRequest(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAuthCaptureRequest unmarshals polymorphic AuthCaptureRequest
func UnmarshalAuthCaptureRequest(reader io.Reader, consumer runtime.Consumer) (AuthCaptureRequest, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAuthCaptureRequest(data, consumer)
}

func unmarshalAuthCaptureRequest(data []byte, consumer runtime.Consumer) (AuthCaptureRequest, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		AtType string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.AtType); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.AtType {
	case "AuthCaptureRequest":
		var result authCaptureRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "BraintreeAuthCaptureRequest":
		var result BraintreeAuthCaptureRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "StripeAuthCaptureRequest":
		var result StripeAuthCaptureRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}
