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
Entity for requesting that an 'aggregating subscription' (i.e. a 'parent subscription' which collects the charges raised by many 'child subscriptions') be created.

swagger:model CreateAggregatingSubscriptionRequest
*/
type CreateAggregatingSubscriptionRequest struct {

	/* {"description":"ID of the BillForward Account who will own this aggregating subscription. You should ensure beforehand that the customer has had a BillForward Account created for them.","verbs":["POST"]}

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* {"default":false,"description":"Whether this 'aggregating subscription' should collect charges (starting now) from all other subscriptions (current and future) belonging to this BillForward Account.","verbs":["POST"]}
	 */
	AggregateAllSubscriptionsOnAccount bool `json:"aggregateAllSubscriptionsOnAccount,omitempty"`

	/* {"default":"(empty list)","description":"[Required if and only if `productRatePlan` is omitted] List of components whose prices should be recalculated upon invoice aggregation. For example: two subscriptions' individual consumptions may neither of them be large enough to achieve bulk buy discounts. When aggregated, though, the same two subscriptions' consumption may add up to a quantity which does merit a bulk buy discount within your tiering system.","verbs":["POST"]}
	 */
	AggregatingComponents []*CreateAggregatingComponentRequest `json:"aggregatingComponents,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] The currency of the product-rate-plan &mdash; as specified by a three-character ISO 4217 currency code (i.e. USD).","verbs":["POST"]}

	Required: true
	*/
	Currency string `json:"currency,omitempty"`

	/* {"default":"(null)","description":"Description of the created subscription. This is primarily for your benefit &mdash; for example, you could write here the mechanism through which you obtained this customer. (e.g. 'Business signed up using BUSYGUYS coupon, at management trade show').","verbs":["POST"]}
	 */
	Description string `json:"description,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] Number of length-measures which constitute the rate plan's period.","verbs":["POST"]}
	 */
	Duration int32 `json:"duration,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] Measure describing the magnitude of the rate plan's period.","verbs":["POST"]}
	 */
	DurationPeriod string `json:"durationPeriod,omitempty"`

	/* {"default":"(Subscription will be named after the rate plan to which the subscription subscribes)","description":"Name of the created subscription. This is primarily for your benefit &mdash; for example, to enable you to identify subscriptions at a glance in the BillForward web interface (e.g. 'BusinessCorp subscriptions, care of Mr Business (mr@busy.com)').","verbs":["POST"]}
	 */
	Name string `json:"name,omitempty"`

	/* {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested Subscription should be created. If omitted, this will be auto-populated using your authentication credentials.","verbs":["POST"]}
	 */
	OrganizationID string `json:"organizationID,omitempty"`

	/* {"description":"ID of the rate plan to which the subscription will be subscribing. If omitted: it will be assumed you wish to create a new rate plan as part of this request &mdash; this subscription will subscribe to that newly-created rate plan.","verbs":["POST"]}
	 */
	ProductRatePlan string `json:"productRatePlan,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] The frequency of the rate plan &mdash; either recurring or non-recurring.","verbs":["POST"]}
	 */
	ProductType string `json:"productType,omitempty"`

	/* {"default":"(ServerNow upon receiving request)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should enter its first service period.","verbs":["POST"]}
	 */
	Start *strfmt.DateTime `json:"start,omitempty"`

	/* {"default":"Provisioned","description":"The state in which the created subscription will begin.<br><span class=\"label label-default\">Provisioned</span> &mdash; The subscription will wait (without raising any invoices or beginning its service) until explicit action is taken to change its state.<br><span class=\"label label-default\">AwaitingPayment</span> &mdash; The subscription is activated. After `start` time is surpassed, it will begin service and raise its first invoice.","verbs":["POST"]}
	 */
	State string `json:"state,omitempty"`
}

// Validate validates this create aggregating subscription request
func (m *CreateAggregatingSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurationPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestBillingEntityEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateBillingEntityEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestBillingEntityEnum = append(createAggregatingSubscriptionRequestBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, createAggregatingSubscriptionRequestBillingEntityEnum)
}

func (m *CreateAggregatingSubscriptionRequest) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", string(m.Currency)); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestDurationPeriodEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateDurationPeriodEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestDurationPeriodEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["minutes","days","months","years"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestDurationPeriodEnum = append(createAggregatingSubscriptionRequestDurationPeriodEnum, v)
		}
	}
	return validate.Enum(path, location, value, createAggregatingSubscriptionRequestDurationPeriodEnum)
}

func (m *CreateAggregatingSubscriptionRequest) validateDurationPeriod(formats strfmt.Registry) error {

	if err := m.validateDurationPeriodEnum("durationPeriod", "body", m.DurationPeriod); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestProductTypeEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateProductTypeEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestProductTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["nonrecurring","recurring"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestProductTypeEnum = append(createAggregatingSubscriptionRequestProductTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, createAggregatingSubscriptionRequestProductTypeEnum)
}

func (m *CreateAggregatingSubscriptionRequest) validateProductType(formats strfmt.Registry) error {

	if err := m.validateProductTypeEnum("productType", "body", m.ProductType); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestStateEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateStateEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Trial","Provisioned","Paid","AwaitingPayment","Cancelled","Failed","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestStateEnum = append(createAggregatingSubscriptionRequestStateEnum, v)
		}
	}
	return validate.Enum(path, location, value, createAggregatingSubscriptionRequestStateEnum)
}

func (m *CreateAggregatingSubscriptionRequest) validateState(formats strfmt.Registry) error {

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}
