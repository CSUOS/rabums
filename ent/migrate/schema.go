// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClientServersColumns holds the columns for the "client_servers" table.
	ClientServersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "client_name", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "available", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ClientServersTable holds the schema information for the "client_servers" table.
	ClientServersTable = &schema.Table{
		Name:        "client_servers",
		Columns:     ClientServersColumns,
		PrimaryKey:  []*schema.Column{ClientServersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event", Type: field.TypeString},
		{Name: "message", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "event_user", Type: field.TypeInt, Nullable: true},
		{Name: "event_clientserver", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "events_users_user",
				Columns: []*schema.Column{EventsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "events_client_servers_clientserver",
				Columns: []*schema.Column{EventsColumns[5]},

				RefColumns: []*schema.Column{ClientServersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "user_pw", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
		{Name: "user_number", Type: field.TypeInt},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ClientServerOwnerColumns holds the columns for the "client_server_owner" table.
	ClientServerOwnerColumns = []*schema.Column{
		{Name: "client_server_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ClientServerOwnerTable holds the schema information for the "client_server_owner" table.
	ClientServerOwnerTable = &schema.Table{
		Name:       "client_server_owner",
		Columns:    ClientServerOwnerColumns,
		PrimaryKey: []*schema.Column{ClientServerOwnerColumns[0], ClientServerOwnerColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "client_server_owner_client_server_id",
				Columns: []*schema.Column{ClientServerOwnerColumns[0]},

				RefColumns: []*schema.Column{ClientServersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "client_server_owner_user_id",
				Columns: []*schema.Column{ClientServerOwnerColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClientServersTable,
		EventsTable,
		UsersTable,
		ClientServerOwnerTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = UsersTable
	EventsTable.ForeignKeys[1].RefTable = ClientServersTable
	ClientServerOwnerTable.ForeignKeys[0].RefTable = ClientServersTable
	ClientServerOwnerTable.ForeignKeys[1].RefTable = UsersTable
}
