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
Product

swagger:model Product
*/
type Product struct {

	/* {"description":"","verbs":[]}
	 */
	AccountID string `json:"accountID,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* {"default":"true","description":"Whether invoices are created if they have a zero valued cost before any discounts are applied.","verbs":["POST","PUT","GET"] }

	Required: true
	*/
	CreateZeroValuedInvoices bool `json:"createZeroValuedInvoices,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* Crmid crmid
	 */
	Crmid string `json:"crmid,omitempty"`

	/* {"description":"","verbs":["GET"]}

	Required: true
	*/
	Deleted bool `json:"deleted,omitempty"`

	/* {"description":"A description &mdash; for your benefit &mdash; of the product. For example: you could explain what service this product entitles a customer to.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	Description string `json:"description,omitempty"`

	/* {"description":"Number of length-measures which constitute the product's period.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	Duration int32 `json:"duration,omitempty"`

	/* {"description":"Measure describing the magnitude of the product's period.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	DurationPeriod string `json:"durationPeriod,omitempty"`

	/* {"description":"","verbs":[]}
	 */
	GetstartDate strfmt.DateTime `json:"getstartDate,omitempty"`

	/* { "description" : "", "verbs":["GET", "PUT"] }
	 */
	ID string `json:"id,omitempty"`

	/* {"description":"A unique name &mdash; for your benefit &mdash; used to identify this product within BillForward. It should reflect the fact that this product confers some service to a customer (e.g. \"Gold membership\").<br>The product can also be defined by the frequency with which it recurs (e.g. \"Monthly Gold membership\").<br>Remember also that rate plans can override the timing prescribed by their product. If you intend to override that timing, you may consider the product's period duration to be an unimportant factor when it comes to naming it.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* {"default":"recurring","description":"The frequency of the product &mdash; either recurring or non-recurring.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	ProductType string `json:"productType,omitempty"`

	/* {"description":"A friendly non-unique name used to identify this product","verbs":["POST","PUT","GET"]}
	 */
	PublicName string `json:"publicName,omitempty"`

	/* {"default":0,"description":"Number of trial-length-measures which constitute the product's trial period","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	Trial int32 `json:"trial,omitempty"`

	/* {"default":"none","description":"Measure describing the magnitude of the product's trial period.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	TrialPeriod string `json:"trialPeriod,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateZeroValuedInvoices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurationPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrialPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var productBillingEntityEnum []interface{}

func (m *Product) validateBillingEntityEnum(path, location string, value string) error {
	if productBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			productBillingEntityEnum = append(productBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, productBillingEntityEnum)
}

func (m *Product) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateCreateZeroValuedInvoices(formats strfmt.Registry) error {

	if err := validate.Required("createZeroValuedInvoices", "body", bool(m.CreateZeroValuedInvoices)); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", int32(m.Duration)); err != nil {
		return err
	}

	return nil
}

var productDurationPeriodEnum []interface{}

func (m *Product) validateDurationPeriodEnum(path, location string, value string) error {
	if productDurationPeriodEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["minutes","days","months","years"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			productDurationPeriodEnum = append(productDurationPeriodEnum, v)
		}
	}
	return validate.Enum(path, location, value, productDurationPeriodEnum)
}

func (m *Product) validateDurationPeriod(formats strfmt.Registry) error {

	if err := validate.Required("durationPeriod", "body", string(m.DurationPeriod)); err != nil {
		return err
	}

	if err := m.validateDurationPeriodEnum("durationPeriod", "body", m.DurationPeriod); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

var productProductTypeEnum []interface{}

func (m *Product) validateProductTypeEnum(path, location string, value string) error {
	if productProductTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["nonrecurring","recurring"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			productProductTypeEnum = append(productProductTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, productProductTypeEnum)
}

func (m *Product) validateProductType(formats strfmt.Registry) error {

	if err := validate.Required("productType", "body", string(m.ProductType)); err != nil {
		return err
	}

	if err := m.validateProductTypeEnum("productType", "body", m.ProductType); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateTrial(formats strfmt.Registry) error {

	if err := validate.Required("trial", "body", int32(m.Trial)); err != nil {
		return err
	}

	return nil
}

var productTrialPeriodEnum []interface{}

func (m *Product) validateTrialPeriodEnum(path, location string, value string) error {
	if productTrialPeriodEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["none","minutes","days","months"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			productTrialPeriodEnum = append(productTrialPeriodEnum, v)
		}
	}
	return validate.Enum(path, location, value, productTrialPeriodEnum)
}

func (m *Product) validateTrialPeriod(formats strfmt.Registry) error {

	if err := validate.Required("trialPeriod", "body", string(m.TrialPeriod)); err != nil {
		return err
	}

	if err := m.validateTrialPeriodEnum("trialPeriod", "body", m.TrialPeriod); err != nil {
		return err
	}

	return nil
}
