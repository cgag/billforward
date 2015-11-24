package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
FixedTerm

swagger:model FixedTerm
*/
type FixedTerm struct {

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "compoundUplift", "verbs":["POST","PUT","GET"] }The proportional INCREASE in price applied every time the fixed terms recur. e.g. 0.03 is a 3% increase. -0.5 is a 50% decrease. 3 is a 300% increase

	Required: true
	*/
	CompoundUplift float64 `json:"compoundUplift,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "Is the fixedTerm deleted.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Deleted bool `json:"deleted,omitempty"`

	/* fixedTermExpiryBehaviour

	Required: true
	*/
	ExpiryBehaviour string `json:"expiryBehaviour,omitempty"`

	/* expiry_time

	Required: true
	*/
	ExpiryTime strfmt.DateTime `json:"expiryTime,omitempty"`

	/* FixedTermDefinition fixed term definition

	Required: true
	*/
	FixedTermDefinition *MutableBillingEntity `json:"fixedTermDefinition,omitempty"`

	/* { "description" : "fixedTermDefinitionID", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	FixedTermDefinitionID string `json:"fixedTermDefinitionID,omitempty"`

	/* { "description" : "The number of sequential fixed terms previous to this one for the subscription (i.e. zero indexed 'fixedTermCount').", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	FixedTermNumber int32 `json:"fixedTermNumber,omitempty"`

	/* { "description" : "", "verbs":["GET", "PUT"] }
	 */
	ID string `json:"id,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "The ID of the organization associated with the amendment.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "The number of billing periods that this fixed term lasts for.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Periods int32 `json:"periods,omitempty"`

	/* productRatePlanAsOfTime

	Required: true
	*/
	ProductRatePlanAsOfTime strfmt.DateTime `json:"productRatePlanAsOfTime,omitempty"`

	/* start_time

	Required: true
	*/
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	/* state

	Required: true
	*/
	State string `json:"state,omitempty"`

	/* Subscription subscription
	 */
	Subscription *Subscription `json:"subscription,omitempty"`

	/* { "description" : "subscriptionID", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	SubscriptionID string `json:"subscriptionID,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this fixed term
func (m *FixedTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompoundUplift(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryBehaviour(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedTermDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedTermDefinitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedTermNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRatePlanAsOfTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fixedTermBillingEntityEnum []interface{}

func (m *FixedTerm) validateBillingEntityEnum(path, location string, value string) error {
	if fixedTermBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			fixedTermBillingEntityEnum = append(fixedTermBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, fixedTermBillingEntityEnum)
}

func (m *FixedTerm) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateCompoundUplift(formats strfmt.Registry) error {

	if err := validate.Required("compoundUplift", "body", float64(m.CompoundUplift)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

var fixedTermExpiryBehaviourEnum []interface{}

func (m *FixedTerm) validateExpiryBehaviourEnum(path, location string, value string) error {
	if fixedTermExpiryBehaviourEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["ExpireSubscription","EvergreenSubscription","RecurUplift","RecurLatestPricing"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			fixedTermExpiryBehaviourEnum = append(fixedTermExpiryBehaviourEnum, v)
		}
	}
	return validate.Enum(path, location, value, fixedTermExpiryBehaviourEnum)
}

func (m *FixedTerm) validateExpiryBehaviour(formats strfmt.Registry) error {

	if err := validate.Required("expiryBehaviour", "body", string(m.ExpiryBehaviour)); err != nil {
		return err
	}

	if err := m.validateExpiryBehaviourEnum("expiryBehaviour", "body", m.ExpiryBehaviour); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateExpiryTime(formats strfmt.Registry) error {

	if err := validate.Required("expiryTime", "body", strfmt.DateTime(m.ExpiryTime)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateFixedTermDefinition(formats strfmt.Registry) error {

	if m.FixedTermDefinition != nil {

		if err := m.FixedTermDefinition.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *FixedTerm) validateFixedTermDefinitionID(formats strfmt.Registry) error {

	if err := validate.Required("fixedTermDefinitionID", "body", string(m.FixedTermDefinitionID)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateFixedTermNumber(formats strfmt.Registry) error {

	if err := validate.Required("fixedTermNumber", "body", int32(m.FixedTermNumber)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validatePeriods(formats strfmt.Registry) error {

	if err := validate.Required("periods", "body", int32(m.Periods)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateProductRatePlanAsOfTime(formats strfmt.Registry) error {

	if err := validate.Required("productRatePlanAsOfTime", "body", strfmt.DateTime(m.ProductRatePlanAsOfTime)); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", strfmt.DateTime(m.StartTime)); err != nil {
		return err
	}

	return nil
}

var fixedTermStateEnum []interface{}

func (m *FixedTerm) validateStateEnum(path, location string, value string) error {
	if fixedTermStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["NeedsAmendments","Active","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			fixedTermStateEnum = append(fixedTermStateEnum, v)
		}
	}
	return validate.Enum(path, location, value, fixedTermStateEnum)
}

func (m *FixedTerm) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", string(m.State)); err != nil {
		return err
	}

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *FixedTerm) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", string(m.SubscriptionID)); err != nil {
		return err
	}

	return nil
}