// Package journey provides OpenTelemetry Semantic Conventions for User Journey observability.
//
// These attribute names extend OpenTelemetry's gen_ai.* namespace with user journey-specific
// concepts for product analytics and user experience observability. They are designed to
// complement the gen_ai.agent.* conventions for AI systems.
//
// # Overview
//
// User journeys represent the paths users take through an application to achieve goals.
// ProductGraph uses these conventions to:
//
//   - Track user sessions and page navigation
//   - Capture UI interactions (clicks, inputs, scrolls)
//   - Record state changes in frontend applications
//   - Match user behavior to predefined journey definitions
//   - Calculate conversion funnels and cohort retention
//
// # Namespace Structure
//
// The conventions are organized into several namespaces:
//
//   - gen_ai.journey.*     - Journey definitions and steps
//   - session.*            - User session attributes
//   - ui.*                 - UI interaction attributes
//   - snapshot.*           - Visual capture attributes
//   - page.*               - Page navigation attributes
//
// # Integration with AgentOps
//
// These conventions work alongside the agent package (gen_ai.agent.*) to provide
// end-to-end observability:
//
//	Frontend (journey)                    Backend (agent)
//	├── session.id ─────────────────────► trace_id correlation
//	├── ui.click AddToCart ────────────► gen_ai.agent.task (process order)
//	│   └── api.request POST /cart        └── gen_ai.agent.tool_call (inventory check)
//	└── page.view /confirmation            └── gen_ai.agent.event (order created)
//
// # Example Usage
//
// Events emitted from frontend applications include these attributes:
//
//	{
//	    "session.id": "sess-abc123",
//	    "event.type": "ui.click",
//	    "ui.component.name": "AddToCartButton",
//	    "ui.component.path": "App/ProductPage/ProductDetails/AddToCartButton",
//	    "ui.action": "click",
//	    "gen_ai.journey.id": "checkout-flow",
//	    "gen_ai.journey.step.id": "add-to-cart",
//	    "page.path": "/products/widget-123",
//	    "timestamp": "2025-01-15T14:32:01Z"
//	}
//
// # ProductGraph Compatibility
//
// These conventions are designed specifically for ProductGraph's event ingestion API
// but follow OpenTelemetry patterns for interoperability with OTLP collectors and
// standard observability backends.
package journey
