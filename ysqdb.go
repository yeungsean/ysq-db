package ysqdb

type (
	// Iterable ...
	Iterable[T any] interface {
		Next() Iterator[T]
	}

	// Iterator 迭代器
	Iterator[T any] func() T
)
