package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Treatment holds the schema definition for the Treatment entity.
type Treatment struct {
	ent.Schema
}

// Fields of the Treatment.
func (Treatment) Fields() []ent.Field {
	return []ent.Field{
		field.String("tablets"),
		field.String("psychologicalTreatment"),
		field.String("survey"),
		field.Int("patientNumber").Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Treatment.
func (Treatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("treat", Patient.Type).
			Ref("treats").
			Field("patientNumber").
			Unique(),
	}
}
