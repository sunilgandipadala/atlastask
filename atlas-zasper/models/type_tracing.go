package models

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracing struct {
	Span trace.Span
	Ctx  context.Context
}

func (t *Tracing) Finish() {
	t.Span.End()
}

func (t *Tracing) GetContext() trace.SpanContext {
	return t.Span.SpanContext()
}
