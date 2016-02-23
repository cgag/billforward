package models

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*flatPricingComponent FlatPricingComponent flat pricing component

swagger:model flatPricingComponent
*/
type FlatPricingComponent struct {
	changedByField *string

	chargeModelField string

	chargeTypeField string

	createdField *strfmt.DateTime

	crmIdField *string

	defaultQuantityField int32

	descriptionField *string

	downgradeModeField *string

	idField string

	invoicingTypeField *string

	maxQuantityField *int32

	minQuantityField *int32

	nameField string

	organizationIdField string

	productRatePlanIdField string

	publicNameField *string

	tiersField []*PricingComponentTier

	unitOfMeasureField *UnitOfMeasure

	unitOfMeasureIdField string

	updatedField *strfmt.DateTime

	upgradeModeField *string

	validFromField strfmt.DateTime

	validTillField *strfmt.DateTime

	versionIdField string
}

func (m *FlatPricingComponent) AtType() string {
	return "flatPricingComponent"
}
func (m *FlatPricingComponent) SetAtType(val string) {

}

func (m *FlatPricingComponent) ChangedBy() *string {
	return m.changedByField
}
func (m *FlatPricingComponent) SetChangedBy(val *string) {
	m.changedByField = val
}

func (m *FlatPricingComponent) ChargeModel() string {
	return m.chargeModelField
}
func (m *FlatPricingComponent) SetChargeModel(val string) {
	m.chargeModelField = val
}

func (m *FlatPricingComponent) ChargeType() string {
	return m.chargeTypeField
}
func (m *FlatPricingComponent) SetChargeType(val string) {
	m.chargeTypeField = val
}

func (m *FlatPricingComponent) Created() *strfmt.DateTime {
	return m.createdField
}
func (m *FlatPricingComponent) SetCreated(val *strfmt.DateTime) {
	m.createdField = val
}

func (m *FlatPricingComponent) CrmID() *string {
	return m.crmIdField
}
func (m *FlatPricingComponent) SetCrmID(val *string) {
	m.crmIdField = val
}

func (m *FlatPricingComponent) DefaultQuantity() int32 {
	return m.defaultQuantityField
}
func (m *FlatPricingComponent) SetDefaultQuantity(val int32) {
	m.defaultQuantityField = val
}

func (m *FlatPricingComponent) Description() *string {
	return m.descriptionField
}
func (m *FlatPricingComponent) SetDescription(val *string) {
	m.descriptionField = val
}

func (m *FlatPricingComponent) DowngradeMode() *string {
	return m.downgradeModeField
}
func (m *FlatPricingComponent) SetDowngradeMode(val *string) {
	m.downgradeModeField = val
}

func (m *FlatPricingComponent) ID() string {
	return m.idField
}
func (m *FlatPricingComponent) SetID(val string) {
	m.idField = val
}

func (m *FlatPricingComponent) InvoicingType() *string {
	return m.invoicingTypeField
}
func (m *FlatPricingComponent) SetInvoicingType(val *string) {
	m.invoicingTypeField = val
}

func (m *FlatPricingComponent) MaxQuantity() *int32 {
	return m.maxQuantityField
}
func (m *FlatPricingComponent) SetMaxQuantity(val *int32) {
	m.maxQuantityField = val
}

func (m *FlatPricingComponent) MinQuantity() *int32 {
	return m.minQuantityField
}
func (m *FlatPricingComponent) SetMinQuantity(val *int32) {
	m.minQuantityField = val
}

func (m *FlatPricingComponent) Name() string {
	return m.nameField
}
func (m *FlatPricingComponent) SetName(val string) {
	m.nameField = val
}

func (m *FlatPricingComponent) OrganizationID() string {
	return m.organizationIdField
}
func (m *FlatPricingComponent) SetOrganizationID(val string) {
	m.organizationIdField = val
}

func (m *FlatPricingComponent) ProductRatePlanID() string {
	return m.productRatePlanIdField
}
func (m *FlatPricingComponent) SetProductRatePlanID(val string) {
	m.productRatePlanIdField = val
}

