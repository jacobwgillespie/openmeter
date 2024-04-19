// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CreditEntriesColumns holds the columns for the "credit_entries" table.
	CreditEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "namespace", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "entry_type", Type: field.TypeEnum, Enums: []string{"GRANT", "VOID_GRANT", "RESET"}},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"USAGE"}},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "priority", Type: field.TypeUint8, Default: 1},
		{Name: "effective_at", Type: field.TypeTime},
		{Name: "expiration_period_duration", Type: field.TypeEnum, Nullable: true, Enums: []string{"HOUR", "DAY", "WEEK", "MONTH", "YEAR"}},
		{Name: "expiration_period_count", Type: field.TypeUint8, Nullable: true},
		{Name: "rollover_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"ORIGINAL_AMOUNT", "REMAINING_AMOUNT"}},
		{Name: "rollover_max_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "parent_id", Type: field.TypeString, Unique: true, Nullable: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "product_id", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// CreditEntriesTable holds the schema information for the "credit_entries" table.
	CreditEntriesTable = &schema.Table{
		Name:       "credit_entries",
		Columns:    CreditEntriesColumns,
		PrimaryKey: []*schema.Column{CreditEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credit_entries_credit_entries_children",
				Columns:    []*schema.Column{CreditEntriesColumns[15]},
				RefColumns: []*schema.Column{CreditEntriesColumns[0]},
				OnDelete:   schema.Restrict,
			},
			{
				Symbol:     "credit_entries_products_credit_grants",
				Columns:    []*schema.Column{CreditEntriesColumns[16]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Restrict,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "creditentry_namespace_subject",
				Unique:  false,
				Columns: []*schema.Column{CreditEntriesColumns[3], CreditEntriesColumns[4]},
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "namespace", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "meter_slug", Type: field.TypeString},
		{Name: "meter_group_by_filters", Type: field.TypeJSON, Nullable: true},
		{Name: "archived", Type: field.TypeBool, Default: false},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "product_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{ProductsColumns[3], ProductsColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CreditEntriesTable,
		ProductsTable,
	}
)

func init() {
	CreditEntriesTable.ForeignKeys[0].RefTable = CreditEntriesTable
	CreditEntriesTable.ForeignKeys[1].RefTable = ProductsTable
}
