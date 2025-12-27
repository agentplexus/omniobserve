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

// AgentEvent represents a generic event in the agent system.
// This is an extensible event store for domain-specific and custom events.
type AgentEvent struct {
	ent.Schema
}

// Annotations of the AgentEvent.
func (AgentEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "agent_events"},
	}
}

// Fields of the AgentEvent.
func (AgentEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique event identifier (UUID)"),

		field.String("event_type").
			NotEmpty().
			Comment("Event type (e.g., 'agent.task.started', 'domain.statistic.extracted')"),

		field.String("event_category").
			Default("agent").
			Comment("Event category: agent, workflow, tool, domain, system"),

		field.String("workflow_id").
			Optional().
			Comment("Associated workflow ID"),

		field.String("task_id").
			Optional().
			Comment("Associated task ID"),

		field.String("agent_id").
			Optional().
			Comment("Associated agent ID"),

		field.String("trace_id").
			Optional().
			Comment("OpenTelemetry trace ID"),

		field.String("span_id").
			Optional().
			Comment("OpenTelemetry span ID"),

		field.String("severity").
			Default("info").
			Comment("Event severity: debug, info, warn, error"),

		field.JSON("data", map[string]any{}).
			Optional().
			Comment("Event payload data"),

		field.JSON("metadata", map[string]any{}).
			Optional().
			Comment("Additional metadata"),

		field.JSON("tags", []string{}).
			Optional().
			Comment("Event tags for filtering"),

		field.String("source").
			Optional().
			Comment("Source of the event (service name, component)"),

		field.String("version").
			Default("1.0").
			Comment("Event schema version"),

		field.Time("timestamp").
			Default(time.Now).
			Immutable().
			Comment("When the event occurred"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the AgentEvent.
func (AgentEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workflow", Workflow.Type).
			Ref("events").
			Field("workflow_id").
			Unique(),
		edge.From("task", AgentTask.Type).
			Ref("events").
			Field("task_id").
			Unique(),
	}
}

// Indexes of the AgentEvent.
func (AgentEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type"),
		index.Fields("event_category"),
		index.Fields("workflow_id"),
		index.Fields("task_id"),
		index.Fields("agent_id"),
		index.Fields("trace_id"),
		index.Fields("severity"),
		index.Fields("timestamp"),
		// Composite index for common queries
		index.Fields("event_type", "timestamp"),
		index.Fields("agent_id", "event_type"),
	}
}
