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

// ToolInvocation represents a tool or function call made by an agent.
// This includes external API calls, database queries, file operations, etc.
type ToolInvocation struct {
	ent.Schema
}

// Annotations of the ToolInvocation.
func (ToolInvocation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tool_invocations"},
	}
}

// Fields of the ToolInvocation.
func (ToolInvocation) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique invocation identifier (UUID)"),

		field.String("task_id").
			Optional().
			Comment("Parent task ID"),

		field.String("agent_id").
			NotEmpty().
			Comment("Agent that invoked the tool"),

		field.String("tool_name").
			NotEmpty().
			Comment("Name of the tool (e.g., 'web_fetch', 'database_query', 'calculator')"),

		field.String("tool_type").
			Optional().
			Comment("Category of tool: http, database, file, function, external_api"),

		field.String("status").
			Default("running").
			Comment("Invocation status: running, success, failed, timeout"),

		field.String("trace_id").
			Optional().
			Comment("OpenTelemetry trace ID"),

		field.String("span_id").
			Optional().
			Comment("OpenTelemetry span ID"),

		field.JSON("input", map[string]any{}).
			Optional().
			Comment("Tool input/arguments"),

		field.JSON("output", map[string]any{}).
			Optional().
			Comment("Tool output/result"),

		field.JSON("metadata", map[string]any{}).
			Optional().
			Comment("Additional metadata"),

		// HTTP-specific fields
		field.String("http_method").
			Optional().
			Comment("HTTP method if applicable"),

		field.String("http_url").
			Optional().
			Comment("HTTP URL if applicable"),

		field.Int("http_status_code").
			Optional().
			Comment("HTTP status code if applicable"),

		// Performance metrics
		field.Int64("duration_ms").
			Optional().
			Comment("Invocation duration in milliseconds"),

		field.Int("request_size_bytes").
			Default(0).
			Comment("Size of the request in bytes"),

		field.Int("response_size_bytes").
			Default(0).
			Comment("Size of the response in bytes"),

		field.Int("retry_count").
			Default(0).
			Comment("Number of retries"),

		field.String("error_type").
			Optional().
			Comment("Error type if invocation failed"),

		field.String("error_message").
			Optional().
			Comment("Error message if invocation failed"),

		field.Time("started_at").
			Default(time.Now).
			Immutable().
			Comment("When the invocation started"),

		field.Time("ended_at").
			Optional().
			Nillable().
			Comment("When the invocation ended"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ToolInvocation.
func (ToolInvocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", AgentTask.Type).
			Ref("tool_invocations").
			Field("task_id").
			Unique(),
	}
}

// Indexes of the ToolInvocation.
func (ToolInvocation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id"),
		index.Fields("agent_id"),
		index.Fields("tool_name"),
		index.Fields("tool_type"),
		index.Fields("status"),
		index.Fields("trace_id"),
		index.Fields("started_at"),
		index.Fields("http_url"),
	}
}
