// Code generated by ent, DO NOT EDIT.

package db

import (
	"time"

	"github.com/openmeterio/openmeter/internal/productcatalog/postgres_connector/ent/db/feature"
	"github.com/openmeterio/openmeter/internal/productcatalog/postgres_connector/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	featureMixin := schema.Feature{}.Mixin()
	featureMixinFields0 := featureMixin[0].Fields()
	_ = featureMixinFields0
	featureMixinFields1 := featureMixin[1].Fields()
	_ = featureMixinFields1
	featureFields := schema.Feature{}.Fields()
	_ = featureFields
	// featureDescCreatedAt is the schema descriptor for created_at field.
	featureDescCreatedAt := featureMixinFields1[0].Descriptor()
	// feature.DefaultCreatedAt holds the default value on creation for the created_at field.
	feature.DefaultCreatedAt = featureDescCreatedAt.Default.(func() time.Time)
	// featureDescUpdatedAt is the schema descriptor for updated_at field.
	featureDescUpdatedAt := featureMixinFields1[1].Descriptor()
	// feature.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feature.DefaultUpdatedAt = featureDescUpdatedAt.Default.(func() time.Time)
	// feature.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feature.UpdateDefaultUpdatedAt = featureDescUpdatedAt.UpdateDefault.(func() time.Time)
	// featureDescNamespace is the schema descriptor for namespace field.
	featureDescNamespace := featureFields[0].Descriptor()
	// feature.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	feature.NamespaceValidator = featureDescNamespace.Validators[0].(func(string) error)
	// featureDescName is the schema descriptor for name field.
	featureDescName := featureFields[1].Descriptor()
	// feature.NameValidator is a validator for the "name" field. It is called by the builders before save.
	feature.NameValidator = featureDescName.Validators[0].(func(string) error)
	// featureDescMeterSlug is the schema descriptor for meter_slug field.
	featureDescMeterSlug := featureFields[2].Descriptor()
	// feature.MeterSlugValidator is a validator for the "meter_slug" field. It is called by the builders before save.
	feature.MeterSlugValidator = featureDescMeterSlug.Validators[0].(func(string) error)
	// featureDescArchived is the schema descriptor for archived field.
	featureDescArchived := featureFields[4].Descriptor()
	// feature.DefaultArchived holds the default value on creation for the archived field.
	feature.DefaultArchived = featureDescArchived.Default.(bool)
	// featureDescID is the schema descriptor for id field.
	featureDescID := featureMixinFields0[0].Descriptor()
	// feature.DefaultID holds the default value on creation for the id field.
	feature.DefaultID = featureDescID.Default.(func() string)
}