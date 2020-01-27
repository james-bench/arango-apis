//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	bearerPrefix = "bearer "
)

// GetAccessToken fetches the access token from the given (GRPC) context.
// If no such token is found, false is returned.
func GetAccessToken(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	auth := md.Get("Authorization")
	if len(auth) == 0 {
		return "", false
	}
	for _, token := range auth {
		if strings.HasPrefix(strings.ToLower(token), bearerPrefix) {
			result := strings.TrimSpace(token[len(bearerPrefix):])
			return result, true
		}
	}
	return "", false
}

// WithAccessToken returns a context that containes the given access token
// for outgoing (GRPC) calls and is derived from the given context.WithAccessToken
func WithAccessToken(ctx context.Context, token string) context.Context {
	if token == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, "Authorization", bearerPrefix+token)
}
