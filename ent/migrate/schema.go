// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProvidersColumns holds the columns for the "providers" table.
	ProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 30},
		{Name: "path", Type: field.TypeString, Size: 100},
		{Name: "detour", Type: field.TypeString, Nullable: true, Size: 30},
		{Name: "remote", Type: field.TypeBool},
	}
	// ProvidersTable holds the schema information for the "providers" table.
	ProvidersTable = &schema.Table{
		Name:       "providers",
		Columns:    ProvidersColumns,
		PrimaryKey: []*schema.Column{ProvidersColumns[0]},
	}
	// RuleSetsColumns holds the columns for the "rule_sets" table.
	RuleSetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag", Type: field.TypeString, Unique: true, Size: 30},
		{Name: "type", Type: field.TypeString, Size: 10},
		{Name: "path", Type: field.TypeString, Size: 100},
		{Name: "format", Type: field.TypeString, Size: 10},
		{Name: "label", Type: field.TypeString, Size: 30},
		{Name: "download_detour", Type: field.TypeString, Size: 30},
		{Name: "update_interval", Type: field.TypeString, Size: 10},
		{Name: "outbound", Type: field.TypeString, Size: 30},
		{Name: "china", Type: field.TypeBool},
	}
	// RuleSetsTable holds the schema information for the "rule_sets" table.
	RuleSetsTable = &schema.Table{
		Name:       "rule_sets",
		Columns:    RuleSetsColumns,
		PrimaryKey: []*schema.Column{RuleSetsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProvidersTable,
		RuleSetsTable,
	}
)

func init() {
}
