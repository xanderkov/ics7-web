package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique(),
		field.String("password_hash"),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("is", Doctor.Type),
	}
}
