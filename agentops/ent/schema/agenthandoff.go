package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AgentHandoff represents a handoff of work from one agent to another.
// This tracks inter-agent communication and coordination.
type AgentHandoff struct {
	ent.Schema
}

// Annotations of the AgentHandoff.
func (AgentHandoff) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "agent_handoffs"},
	}
}

// Fields of the AgentHandoff.
func (AgentHandoff) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique handoff identifier (UUID)"),

		field.String("workflow_id").
			Optional().
			Comment("Parent workflow ID"),

		field.String("from_agent_id").
			NotEmpty().
			Comment("Agent initiating the handoff"),

		field.String("from_agent_type").
			Optional().
			Comment("Type of source agent"),

		field.String("to_agent_id").
			NotEmpty().
			Comment("Agent receiving the handoff"),

		field.String("to_agent_type").
			Optional().
			Comment("Type of target agent"),

		field.String("handoff_type").
			Default("request").
			Comment("Type of handoff: request, response, broadcast, delegate"),

		field.String("status").
			Default("pending").
			Comment("Handoff status: pending, accepted, rejected, completed, failed"),

		field.String("trace_id").
			Optional().
			Comment("OpenTelemetry trace ID"),

		field.String("from_task_id").
			Optional().
			Comment("Task ID in the source agent"),

		field.String("to_task_id").
			Optional().
			Comment("Task ID in the target agent"),

		field.JSON("payload", map[string]any{}).
			Optional().
			Comment("Data being passed between agents"),

		field.JSON("metadata", map[string]any{}).
			Optional().
			Comment("Additional metadata"),

		field.Int("payload_size_bytes").
			Default(0).
			Comment("Size of the payload in bytes"),

		field.Int64("latency_ms").
			Optional().
			Comment("Time from handoff initiation to acceptance"),

		field.String("error_message").
			Optional().
			Comment("Error message if handoff failed"),

		field.Time("initiated_at").
			Default(time.Now).
			Immutable().
			Comment("When the handoff was initiated"),

		field.Time("accepted_at").
			Optional().
			Nillable().
			Comment("When the handoff was accepted"),

		field.Time("completed_at").
			Optional().
			Nillable().
			Comment("When the handoff was completed"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the AgentHandoff.
func (AgentHandoff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workflow", Workflow.Type).
			Ref("handoffs").
			Field("workflow_id").
			Unique(),
	}
}

// Indexes of the AgentHandoff.
func (AgentHandoff) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("workflow_id"),
		index.Fields("from_agent_id"),
		index.Fields("to_agent_id"),
		index.Fields("handoff_type"),
		index.Fields("status"),
		index.Fields("trace_id"),
		index.Fields("initiated_at"),
		// Composite index for agent pair analysis
		index.Fields("from_agent_id", "to_agent_id"),
	}
}
