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

// AgentTask represents a task executed by an agent.
// Each task is part of a workflow and may involve LLM calls, tool invocations, etc.
type AgentTask struct {
	ent.Schema
}

// Annotations of the AgentTask.
func (AgentTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "agent_tasks"},
	}
}

// Fields of the AgentTask.
func (AgentTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique task identifier (UUID)"),

		field.String("workflow_id").
			Optional().
			Comment("Parent workflow ID"),

		field.String("agent_id").
			NotEmpty().
			Comment("Agent that executed this task"),

		field.String("agent_type").
			Optional().
			Comment("Type of agent (e.g., 'synthesis', 'verification', 'research')"),

		field.String("task_type").
			NotEmpty().
			Comment("Type of task (e.g., 'extract_statistics', 'verify_source')"),

		field.String("name").
			NotEmpty().
			Comment("Human-readable task name"),

		field.String("status").
			Default("running").
			Comment("Task status: pending, running, completed, failed, cancelled"),

		field.String("trace_id").
			Optional().
			Comment("OpenTelemetry trace ID"),

		field.String("span_id").
			Optional().
			Comment("OpenTelemetry span ID"),

		field.String("parent_span_id").
			Optional().
			Comment("Parent span ID for nested tasks"),

		field.JSON("input", map[string]any{}).
			Optional().
			Comment("Task input data"),

		field.JSON("output", map[string]any{}).
			Optional().
			Comment("Task output/result data"),

		field.JSON("metadata", map[string]any{}).
			Optional().
			Comment("Additional metadata"),

		field.Int("llm_call_count").
			Default(0).
			Comment("Number of LLM calls made"),

		field.Int("tool_call_count").
			Default(0).
			Comment("Number of tool invocations"),

		field.Int("retry_count").
			Default(0).
			Comment("Number of retries attempted"),

		field.Int("tokens_prompt").
			Default(0).
			Comment("Prompt tokens used"),

		field.Int("tokens_completion").
			Default(0).
			Comment("Completion tokens used"),

		field.Int("tokens_total").
			Default(0).
			Comment("Total tokens used"),

		field.Float("cost_usd").
			Default(0).
			Comment("Cost in USD"),

		field.Int64("duration_ms").
			Optional().
			Comment("Task duration in milliseconds"),

		field.String("error_type").
			Optional().
			Comment("Error type if task failed"),

		field.String("error_message").
			Optional().
			Comment("Error message if task failed"),

		field.Time("started_at").
			Default(time.Now).
			Immutable().
			Comment("When the task started"),

		field.Time("ended_at").
			Optional().
			Nillable().
			Comment("When the task ended"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the AgentTask.
func (AgentTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workflow", Workflow.Type).
			Ref("tasks").
			Field("workflow_id").
			Unique(),
		edge.To("tool_invocations", ToolInvocation.Type).
			Comment("Tool invocations made during this task"),
		edge.To("events", AgentEvent.Type).
			Comment("Events during this task"),
	}
}

// Indexes of the AgentTask.
func (AgentTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("workflow_id"),
		index.Fields("agent_id"),
		index.Fields("agent_type"),
		index.Fields("task_type"),
		index.Fields("status"),
		index.Fields("trace_id"),
		index.Fields("started_at"),
	}
}
