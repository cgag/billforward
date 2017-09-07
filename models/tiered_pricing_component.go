// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TieredPricingComponent tiered pricing component
// swagger:model tieredPricingComponent
type TieredPricingComponent struct {
	changedByField string

	chargeModelField *string

	chargeTypeField *string

	createdField strfmt.DateTime

	crmIdField string

	defaultQuantityField *int32

	descriptionField string

	downgradeModeField string

	idField *string

	invoicingTypeField string

	maxQuantityField int32

	minQuantityField int32

	nameField *string

	organizationIdField *string

	productRatePlanIdField *string

	publicNameField string

	tiersField []*PricingComponentTier

	unitOfMeasureField *UnitOfMeasure

	unitOfMeasureIdField *string

	updatedField strfmt.DateTime

	upgradeModeField string

	validFromField *strfmt.DateTime

	validTillField *strfmt.DateTime

	versionIdField *string
}

func (m *TieredPricingComponent) AtType() string {
	return "tieredPricingComponent"
}
func (m *TieredPricingComponent) SetAtType(val string) {

}

func (m *TieredPricingComponent) ChangedBy() string {
	return m.changedByField
}
func (m *TieredPricingComponent) SetChangedBy(val string) {
	m.changedByField = val
}

func (m *TieredPricingComponent) ChargeModel() *string {
	return m.chargeModelField
}
func (m *TieredPricingComponent) SetChargeModel(val *string) {
	m.chargeModelField = val
}

func (m *TieredPricingComponent) ChargeType() *string {
	return m.chargeTypeField
}
func (m *TieredPricingComponent) SetChargeType(val *string) {
	m.chargeTypeField = val
}

func (m *TieredPricingComponent) Created() strfmt.DateTime {
	return m.createdField
}
func (m *TieredPricingComponent) SetCreated(val strfmt.DateTime) {
	m.createdField = val
}

func (m *TieredPricingComponent) CrmID() string {
	return m.crmIdField
}
func (m *TieredPricingComponent) SetCrmID(val string) {
	m.crmIdField = val
}

func (m *TieredPricingComponent) DefaultQuantity() *int32 {
	return m.defaultQuantityField
}
func (m *TieredPricingComponent) SetDefaultQuantity(val *int32) {
	m.defaultQuantityField = val
}

func (m *TieredPricingComponent) Description() string {
	return m.descriptionField
}
func (m *TieredPricingComponent) SetDescription(val string) {
	m.descriptionField = val
}

func (m *TieredPricingComponent) DowngradeMode() string {
	return m.downgradeModeField
}
func (m *TieredPricingComponent) SetDowngradeMode(val string) {
	m.downgradeModeField = val
}

func (m *TieredPricingComponent) ID() *string {
	return m.idField
}
func (m *TieredPricingComponent) SetID(val *string) {
	m.idField = val
}

func (m *TieredPricingComponent) InvoicingType() string {
	return m.invoicingTypeField
}
func (m *TieredPricingComponent) SetInvoicingType(val string) {
	m.invoicingTypeField = val
}

func (m *TieredPricingComponent) MaxQuantity() int32 {
	return m.maxQuantityField
}
func (m *TieredPricingComponent) SetMaxQuantity(val int32) {
	m.maxQuantityField = val
}

func (m *TieredPricingComponent) MinQuantity() int32 {
	return m.minQuantityField
}
func (m *TieredPricingComponent) SetMinQuantity(val int32) {
	m.minQuantityField = val
}

func (m *TieredPricingComponent) Name() *string {
	return m.nameField
}
func (m *TieredPricingComponent) SetName(val *string) {
	m.nameField = val
}

func (m *TieredPricingComponent) OrganizationID() *string {
	return m.organizationIdField
}
func (m *TieredPricingComponent) SetOrganizationID(val *string) {
	m.organizationIdField = val
}

func (m *TieredPricingComponent) ProductRatePlanID() *string {
	return m.productRatePlanIdField
}
func (m *TieredPricingComponent) SetProductRatePlanID(val *string) {
	m.productRatePlanIdField = val
}

func (m *TieredPricingComponent) PublicName() string {
	return m.publicNameField
}
func (m *TieredPricingComponent) SetPublicName(val string) {
	m.publicNameField = val
}

func (m *TieredPricingComponent) Tiers() []*PricingComponentTier {
	return m.tiersField
}
func (m *TieredPricingComponent) SetTiers(val []*PricingComponentTier) {
	m.tiersField = val
}

