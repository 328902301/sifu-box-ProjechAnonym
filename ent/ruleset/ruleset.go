// Code generated by ent, DO NOT EDIT.

package ruleset

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the ruleset type in the database.
	Label = "rule_set"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldFormat holds the string denoting the format field in the database.
	FieldFormat = "format"
	// FieldLabel holds the string denoting the label field in the database.
	FieldLabel = "label"
	// FieldDownloadDetour holds the string denoting the download_detour field in the database.
	FieldDownloadDetour = "download_detour"
	// FieldUpdateInterval holds the string denoting the update_interval field in the database.
	FieldUpdateInterval = "update_interval"
	// FieldNameServer holds the string denoting the name_server field in the database.
	FieldNameServer = "name_server"
	// FieldChina holds the string denoting the china field in the database.
	FieldChina = "china"
	// Table holds the table name of the ruleset in the database.
	Table = "rule_sets"
)

// Columns holds all SQL columns for ruleset fields.
var Columns = []string{
	FieldID,
	FieldTag,
	FieldType,
	FieldPath,
	FieldFormat,
	FieldLabel,
	FieldDownloadDetour,
	FieldUpdateInterval,
	FieldNameServer,
	FieldChina,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// FormatValidator is a validator for the "format" field. It is called by the builders before save.
	FormatValidator func(string) error
	// LabelValidator is a validator for the "label" field. It is called by the builders before save.
	LabelValidator func(string) error
	// DownloadDetourValidator is a validator for the "download_detour" field. It is called by the builders before save.
	DownloadDetourValidator func(string) error
	// UpdateIntervalValidator is a validator for the "update_interval" field. It is called by the builders before save.
	UpdateIntervalValidator func(string) error
	// NameServerValidator is a validator for the "name_server" field. It is called by the builders before save.
	NameServerValidator func(string) error
)

// OrderOption defines the ordering options for the RuleSet queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTag orders the results by the tag field.
func ByTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTag, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByFormat orders the results by the format field.
func ByFormat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormat, opts...).ToFunc()
}

// ByLabel orders the results by the label field.
func ByLabel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLabel, opts...).ToFunc()
}

// ByDownloadDetour orders the results by the download_detour field.
func ByDownloadDetour(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadDetour, opts...).ToFunc()
}

// ByUpdateInterval orders the results by the update_interval field.
func ByUpdateInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateInterval, opts...).ToFunc()
}

// ByNameServer orders the results by the name_server field.
func ByNameServer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameServer, opts...).ToFunc()
}

// ByChina orders the results by the china field.
func ByChina(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChina, opts...).ToFunc()
}
