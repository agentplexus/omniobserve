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

// Workflow represents an end-to-end workflow/session in a multi-agent system.
// A workflow contains multiple agent tasks and represents a complete unit of work.
type Workflow struct {
	ent.Schema
}

// Annotations of the Workflow.
func (Workflow) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "workflows"},
	}
}

// Fields of the Workflow.
func (Workflow) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique workflow identifier (UUID)"),

		field.String("name").
			NotEmpty().
			Comment("Workflow name/type (e.g., 'statistics-extraction')"),

		field.String("status").
			Default("running").
			Comment("Workflow status: running, completed, failed, cancelled"),

		field.String("trace_id").
			Optional().
			Comment("OpenTelemetry trace ID for correlation"),

		field.String("parent_workflow_id").
			Optional().
			Comment("Parent workflow ID for nested workflows"),

		field.String("initiator").
			Optional().
			Comment("What initiated the workflow (user_id, api_key, system)"),

		field.JSON("input", map[string]any{}).
			Optional().
			Comment("Workflow input data"),

		field.JSON("output", map[string]any{}).
			Optional().
			Comment("Workflow output/result data"),

		field.JSON("metadata", map[string]any{}).
			Optional().
			Comment("Additional metadata"),

		field.Int("task_count").
			Default(0).
			Comment("Number of tasks in this workflow"),

		field.Int("completed_task_count").
			Default(0).
			Comment("Number of completed tasks"),

		field.Int("failed_task_count").
			Default(0).
			Comment("Number of failed tasks"),

		field.Float("total_cost_usd").
			Default(0).
			Comment("Total cost in USD"),

		field.Int("total_tokens").
			Default(0).
			Comment("Total LLM tokens used"),

		field.Int64("duration_ms").
			Optional().
			Comment("Total duration in milliseconds"),

		field.String("error_message").
			Optional().
			Comment("Error message if workflow failed"),

		field.Time("started_at").
			Default(time.Now).
			Immutable().
			Comment("When the workflow started"),

		field.Time("ended_at").
			Optional().
			Nillable().
			Comment("When the workflow ended"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Workflow.
func (Workflow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", AgentTask.Type).
			Comment("Tasks belonging to this workflow"),
		edge.To("handoffs", AgentHandoff.Type).
			Comment("Agent handoffs in this workflow"),
		edge.To("events", AgentEvent.Type).
			Comment("Events in this workflow"),
	}
}

// Indexes of the Workflow.
func (Workflow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("status"),
		index.Fields("trace_id"),
		index.Fields("started_at"),
		index.Fields("initiator"),
	}
}
