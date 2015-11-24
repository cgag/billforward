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
The aggregating component defines a component which should be re-priced upon invoice aggregation

swagger:model AggregatingComponent
*/
type AggregatingComponent struct {

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* {"default":false,"description":"Whether the AggregatingComponent has been retired.","verbs":["GET"]}
	 */
	Deleted bool `json:"deleted,omitempty"`

	/* { "description" : "", "verbs":["GET", "PUT"] }
	 */
	ID string `json:"id,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested pricing component should be created. If omitted: this will be auto-populated using your authentication credentials.","verbs":["POST","GET"]}
	 */
	OrganizationID string `json:"organizationID,omitempty"`

	/* {"description":"ID of the pricing component to which this AggregatingComponent's aggregation applies. The subscriber to the aggregating rate plan (which contains this AggregatingComponent), will consult its children at the end of each billing period, and collect from those children all charges whose pricing component matches this ID. Those charges' quantities will be counted, and used when calculating the price of consuming this AggregatingComponent. The aggregating subscription then raises a discount charge &mdash; to account for the more favourable price tiering that emerges when aggregating.","verbs":["POST"]}

	Required: true
	*/
	PricingComponentID string `json:"pricingComponentID,omitempty"`

	/* {"description":"Name of the pricing component to which this AggregatingComponent's aggregation applies. The subscriber to the aggregating rate plan (which contains this AggregatingComponent), will consult its children at the end of each billing period, and collect from those children all charges whose pricing component matches this ID. Those charges' quantities will be counted, and used when calculating the price of consuming this AggregatingComponent. The aggregating subscription then raises a discount charge &mdash; to account for the more favourable price tiering that emerges when aggregating.","verbs":["POST"]}

	Required: true
	*/
	PricingComponentName string `json:"pricingComponentName,omitempty"`

	/* { "description" : "The product-rate-plan associated with the aggregating-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ProductRatePlan *ProductRatePlan `json:"productRatePlan,omitempty"`

	/* {"description":"ID of the rate plan upon which this AggregatingComponent is defined.","verbs":["POST","GET"]}
	 */
	ProductRatePlanID string `json:"productRatePlanID,omitempty"`

	/* {"description":"Name of the rate plan upon which this AggregatingComponent is defined.","verbs":["POST","GET"]}
	 */
	ProductRatePlanName string `json:"productRatePlanName,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this aggregating component
func (m *AggregatingComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingComponentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingComponentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aggregatingComponentBillingEntityEnum []interface{}

func (m *AggregatingComponent) validateBillingEntityEnum(path, location string, value string) error {
	if aggregatingComponentBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			aggregatingComponentBillingEntityEnum = append(aggregatingComponentBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, aggregatingComponentBillingEntityEnum)
}

func (m *AggregatingComponent) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *AggregatingComponent) validatePricingComponentID(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentID", "body", string(m.PricingComponentID)); err != nil {
		return err
	}

	return nil
}

func (m *AggregatingComponent) validatePricingComponentName(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentName", "body", string(m.PricingComponentName)); err != nil {
		return err
	}

	return nil
}

func (m *AggregatingComponent) validateProductRatePlan(formats strfmt.Registry) error {

	if m.ProductRatePlan != nil {

		if err := m.ProductRatePlan.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