func (m *FlatPricingComponent) PublicName() *string {
	return m.publicNameField
}
func (m *FlatPricingComponent) SetPublicName(val *string) {
	m.publicNameField = val
}

func (m *FlatPricingComponent) Tiers() []*PricingComponentTier {
	return m.tiersField
}
func (m *FlatPricingComponent) SetTiers(val []*PricingComponentTier) {
	m.tiersField = val
}

func (m *FlatPricingComponent) UnitOfMeasure() *UnitOfMeasure {
	return m.unitOfMeasureField
}
func (m *FlatPricingComponent) SetUnitOfMeasure(val *UnitOfMeasure) {
	m.unitOfMeasureField = val
}

func (m *FlatPricingComponent) UnitOfMeasureID() string {
	return m.unitOfMeasureIdField
}
func (m *FlatPricingComponent) SetUnitOfMeasureID(val string) {
	m.unitOfMeasureIdField = val
}

func (m *FlatPricingComponent) Updated() *strfmt.DateTime {
	return m.updatedField
}
func (m *FlatPricingComponent) SetUpdated(val *strfmt.DateTime) {
	m.updatedField = val
}

func (m *FlatPricingComponent) UpgradeMode() *string {
	return m.upgradeModeField
}
func (m *FlatPricingComponent) SetUpgradeMode(val *string) {
	m.upgradeModeField = val
}

func (m *FlatPricingComponent) ValidFrom() strfmt.DateTime {
	return m.validFromField
}
func (m *FlatPricingComponent) SetValidFrom(val strfmt.DateTime) {
	m.validFromField = val
}

func (m *FlatPricingComponent) ValidTill() *strfmt.DateTime {
	return m.validTillField
}
func (m *FlatPricingComponent) SetValidTill(val *strfmt.DateTime) {
	m.validTillField = val
}

func (m *FlatPricingComponent) VersionID() string {
	return m.versionIdField
}
func (m *FlatPricingComponent) SetVersionID(val string) {
	m.versionIdField = val
}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *FlatPricingComponent) UnmarshalJSON(raw []byte) error {
	var data struct {
		AtType string `json:"@type,omitempty"`

		ChangedBy *string `json:"changedBy,omitempty"`

		ChargeModel string `json:"chargeModel,omitempty"`

		ChargeType string `json:"chargeType,omitempty"`

		Created *strfmt.DateTime `json:"created,omitempty"`

		CrmID *string `json:"crmID,omitempty"`

		DefaultQuantity int32 `json:"defaultQuantity,omitempty"`

		Description *string `json:"description,omitempty"`

		DowngradeMode *string `json:"downgradeMode,omitempty"`

		ID string `json:"id,omitempty"`

		InvoicingType *string `json:"invoicingType,omitempty"`

		MaxQuantity *int32 `json:"maxQuantity,omitempty"`

		MinQuantity *int32 `json:"minQuantity,omitempty"`

		Name string `json:"name,omitempty"`

		OrganizationID string `json:"organizationID,omitempty"`

		ProductRatePlanID string `json:"productRatePlanID,omitempty"`

		PublicName *string `json:"publicName,omitempty"`

		Tiers []*PricingComponentTier `json:"tiers,omitempty"`

		UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

		UnitOfMeasureID string `json:"unitOfMeasureID,omitempty"`

		Updated *strfmt.DateTime `json:"updated,omitempty"`

		UpgradeMode *string `json:"upgradeMode,omitempty"`

		ValidFrom strfmt.DateTime `json:"validFrom,omitempty"`

		ValidTill *strfmt.DateTime `json:"validTill,omitempty"`

		VersionID string `json:"versionID,omitempty"`
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		return err
	}

	m.changedByField = data.ChangedBy
	m.chargeModelField = data.ChargeModel
	m.chargeTypeField = data.ChargeType
	m.createdField = data.Created
	m.crmIdField = data.CrmID
	m.defaultQuantityField = data.DefaultQuantity
	m.descriptionField = data.Description
	m.downgradeModeField = data.DowngradeMode
	m.idField = data.ID
	m.invoicingTypeField = data.InvoicingType
	m.maxQuantityField = data.MaxQuantity
	m.minQuantityField = data.MinQuantity
	m.nameField = data.Name
	m.organizationIdField = data.OrganizationID
	m.productRatePlanIdField = data.ProductRatePlanID
	m.publicNameField = data.PublicName
	m.tiersField = data.Tiers
	m.unitOfMeasureField = data.UnitOfMeasure
	m.unitOfMeasureIdField = data.UnitOfMeasureID
	m.updatedField = data.Updated
	m.upgradeModeField = data.UpgradeMode
	m.validFromField = data.ValidFrom
	m.validTillField = data.ValidTill
	m.versionIdField = data.VersionID

	return nil
}

