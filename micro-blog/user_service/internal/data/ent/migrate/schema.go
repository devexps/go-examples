// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 50, Comment: "username"},
		{Name: "nickname", Type: field.TypeString, Size: 128, Comment: "nickname"},
		{Name: "password", Type: field.TypeString, Size: 255, Comment: "password"},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UserColumns[0]},
			},
			{
				Name:    "user_id_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0], UserColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserTable,
	}
)

func init() {
	UserTable.Annotation = &entsql.Annotation{
		Table:     "user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}