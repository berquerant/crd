package util

import (
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
)

type Ring[T any] []T

func MustNewRing[T any](v ...T) Ring[T] {
	v, err := NewRing[T](v...)
	logx.PanicOnError(err)
	return v
}

// NewRing returns a new ring buffer.
func NewRing[T any](v ...T) (Ring[T], error) {
	if len(v) == 0 {
		return nil, errorx.Invalid("Ring should have 1 or more elements")
	}
	return Ring[T](v), nil
}

func (r Ring[T]) At(i int) T {
	if len(r) == 1 {
		return r[0]
	}

	index := i % len(r)
	if index < 0 {
		index += len(r)
	}
	return r[index]
}

func (r Ring[T]) Len() int { return len(r) }