// MarshalJSON marshals this polymorphic type to a JSON structure
func (m FlatPricingComponent) MarshalJSON() ([]byte, error) {
	var data struct {
		AtType string `json:"@type,omitempty"`

		ChangedBy *string `json:"changedBy,omitempty"`

		ChargeModel string `json:"chargeModel,omitempty"`

		ChargeType string `json:"chargeType,omitempty"`

		Created *strfmt.DateTime `json:"created,omitempty"`

		CrmID *string `json:"crmID,omitempty"`

		DefaultQuantity int32 `json:"defaultQuantity,omitempty"`

		Description *string `json:"description,omitempty"`

		DowngradeMode *string `json:"downgradeMode,omitempty"`

		ID string `json:"id,omitempty"`

		InvoicingType *string `json:"invoicingType,omitempty"`

		MaxQuantity *int32 `json:"maxQuantity,omitempty"`

		MinQuantity *int32 `json:"minQuantity,omitempty"`

		Name string `json:"name,omitempty"`

		OrganizationID string `json:"organizationID,omitempty"`

		ProductRatePlanID string `json:"productRatePlanID,omitempty"`

		PublicName *string `json:"publicName,omitempty"`

		Tiers []*PricingComponentTier `json:"tiers,omitempty"`

		UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

		UnitOfMeasureID string `json:"unitOfMeasureID,omitempty"`

		Updated *strfmt.DateTime `json:"updated,omitempty"`

		UpgradeMode *string `json:"upgradeMode,omitempty"`

		ValidFrom strfmt.DateTime `json:"validFrom,omitempty"`

		ValidTill *strfmt.DateTime `json:"validTill,omitempty"`

		VersionID string `json:"versionID,omitempty"`
	}

	data.ChangedBy = m.changedByField
	data.ChargeModel = m.chargeModelField
	data.ChargeType = m.chargeTypeField
	data.Created = m.createdField
	data.CrmID = m.crmIdField
	data.DefaultQuantity = m.defaultQuantityField
	data.Description = m.descriptionField
	data.DowngradeMode = m.downgradeModeField
	data.ID = m.idField
	data.InvoicingType = m.invoicingTypeField
	data.MaxQuantity = m.maxQuantityField
	data.MinQuantity = m.minQuantityField
	data.Name = m.nameField
	data.OrganizationID = m.organizationIdField
	data.ProductRatePlanID = m.productRatePlanIdField
	data.PublicName = m.publicNameField
	data.Tiers = m.tiersField
	data.UnitOfMeasure = m.unitOfMeasureField
	data.UnitOfMeasureID = m.unitOfMeasureIdField
	data.Updated = m.updatedField
	data.UpgradeMode = m.upgradeModeField
	data.ValidFrom = m.validFromField
	data.ValidTill = m.validTillField
	data.VersionID = m.versionIdField
	data.AtType = "flatPricingComponent"
	return json.Marshal(data)
}

