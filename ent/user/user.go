// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUserPw holds the string denoting the user_pw field in the database.
	FieldUserPw = "user_pw"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldUserNumber holds the string denoting the user_number field in the database.
	FieldUserNumber = "user_number"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// EdgeOwns holds the string denoting the owns edge name in mutations.
	EdgeOwns = "owns"

	// Table holds the table name of the user in the database.
	Table = "users"
	// EventsTable is the table the holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "event_user"
	// OwnsTable is the table the holds the owns relation/edge. The primary key declared below.
	OwnsTable = "client_server_owner"
	// OwnsInverseTable is the table name for the ClientServer entity.
	// It exists in this package in order to avoid circular dependency with the "clientserver" package.
	OwnsInverseTable = "client_servers"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldUserPw,
	FieldUserName,
	FieldUserNumber,
	FieldEmail,
	FieldCreatedAt,
	FieldDeletedAt,
}

var (
	// OwnsPrimaryKey and OwnsColumn2 are the table columns denoting the
	// primary key for the owns relation (M2M).
	OwnsPrimaryKey = []string{"client_server_id", "user_id"}
)

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)