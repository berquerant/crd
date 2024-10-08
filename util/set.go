package util

import "iter"

type Set[T comparable] map[T]bool

func NewSet[T comparable](values ...T) Set[T] {
	s := map[T]bool{}
	for _, v := range values {
		s[v] = true
	}
	return s
}

func (s Set[T]) In(v T) bool { return s[v] }
func (s Set[T]) Len() int    { return len(s) }
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}
