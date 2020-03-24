//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
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