func (m *TieredPricingComponent) UnitOfMeasure() *UnitOfMeasure {
	return m.unitOfMeasureField
}
func (m *TieredPricingComponent) SetUnitOfMeasure(val *UnitOfMeasure) {
	m.unitOfMeasureField = val
}

func (m *TieredPricingComponent) UnitOfMeasureID() *string {
	return m.unitOfMeasureIdField
}
func (m *TieredPricingComponent) SetUnitOfMeasureID(val *string) {
	m.unitOfMeasureIdField = val
}

func (m *TieredPricingComponent) Updated() strfmt.DateTime {
	return m.updatedField
}
func (m *TieredPricingComponent) SetUpdated(val strfmt.DateTime) {
	m.updatedField = val
}

func (m *TieredPricingComponent) UpgradeMode() string {
	return m.upgradeModeField
}
func (m *TieredPricingComponent) SetUpgradeMode(val string) {
	m.upgradeModeField = val
}

func (m *TieredPricingComponent) ValidFrom() *strfmt.DateTime {
	return m.validFromField
}
func (m *TieredPricingComponent) SetValidFrom(val *strfmt.DateTime) {
	m.validFromField = val
}

func (m *TieredPricingComponent) ValidTill() *strfmt.DateTime {
	return m.validTillField
}
func (m *TieredPricingComponent) SetValidTill(val *strfmt.DateTime) {
	m.validTillField = val
}

