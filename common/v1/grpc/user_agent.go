//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//

package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

const (
	metadataKeyUA = "x-arango-graph-user-agent"
)

// UserAgentFromContext returns an user agent containing the name and version of the Tool used to call this func.
func UserAgentFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	vals := md.Get(metadataKeyUA)
	if len(vals) == 0 {
		return ""
	}
	// Return the first in the array
	return vals[0]
}

// CreateUserAgent creates a User agent string based on the provided information
func CreateUserAgent(applicationName, version string) string {
	return fmt.Sprintf("%s (%s)", applicationName, version)
}

// WithUserAgent sets grpc outgoing context to have given user agent
func WithUserAgent(ctx context.Context, userAgent string) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}
	md.Set(metadataKeyUA, userAgent)
	return metadata.NewOutgoingContext(ctx, md)
}

// WithUserAgentForTest sets grpc incoming context to have given user agent
// This can be used to mimic incoming traffic.
func WithUserAgentForTest(ctx context.Context, userAgent string) context.Context {
	md := metadata.New(nil)
	md.Set(metadataKeyUA, userAgent)
	return metadata.NewIncomingContext(ctx, md)
}
