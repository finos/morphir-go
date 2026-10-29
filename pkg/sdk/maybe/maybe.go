// Package maybe provides a Go implementation of the Morphir.SDK.Maybe module.
//
// A Maybe represents an optional value: either Just(value) or Nothing.
// It is an immutable value type; every operation returns a new Maybe and the
// zero value of Maybe[T] is Nothing.
//
// The API mirrors Morphir.SDK.Maybe (see Morphir.IR.SDK.Maybe in morphir-elm
// and morphir.sdk.Maybe in morphir-jvm): Just, Nothing, Map, Map2..Map5,
// AndThen, WithDefault and HasValue.
package maybe

// Maybe is an optional value of type T.
//
// Maybe[T] is comparable whenever T is comparable. Its zero value is Nothing.
type Maybe[T any] struct {
	value T
	valid bool
}

// Just wraps a value in a Maybe that has a value.
// The wrapped value may itself be the zero value of T; Just(0) is not Nothing.
func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value: value, valid: true}
}

// Nothing returns a Maybe that has no value.
func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

// Value returns the wrapped value and true if m is Just, or the zero value of T
// and false if m is Nothing.
func (m Maybe[T]) Value() (T, bool) {
	return m.value, m.valid
}

// HasValue reports whether m is a Just.
func HasValue[T any](m Maybe[T]) bool {
	return m.valid
}

// WithDefault provides a default value, turning an optional value into a
// normal value. It returns the wrapped value if m is Just, otherwise def.
func WithDefault[T any](def T, m Maybe[T]) T {
	if m.valid {
		return m.value
	}
	return def
}

// Map transforms the value inside m with f. If m is Nothing, f is not called
// and Nothing is returned.
func Map[A, B any](f func(A) B, m Maybe[A]) Maybe[B] {
	if !m.valid {
		return Nothing[B]()
	}
	return Just(f(m.value))
}

// Map2 applies f if both arguments are Just. If either is Nothing, f is not
// called and Nothing is returned.
func Map2[A, B, R any](f func(A, B) R, m1 Maybe[A], m2 Maybe[B]) Maybe[R] {
	if !m1.valid || !m2.valid {
		return Nothing[R]()
	}
	return Just(f(m1.value, m2.value))
}

// Map3 applies f if all three arguments are Just. If any is Nothing, f is not
// called and Nothing is returned.
func Map3[A, B, C, R any](f func(A, B, C) R, m1 Maybe[A], m2 Maybe[B], m3 Maybe[C]) Maybe[R] {
	if !m1.valid || !m2.valid || !m3.valid {
		return Nothing[R]()
	}
	return Just(f(m1.value, m2.value, m3.value))
}

// Map4 applies f if all four arguments are Just. If any is Nothing, f is not
// called and Nothing is returned.
func Map4[A, B, C, D, R any](f func(A, B, C, D) R, m1 Maybe[A], m2 Maybe[B], m3 Maybe[C], m4 Maybe[D]) Maybe[R] {
	if !m1.valid || !m2.valid || !m3.valid || !m4.valid {
		return Nothing[R]()
	}
	return Just(f(m1.value, m2.value, m3.value, m4.value))
}

// Map5 applies f if all five arguments are Just. If any is Nothing, f is not
// called and Nothing is returned.
func Map5[A, B, C, D, E, R any](f func(A, B, C, D, E) R, m1 Maybe[A], m2 Maybe[B], m3 Maybe[C], m4 Maybe[D], m5 Maybe[E]) Maybe[R] {
	if !m1.valid || !m2.valid || !m3.valid || !m4.valid || !m5.valid {
		return Nothing[R]()
	}
	return Just(f(m1.value, m2.value, m3.value, m4.value, m5.value))
}

// AndThen chains together computations that may fail. If m is Just, f is
// applied to the wrapped value and its result is returned; otherwise f is not
// called and Nothing is returned.
func AndThen[A, B any](f func(A) Maybe[B], m Maybe[A]) Maybe[B] {
	if !m.valid {
		return Nothing[B]()
	}
	return f(m.value)
}
