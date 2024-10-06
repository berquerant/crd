package util

type Opt[T any] struct {
	value   T
	updated bool
}

func NewOpt[T any](v T) *Opt[T] {
	return &Opt[T]{
		value:   v,
		updated: true,
	}
}

func (p Opt[T]) Unwrap() T {
	return p.value
}

func (p *Opt[T]) Update(v T) {
	p.value = v
	p.updated = true
}

// WhenUpdated calls f the first time this is called after being Update() or NewOpt().
func (p *Opt[T]) WhenUpdated(f func(v T)) {
	if p.updated {
		p.updated = false
		f(p.value)
	}
}
