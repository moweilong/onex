package middleware

import (
	"context"

	"github.com/moweilong/onex/pkg/authz/engine"
)

type Option func(*options)

type options struct{}

func NewContext(ctx context.Context, claims *engine.AuthClaims) context.Context {
	return engine.ContextWithAuthClaims(ctx, claims)
}

func FromContext(ctx context.Context) (*engine.AuthClaims, bool) {
	return engine.AuthClaimsFromContext(ctx)
}
