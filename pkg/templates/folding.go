package templates

import "github.com/cheekybits/genny/generic"

// Accumulator is the type used by accumulator values.
type Accumulator generic.Type

// FoldForAccumulator applies a reducer to the Iterator.
func (i IteratorForElement) FoldForAccumulator(init Accumulator, reducer func(acc Accumulator, item Element) Accumulator) Accumulator {
	acc := init

	item := i.Next()
	for item.IsSome() {
		acc = reducer(acc, item.Unwrap())

		item = i.Next()
	}

	return acc
}

// TryFoldForAccumulator folds over the Iterator and stops if it reaches the end of the Iterator or got a Break (bool, interface{}).
func (i IteratorForElement) TryFoldForAccumulator(init Accumulator, reducer func(acc Accumulator, item Element) (Accumulator, bool)) (Accumulator, bool) {
	acc := init

	item := i.Next()
	for item.IsSome() {
		r, ok := reducer(acc, item.Unwrap())
		if !ok {
			return r, ok
		}

		acc = r
		item = i.Next()
	}

	return acc, true
}
