package pkg

import "context"

// CtxKey key定义
type CtxKey uint8

const (
	// CtxKeyLambda ...
	CtxKeyLambda CtxKey = iota
	// CtxKeyDBProvider 数据源provider
	CtxKeyDBProvider
	// CtxKeyFilterColumnIndex 过滤条件的列编号
	CtxKeyFilterColumnIndex
	// CtxKeyCacheProvider 缓存provider
	CtxKeyCacheProvider
	// CtxKeyTx 事务
	CtxKeyTx
	// CtxKeyDB db实例
	CtxKeyDB
)

// CtxResetFilterColumnIndex 获取过滤条件的列编号
func CtxResetFilterColumnIndex(ctx context.Context) context.Context {
	idx := 1
	return context.WithValue(ctx, CtxKeyFilterColumnIndex, &idx)
}
