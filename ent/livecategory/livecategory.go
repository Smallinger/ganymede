// Code generated by ent, DO NOT EDIT.

package livecategory

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the livecategory type in the database.
	Label = "live_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeLive holds the string denoting the live edge name in mutations.
	EdgeLive = "live"
	// Table holds the table name of the livecategory in the database.
	Table = "live_categories"
	// LiveTable is the table that holds the live relation/edge.
	LiveTable = "live_categories"
	// LiveInverseTable is the table name for the Live entity.
	// It exists in this package in order to avoid circular dependency with the "live" package.
	LiveInverseTable = "lives"
	// LiveColumn is the table column denoting the live relation/edge.
	LiveColumn = "live_id"
)

// Columns holds all SQL columns for livecategory fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "live_categories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"live_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)