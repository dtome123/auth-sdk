package client

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	headerClientAssertionKey = "x-client-assertion"
	headerClientIDKey        = "x-client-id"
)

// WithClientAssertion adds the client JWT assertion to the outgoing context.
func WithClientAssertion(ctx context.Context, assertion string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, headerClientAssertionKey, assertion)
}

func WithClientID(ctx context.Context, clientID string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, headerClientIDKey, clientID)
}

// ClientAssertionFromContext extracts the client JWT assertion from the incoming context.
// Returns empty string if not found.
func ClientAssertionFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	values := md[headerClientAssertionKey]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func ClientIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	values := md[headerClientIDKey]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
