package schema

import "entgo.io/ent"

// ToDo holds the schema definition for the ToDo entity.
type ToDo struct {
	ent.Schema
}

// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return nil
}

// Edges of the ToDo.
func (ToDo) Edges() []ent.Edge {
	return nil
}