func (m *TieredPricingComponent) VersionID() *string {
	return m.versionIdField
}
func (m *TieredPricingComponent) SetVersionID(val *string) {
	m.versionIdField = val
}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *TieredPricingComponent) UnmarshalJSON(raw []byte) error {
	var data struct {
		AtType string `json:"@type"`

		ChangedBy string `json:"changedBy,omitempty"`

		ChargeModel *string `json:"chargeModel"`

		ChargeType *string `json:"chargeType"`

		Created strfmt.DateTime `json:"created,omitempty"`

		CrmID string `json:"crmID,omitempty"`

		DefaultQuantity *int32 `json:"defaultQuantity"`

		Description string `json:"description,omitempty"`

		DowngradeMode string `json:"downgradeMode,omitempty"`

		ID *string `json:"id"`

		InvoicingType string `json:"invoicingType,omitempty"`

		MaxQuantity int32 `json:"maxQuantity,omitempty"`

		MinQuantity int32 `json:"minQuantity,omitempty"`

		Name *string `json:"name"`

		OrganizationID *string `json:"organizationID"`

		ProductRatePlanID *string `json:"productRatePlanID"`

		PublicName string `json:"publicName,omitempty"`

		Tiers []*PricingComponentTier `json:"tiers"`

		UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

		UnitOfMeasureID *string `json:"unitOfMeasureID"`

		Updated strfmt.DateTime `json:"updated,omitempty"`

		UpgradeMode string `json:"upgradeMode,omitempty"`

		ValidFrom *strfmt.DateTime `json:"validFrom"`

		ValidTill *strfmt.DateTime `json:"validTill,omitempty"`

		VersionID *string `json:"versionID"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
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
func (m TieredPricingComponent) MarshalJSON() ([]byte, error) {
	var data struct {
		AtType string `json:"@type"`

		ChangedBy string `json:"changedBy,omitempty"`

		ChargeModel *string `json:"chargeModel"`

		ChargeType *string `json:"chargeType"`

		Created strfmt.DateTime `json:"created,omitempty"`

		CrmID string `json:"crmID,omitempty"`

		DefaultQuantity *int32 `json:"defaultQuantity"`

		Description string `json:"description,omitempty"`

		DowngradeMode string `json:"downgradeMode,omitempty"`

		ID *string `json:"id"`

		InvoicingType string `json:"invoicingType,omitempty"`

		MaxQuantity int32 `json:"maxQuantity,omitempty"`

		MinQuantity int32 `json:"minQuantity,omitempty"`

		Name *string `json:"name"`

		OrganizationID *string `json:"organizationID"`

		ProductRatePlanID *string `json:"productRatePlanID"`

		PublicName string `json:"publicName,omitempty"`

		Tiers []*PricingComponentTier `json:"tiers"`

		UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

		UnitOfMeasureID *string `json:"unitOfMeasureID"`

		Updated strfmt.DateTime `json:"updated,omitempty"`

		UpgradeMode string `json:"upgradeMode,omitempty"`

		ValidFrom *strfmt.DateTime `json:"validFrom"`

		ValidTill *strfmt.DateTime `json:"validTill,omitempty"`

		VersionID *string `json:"versionID"`
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
	data.AtType = "tieredPricingComponent"
	return json.Marshal(data)
}

// Validate validates this tiered pricing component
func (m *TieredPricingComponent) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUnitOfMeasure(formats); err != nil {
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

var tieredPricingComponentTypeChargeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["flat","tiered","tiered_volume"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tieredPricingComponentTypeChargeModelPropEnum = append(tieredPricingComponentTypeChargeModelPropEnum, v)
	}
}

// property enum
func (m *TieredPricingComponent) validateChargeModelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tieredPricingComponentTypeChargeModelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TieredPricingComponent) validateChargeModel(formats strfmt.Registry) error {

	if err := validate.Required("chargeModel", "body", m.ChargeModel()); err != nil {
		return err
	}

	// value enum
	if err := m.validateChargeModelEnum("chargeModel", "body", *m.ChargeModel()); err != nil {
		return err
	}

	return nil
}

var tieredPricingComponentTypeChargeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["setup","subscription","arrears","usage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tieredPricingComponentTypeChargeTypePropEnum = append(tieredPricingComponentTypeChargeTypePropEnum, v)
	}
}

// property enum
func (m *TieredPricingComponent) validateChargeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tieredPricingComponentTypeChargeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TieredPricingComponent) validateChargeType(formats strfmt.Registry) error {

	if err := validate.Required("chargeType", "body", m.ChargeType()); err != nil {
		return err
	}

	// value enum
	if err := m.validateChargeTypeEnum("chargeType", "body", *m.ChargeType()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateDefaultQuantity(formats strfmt.Registry) error {

	if err := validate.Required("defaultQuantity", "body", m.DefaultQuantity()); err != nil {
		return err
	}

	return nil
}

var tieredPricingComponentTypeDowngradeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tieredPricingComponentTypeDowngradeModePropEnum = append(tieredPricingComponentTypeDowngradeModePropEnum, v)
	}
}

// property enum
func (m *TieredPricingComponent) validateDowngradeModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tieredPricingComponentTypeDowngradeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TieredPricingComponent) validateDowngradeMode(formats strfmt.Registry) error {

	if swag.IsZero(m.DowngradeMode()) { // not required
		return nil
	}

	// value enum
	if err := m.validateDowngradeModeEnum("downgradeMode", "body", m.DowngradeMode()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

var tieredPricingComponentTypeInvoicingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Immediate","Aggregated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tieredPricingComponentTypeInvoicingTypePropEnum = append(tieredPricingComponentTypeInvoicingTypePropEnum, v)
	}
}

// property enum
func (m *TieredPricingComponent) validateInvoicingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tieredPricingComponentTypeInvoicingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TieredPricingComponent) validateInvoicingType(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoicingType()) { // not required
		return nil
	}

	// value enum
	if err := m.validateInvoicingTypeEnum("invoicingType", "body", m.InvoicingType()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateProductRatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("productRatePlanID", "body", m.ProductRatePlanID()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateTiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Tiers()) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiers()); i++ {

		if swag.IsZero(m.tiersField[i]) { // not required
			continue
		}

		if m.tiersField[i] != nil {

			if err := m.tiersField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TieredPricingComponent) validateUnitOfMeasure(formats strfmt.Registry) error {

	if swag.IsZero(m.UnitOfMeasure()) { // not required
		return nil
	}

	if m.UnitOfMeasure() != nil {

		if err := m.UnitOfMeasure().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unitOfMeasure")
			}
			return err
		}
	}

	return nil
}

func (m *TieredPricingComponent) validateUnitOfMeasureID(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasureID", "body", m.UnitOfMeasureID()); err != nil {
		return err
	}

	return nil
}

var tieredPricingComponentTypeUpgradeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tieredPricingComponentTypeUpgradeModePropEnum = append(tieredPricingComponentTypeUpgradeModePropEnum, v)
	}
}

// property enum
func (m *TieredPricingComponent) validateUpgradeModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tieredPricingComponentTypeUpgradeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TieredPricingComponent) validateUpgradeMode(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradeMode()) { // not required
		return nil
	}

	// value enum
	if err := m.validateUpgradeModeEnum("upgradeMode", "body", m.UpgradeMode()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom()); err != nil {
		return err
	}

	return nil
}

func (m *TieredPricingComponent) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("versionID", "body", m.VersionID()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TieredPricingComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TieredPricingComponent) UnmarshalBinary(b []byte) error {
	var res TieredPricingComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
