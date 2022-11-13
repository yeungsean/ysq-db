package internal

import (
	"context"

	"github.com/yeungsean/ysq-db/internal/provider"
)

// CtxKey ...
type CtxKey uint8

const (
	// CtxKeyLambda ...
	CtxKeyLambda CtxKey = iota
	// CtxKeySourceProvider ...
	CtxKeySourceProvider
	// CtxKeyFilterColumnIndex 过滤条件的列编号
	CtxKeyFilterColumnIndex
)

// CtxGetSourceProvider ...
func CtxGetSourceProvider(ctx context.Context) provider.IProvider {
	return ctx.Value(CtxKeySourceProvider).(provider.IProvider)
}

// CtxResetFilterColumnIndex ...
func CtxResetFilterColumnIndex(ctx context.Context) context.Context {
	idx := 1
	return context.WithValue(ctx, CtxKeyFilterColumnIndex, &idx)
}
