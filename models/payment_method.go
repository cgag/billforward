// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentMethod PaymentMethod
// swagger:model PaymentMethod
type PaymentMethod struct {

	// { "description" : "ID of the account associated with the payment-method.", "verbs":["POST","GET"] }
	// Required: true
	AccountID *string `json:"accountID"`

	// { "description" : "Name of the card holder", "verbs":["POST","PUT","GET"] }
	CardHolderName string `json:"cardHolderName,omitempty"`

	// { "description" : "Type of the card. In the case of card payment methods this is the payment type, for example VISA, MasterCard.", "verbs":["POST","GET"] }
	CardType string `json:"cardType,omitempty"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "CRM ID of the product-rate-plan.", "verbs":] }
	CrmID string `json:"crmID,omitempty"`

	// {"default" : "false", "description" : "Indicates if this is the default payment method for the account.", "verbs":["GET","POST","PUT"]  }
	DefaultPaymentMethod *bool `json:"defaultPaymentMethod,omitempty"`

	// {"default" : "false", "description" : "Indicates if a payment-method has been retired. If a payment-method has been retired it can still be retrieved using the appropriate flag on API requests. Generally payment-methods will be retired by customers wanting to remove the payment method from their account. Caution should be used when requested deleted payment methods.", "verbs":["GET"] }
	Deleted *bool `json:"deleted,omitempty"`

	// { "description" : "Description of the payment-method.", "verbs":["POST","PUT","GET"] }
	Description string `json:"description,omitempty"`

	// { "description" : "In the case of card payment methods this is the expiry date of the card, format should be MMYY including leading 0's. For example 0115 for January 2015.", "verbs":["POST","GET"] }
	// Required: true
	ExpiryDate *string `json:"expiryDate"`

	// expiry month
	ExpiryMonth int32 `json:"expiryMonth,omitempty"`

	// expiry year
	ExpiryYear int32 `json:"expiryYear,omitempty"`

	// first six
	FirstSix string `json:"firstSix,omitempty"`

	// { "description" : "Gateway type for payment-method.", "verbs":["POST","GET"] }
	// Required: true
	Gateway *string `json:"gateway"`

	// { "description" : "ID of the payment-method.", "verbs":["GET"] }
	ID string `json:"id,omitempty"`

	// {"description":"IP address associated with this payment method.","verbs":["POST","PUT","GET"]}
	IPAddress string `json:"ipAddress,omitempty"`

	// {"description":"Country of the IP address associated with this payment method.","verbs":["POST","PUT","GET"]}
	IPAddressCountry string `json:"ipAddressCountry,omitempty"`

	// last four
	LastFour string `json:"lastFour,omitempty"`

	// link ID
	// Required: true
	LinkID *string `json:"linkID"`

	// { "description" : "Name of the payment-method.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Name *string `json:"name"`

	// organization ID
	OrganizationID string `json:"organizationID,omitempty"`

	// province
	Province string `json:"province,omitempty"`

	// { "description" : "State of the payment-method.", "verbs":["POST","GET"] }
	// Required: true
	State *string `json:"state"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this payment method
func (m *PaymentMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExpiryDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinkID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentMethod) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *PaymentMethod) validateExpiryDate(formats strfmt.Registry) error {

	if err := validate.Required("expiryDate", "body", m.ExpiryDate); err != nil {
		return err
	}

	return nil
}

var paymentMethodTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cybersource_token","card_vault","paypal_simple","locustworld","free","coupon","credit_note","stripe","braintree","balanced","paypal","billforward_test","offline","trial","stripeACH","authorizeNet","spreedly","sagePay","trustCommerce","payvision","kash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentMethodTypeGatewayPropEnum = append(paymentMethodTypeGatewayPropEnum, v)
	}
}

