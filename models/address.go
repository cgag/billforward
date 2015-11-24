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
Address

swagger:model Address
*/
type Address struct {

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	AddressLine1 string `json:"addressLine1,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine2 string `json:"addressLine2,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine3 string `json:"addressLine3,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	City string `json:"city,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Country string `json:"country,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "", "verbs":["GET"] }
	 */
	Deleted bool `json:"deleted,omitempty"`

	/* { "description" : "ID of the address.", "verbs":["POST","PUT","GET"] }
	 */
	ID string `json:"id,omitempty"`

	/* { "description" : "Phone number", "verbs":["POST","PUT","GET"] }
	 */
	Landline string `json:"landline,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "", "verbs":[] }
	 */
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "ZIP code or postcode.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Postcode string `json:"postcode,omitempty"`

	/* { "description" : "Is this the primary, default, address for the associated profile?", "verbs":["POST","PUT","GET"] }
	 */
	PrimaryAddress bool `json:"primaryAddress,omitempty"`

	/* Profile profile
	 */
	Profile *Profile `json:"profile,omitempty"`

	/* { "description" : "", "verbs":["GET"] }
	 */
	ProfileID string `json:"profileID,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Province string `json:"province,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostcode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvince(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("addressLine1", "body", string(m.AddressLine1)); err != nil {
		return err
	}

	return nil
}

var addressBillingEntityEnum []interface{}

func (m *Address) validateBillingEntityEnum(path, location string, value string) error {
	if addressBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			addressBillingEntityEnum = append(addressBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, addressBillingEntityEnum)
}

func (m *Address) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", string(m.City)); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", string(m.Country)); err != nil {
		return err
	}

	return nil
}

func (m *Address) validatePostcode(formats strfmt.Registry) error {

	if err := validate.Required("postcode", "body", string(m.Postcode)); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateProvince(formats strfmt.Registry) error {

	if err := validate.Required("province", "body", string(m.Province)); err != nil {
		return err
	}

	return nil
}