// Validate validates this flat pricing component
func (m *FlatPricingComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDowngradeMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoicingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var flatPricingComponentChargeModelEnum []interface{}

func (m *FlatPricingComponent) validateChargeModelEnum(path, location string, value string) error {
	if flatPricingComponentChargeModelEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["flat","tiered","tiered_volume"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			flatPricingComponentChargeModelEnum = append(flatPricingComponentChargeModelEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, flatPricingComponentChargeModelEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlatPricingComponent) validateChargeModel(formats strfmt.Registry) error {

	if err := validate.RequiredString("chargeModel", "body", string(m.ChargeModel())); err != nil {
		return err
	}

	if err := m.validateChargeModelEnum("chargeModel", "body", m.ChargeModel()); err != nil {
		return err
	}

	return nil
}

var flatPricingComponentChargeTypeEnum []interface{}

func (m *FlatPricingComponent) validateChargeTypeEnum(path, location string, value string) error {
	if flatPricingComponentChargeTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["setup","subscription","arrears","usage"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			flatPricingComponentChargeTypeEnum = append(flatPricingComponentChargeTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, flatPricingComponentChargeTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlatPricingComponent) validateChargeType(formats strfmt.Registry) error {

	if err := validate.RequiredString("chargeType", "body", string(m.ChargeType())); err != nil {
		return err
	}

	if err := m.validateChargeTypeEnum("chargeType", "body", m.ChargeType()); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateDefaultQuantity(formats strfmt.Registry) error {

	if err := validate.Required("defaultQuantity", "body", int32(m.DefaultQuantity())); err != nil {
		return err
	}

	return nil
}

var flatPricingComponentDowngradeModeEnum []interface{}

func (m *FlatPricingComponent) validateDowngradeModeEnum(path, location string, value string) error {
	if flatPricingComponentDowngradeModeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			flatPricingComponentDowngradeModeEnum = append(flatPricingComponentDowngradeModeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, flatPricingComponentDowngradeModeEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlatPricingComponent) validateDowngradeMode(formats strfmt.Registry) error {

	if err := m.validateDowngradeModeEnum("downgradeMode", "body", *m.DowngradeMode()); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID())); err != nil {
		return err
	}

	return nil
}

var flatPricingComponentInvoicingTypeEnum []interface{}

func (m *FlatPricingComponent) validateInvoicingTypeEnum(path, location string, value string) error {
	if flatPricingComponentInvoicingTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Immediate","Aggregated"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			flatPricingComponentInvoicingTypeEnum = append(flatPricingComponentInvoicingTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, flatPricingComponentInvoicingTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlatPricingComponent) validateInvoicingType(formats strfmt.Registry) error {

	if err := m.validateInvoicingTypeEnum("invoicingType", "body", *m.InvoicingType()); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name())); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.RequiredString("organizationID", "body", string(m.OrganizationID())); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateProductRatePlanID(formats strfmt.Registry) error {

	if err := validate.RequiredString("productRatePlanID", "body", string(m.ProductRatePlanID())); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateTiers(formats strfmt.Registry) error {

	for i := 0; i < len(m.Tiers()); i++ {

		if m.tiersField[i] != nil {

			if err := m.tiersField[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FlatPricingComponent) validateUnitOfMeasureID(formats strfmt.Registry) error {

	if err := validate.RequiredString("unitOfMeasureID", "body", string(m.UnitOfMeasureID())); err != nil {
		return err
	}

	return nil
}

var flatPricingComponentUpgradeModeEnum []interface{}

func (m *FlatPricingComponent) validateUpgradeModeEnum(path, location string, value string) error {
	if flatPricingComponentUpgradeModeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			flatPricingComponentUpgradeModeEnum = append(flatPricingComponentUpgradeModeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, flatPricingComponentUpgradeModeEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlatPricingComponent) validateUpgradeMode(formats strfmt.Registry) error {

	if err := m.validateUpgradeModeEnum("upgradeMode", "body", *m.UpgradeMode()); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", strfmt.DateTime(m.ValidFrom())); err != nil {
		return err
	}

	return nil
}

func (m *FlatPricingComponent) validateVersionID(formats strfmt.Registry) error {

	if err := validate.RequiredString("versionID", "body", string(m.VersionID())); err != nil {
		return err
	}

	return nil
}