const (
	// PaymentMethodGatewayCybersourceToken captures enum value "cybersource_token"
	PaymentMethodGatewayCybersourceToken string = "cybersource_token"
	// PaymentMethodGatewayCardVault captures enum value "card_vault"
	PaymentMethodGatewayCardVault string = "card_vault"
	// PaymentMethodGatewayPaypalSimple captures enum value "paypal_simple"
	PaymentMethodGatewayPaypalSimple string = "paypal_simple"
	// PaymentMethodGatewayLocustworld captures enum value "locustworld"
	PaymentMethodGatewayLocustworld string = "locustworld"
	// PaymentMethodGatewayFree captures enum value "free"
	PaymentMethodGatewayFree string = "free"
	// PaymentMethodGatewayCoupon captures enum value "coupon"
	PaymentMethodGatewayCoupon string = "coupon"
	// PaymentMethodGatewayCreditNote captures enum value "credit_note"
	PaymentMethodGatewayCreditNote string = "credit_note"
	// PaymentMethodGatewayStripe captures enum value "stripe"
	PaymentMethodGatewayStripe string = "stripe"
	// PaymentMethodGatewayBraintree captures enum value "braintree"
	PaymentMethodGatewayBraintree string = "braintree"
	// PaymentMethodGatewayBalanced captures enum value "balanced"
	PaymentMethodGatewayBalanced string = "balanced"
	// PaymentMethodGatewayPaypal captures enum value "paypal"
	PaymentMethodGatewayPaypal string = "paypal"
	// PaymentMethodGatewayBillforwardTest captures enum value "billforward_test"
	PaymentMethodGatewayBillforwardTest string = "billforward_test"
	// PaymentMethodGatewayOffline captures enum value "offline"
	PaymentMethodGatewayOffline string = "offline"
	// PaymentMethodGatewayTrial captures enum value "trial"
	PaymentMethodGatewayTrial string = "trial"
	// PaymentMethodGatewayStripeACH captures enum value "stripeACH"
	PaymentMethodGatewayStripeACH string = "stripeACH"
	// PaymentMethodGatewayAuthorizeNet captures enum value "authorizeNet"
	PaymentMethodGatewayAuthorizeNet string = "authorizeNet"
	// PaymentMethodGatewaySpreedly captures enum value "spreedly"
	PaymentMethodGatewaySpreedly string = "spreedly"
	// PaymentMethodGatewaySagePay captures enum value "sagePay"
	PaymentMethodGatewaySagePay string = "sagePay"
	// PaymentMethodGatewayTrustCommerce captures enum value "trustCommerce"
	PaymentMethodGatewayTrustCommerce string = "trustCommerce"
	// PaymentMethodGatewayPayvision captures enum value "payvision"
	PaymentMethodGatewayPayvision string = "payvision"
	// PaymentMethodGatewayKash captures enum value "kash"
	PaymentMethodGatewayKash string = "kash"
)

// prop value enum
func (m *PaymentMethod) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentMethodTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentMethod) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", *m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *PaymentMethod) validateLinkID(formats strfmt.Registry) error {

	if err := validate.Required("linkID", "body", m.LinkID); err != nil {
		return err
	}

	return nil
}

func (m *PaymentMethod) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var paymentMethodTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Active","Expiring","Expired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentMethodTypeStatePropEnum = append(paymentMethodTypeStatePropEnum, v)
	}
}

const (
	// PaymentMethodStatePending captures enum value "Pending"
	PaymentMethodStatePending string = "Pending"
	// PaymentMethodStateActive captures enum value "Active"
	PaymentMethodStateActive string = "Active"
	// PaymentMethodStateExpiring captures enum value "Expiring"
	PaymentMethodStateExpiring string = "Expiring"
	// PaymentMethodStateExpired captures enum value "Expired"
	PaymentMethodStateExpired string = "Expired"
)

// prop value enum
func (m *PaymentMethod) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentMethodTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentMethod) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentMethod) UnmarshalBinary(b []byte) error {
	var res PaymentMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
