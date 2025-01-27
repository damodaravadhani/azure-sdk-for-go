//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProviderZeroValues(t *testing.T) {
	pr := Provider{}
	tr := pr.NewTracer("name", "version")
	require.Zero(t, tr)
	require.False(t, tr.Enabled())
	tr.SetAttributes()
	ctx, sp := tr.Start(context.Background(), "spanName", nil)
	require.Equal(t, context.Background(), ctx)
	require.Zero(t, sp)
	sp.AddError(nil)
	sp.AddEvent("event")
	sp.End()
	sp.SetAttributes(Attribute{})
	sp.SetStatus(SpanStatusError, "boom")
	sp, ok := tr.SpanFromContext(ctx)
	require.False(t, ok)
	require.Zero(t, sp)
}

func TestProvider(t *testing.T) {
	var addErrorCalled bool
	var addEventCalled bool
	var endCalled bool
	var setAttributesCalled bool
	var setStatusCalled bool
	var spanFromContextCalled bool

	pr := NewProvider(func(name, version string) Tracer {
		return NewTracer(func(context.Context, string, *SpanOptions) (context.Context, Span) {
			return nil, NewSpan(SpanImpl{
				AddError:      func(error) { addErrorCalled = true },
				AddEvent:      func(string, ...Attribute) { addEventCalled = true },
				End:           func() { endCalled = true },
				SetAttributes: func(...Attribute) { setAttributesCalled = true },
				SetStatus:     func(SpanStatus, string) { setStatusCalled = true },
			})
		}, &TracerOptions{
			SpanFromContext: func(context.Context) (Span, bool) {
				spanFromContextCalled = true
				return Span{}, true
			},
		})
	}, nil)
	tr := pr.NewTracer("name", "version")
	require.NotZero(t, tr)
	require.True(t, tr.Enabled())
	sp, ok := tr.SpanFromContext(context.Background())
	require.True(t, ok)
	require.Zero(t, sp)
	tr.SetAttributes(Attribute{Key: "some", Value: "attribute"})
	require.Len(t, tr.attrs, 1)
	require.EqualValues(t, tr.attrs[0].Key, "some")
	require.EqualValues(t, tr.attrs[0].Value, "attribute")

	ctx, sp := tr.Start(context.Background(), "name", nil)
	require.NotEqual(t, context.Background(), ctx)
	require.NotZero(t, sp)

	sp.AddError(nil)
	sp.AddEvent("event")
	sp.End()
	sp.SetAttributes()
	sp.SetStatus(SpanStatusError, "desc")
	require.True(t, addErrorCalled)
	require.True(t, addEventCalled)
	require.True(t, endCalled)
	require.True(t, setAttributesCalled)
	require.True(t, setStatusCalled)
	require.True(t, spanFromContextCalled)
}
