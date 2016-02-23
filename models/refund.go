package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Refund Refund

swagger:model Refund
*/
type Refund struct {

	/* { "description" : "Identifier of account to refund.", "verbs":["GET"] }
	 */
	AccountID *string `json:"accountID,omitempty"`

	/* { "description" : "Refunded amount", "verbs":[] }

	Required: true
	*/
	ActualRefundedValue float64 `json:"actualRefundedValue,omitempty"`

	/* { "description" : "", "verbs":[] }

	Required: true
	*/
	ActualValue float64 `json:"actualValue,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created *strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "Refund requested by this account", "verbs":["GET"] }
	 */
	CreatedBy *string `json:"createdBy,omitempty"`

	/* { "description" : "Currency of the refund specified by a three character ISO 4217 currency code.", "verbs":["GET"] }

	Required: true
	*/
	Currency string `json:"currency,omitempty"`

	/* { "description" : "", "verbs":["GET", "PUT"] }
	 */
	ID *string `json:"id,omitempty"`

	/* { "description" : "Invoice to refund.", "verbs":["POST","GET"] }
	 */
	InvoiceID *string `json:"invoiceID,omitempty"`

	/* { "description" : "Invoice payment to refund", "verbs":["POST", "GET"] }
	 */
	InvoicePaymentID *string `json:"invoicePaymentID,omitempty"`

	/* { "description" : "Nominal value refunded.", "verbs":[] }

	Required: true
	*/
	NominalRefundedValue float64 `json:"nominalRefundedValue,omitempty"`

	/* { "description" : "", "verbs":[] }

	Required: true
	*/
	NominalValue float64 `json:"nominalValue,omitempty"`

	/* { "description" : "", "verbs":[] }
	 */
	OrganizationID *string `json:"organizationID,omitempty"`

	/* { "description" : "Reference in the gateway to the payment being refunded.", "verbs":["GET"] }
	 */
	OriginalGatewayPaymentReference *string `json:"originalGatewayPaymentReference,omitempty"`

	/* { "description" : "Original payment being refunded.", "verbs":["GET"] }
	 */
	OriginalPaymentID *string `json:"originalPaymentID,omitempty"`

	/* { "description" : "ID of the receipt for the successful payment that this entity refunds.", "verbs":["POST","GET"] }
	 */
	OriginalReceiptID *string `json:"originalReceiptID,omitempty"`

	/* { "description" : "Payment method to refund from", "verbs":["GET"] }
	 */
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`

	/* { "description" : "The reason for the refund.", "verbs":["GET"] }
	 */
	Reason *string `json:"reason,omitempty"`

	/* { "description" : "ID of the receipt for this refund.", "verbs":["POST","GET"] }
	 */
	ReceiptID *string `json:"receiptID,omitempty"`

	/* { "description" : "When the refund was completed", "verbs":["GET"] }
	 */
	RefundCompleted *strfmt.DateTime `json:"refundCompleted,omitempty"`

	/* { "description" : "", "verbs":[] }
	 */
	RefundInvoicePaymentID *string `json:"refundInvoicePaymentID,omitempty"`

	/* {"default":"<span class=\"label label-default\">SingleAttempt</span>","description":The mechanism by which credit is returned to the customer:<br><span class=\"label label-default\">Void</span> &mdash; The original authorized payment is voided before capture.<br><span class=\"label label-default\">Refund</span> &mdash; A refund is issued against an already-captured payment.","verbs":["GET"]}

	Required: true
	*/
	RefundNature string `json:"refundNature,omitempty"`

	/* { "description" : "Identifier for the payment used to refund", "verbs":["GET"] }
	 */
	RefundPaymentID *string `json:"refundPaymentID,omitempty"`

	/* { "default" : "AwaitingRefund", "PUT_description" : "<span class=\"label label-default\">Pending</span> refunds can be set to <span class=\"label label-default\">AwaitingRefund</span> to initiate the refund or <span class=\"label label-default\">Cancelled</span> to stop the refund.", "description" : "Initially a refund is in the <span class=\"label label-default\">AwaitingRefund</span> state. Once the funds are successfully returned the state is <span class=\"label label-default\">Refunded</span>. If the refund fails or is rejected is it set as <span class=\"label label-default\">Failed</span>. Refunds can be set to a <span class=\"label label-default\">Pending</span> state to support authorization flows, and will leave the refund pending until updated to <span class=\"label label-default\">AwaitingRefund</span>. <span class=\"label label-default\">Cancelled</span> is when a refund will not be executed.", "verbs":["GET", "POST", "PUT"] }

	Required: true
	*/
	RefundState string `json:"refundState,omitempty"`

	/* { "description" : "This is the type of refund. Refunds are associated with either an invoice and payment, a payment or unreferenced.", "verbs":[] }

	Required: true
	*/
	RefundType string `json:"refundType,omitempty"`

	/* { "description" : "Value refunded", "verbs":["GET"] }
	 */
	Refunded *float64 `json:"refunded,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "Positive decimal representing the total value to refund. This is at most the amount un-refunded on the payment. If amount is the total payment amount is refunded.", "verbs":["POST","GET"] }
	 */
	Value *float64 `json:"value,omitempty"`
}

// Validate validates this refund
func (m *Refund) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualRefundedValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActualValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNominalRefundedValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNominalValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundNature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Refund) validateActualRefundedValue(formats strfmt.Registry) error {

	if err := validate.Required("actualRefundedValue", "body", float64(m.ActualRefundedValue)); err != nil {
		return err
	}

	return nil
}

func (m *Refund) validateActualValue(formats strfmt.Registry) error {

	if err := validate.Required("actualValue", "body", float64(m.ActualValue)); err != nil {
		return err
	}

	return nil
}

func (m *Refund) validateCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("currency", "body", string(m.Currency)); err != nil {
		return err
	}

	return nil
}

func (m *Refund) validateNominalRefundedValue(formats strfmt.Registry) error {

	if err := validate.Required("nominalRefundedValue", "body", float64(m.NominalRefundedValue)); err != nil {
		return err
	}

	return nil
}

func (m *Refund) validateNominalValue(formats strfmt.Registry) error {

	if err := validate.Required("nominalValue", "body", float64(m.NominalValue)); err != nil {
		return err
	}

	return nil
}

var refundRefundNatureEnum []interface{}

func (m *Refund) validateRefundNatureEnum(path, location string, value string) error {
	if refundRefundNatureEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Refund","Void"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			refundRefundNatureEnum = append(refundRefundNatureEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, refundRefundNatureEnum); err != nil {
		return err
	}
	return nil
}

func (m *Refund) validateRefundNature(formats strfmt.Registry) error {

	if err := validate.RequiredString("refundNature", "body", string(m.RefundNature)); err != nil {
		return err
	}

	if err := m.validateRefundNatureEnum("refundNature", "body", m.RefundNature); err != nil {
		return err
	}

	return nil
}

var refundRefundStateEnum []interface{}

func (m *Refund) validateRefundStateEnum(path, location string, value string) error {
	if refundRefundStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Pending","AwaitingRefund","Refunded","Cancelled","Failed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			refundRefundStateEnum = append(refundRefundStateEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, refundRefundStateEnum); err != nil {
		return err
	}
	return nil
}

func (m *Refund) validateRefundState(formats strfmt.Registry) error {

	if err := validate.RequiredString("refundState", "body", string(m.RefundState)); err != nil {
		return err
	}

	if err := m.validateRefundStateEnum("refundState", "body", m.RefundState); err != nil {
		return err
	}

	return nil
}

var refundRefundTypeEnum []interface{}

func (m *Refund) validateRefundTypeEnum(path, location string, value string) error {
	if refundRefundTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["InvoicePayment","Payment"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			refundRefundTypeEnum = append(refundRefundTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, refundRefundTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *Refund) validateRefundType(formats strfmt.Registry) error {

	if err := validate.RequiredString("refundType", "body", string(m.RefundType)); err != nil {
		return err
	}

	if err := m.validateRefundTypeEnum("refundType", "body", m.RefundType); err != nil {
		return err
	}

	return nil
}
