// Code generated by entc, DO NOT EDIT.

package playback

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zibbp/ganymede/internal/utils"
)

const (
	// Label holds the string label denoting the playback type in the database.
	Label = "playback"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVodID holds the string denoting the vod_id field in the database.
	FieldVodID = "vod_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the playback in the database.
	Table = "playbacks"
)

// Columns holds all SQL columns for playback fields.
var Columns = []string{
	FieldID,
	FieldVodID,
	FieldUserID,
	FieldTime,
	FieldStatus,
	FieldUpdatedAt,
	FieldCreatedAt,
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
	// DefaultTime holds the default value on creation for the "time" field.
	DefaultTime int
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

const DefaultStatus utils.PlaybackStatus = "in_progress"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s utils.PlaybackStatus) error {
	switch s {
	case "in_progress", "finished":
		return nil
	default:
		return fmt.Errorf("playback: invalid enum value for status field: %q", s)
	}
}