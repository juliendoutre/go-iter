package iter

// Iterable describes a struct that can be iterated over.
type Iterable interface {
	Next() interface{}
}

// Iterator embeds an Iterable and provides util functions for it.
type Iterator struct {
	iter Iterable
}

// Next returns the next element of the Iterator.
func (i *Iterator) Next() interface{} {
	return i.iter.Next()
}

// Fold applies a reducer to the Iterator.
func (i *Iterator) Fold(init interface{}, reducer func(acc, item interface{}) interface{}) interface{} {
	acc := init

	item := i.Next()
	for item != nil {
		acc = reducer(acc, item)

		item = i.Next()
	}

	return acc
}

// TryFold folds over the Iterator and stops if it reaches the end of the Iterator or got a Break (bool, interface{}).
func (i *Iterator) TryFold(init interface{}, reducer func(acc, item interface{}) (interface{}, bool)) (interface{}, bool) {
	acc := init

	item := i.Next()
	for item != nil {
		r, ok := reducer(acc, item)
		if !ok {
			return r, ok
		}

		acc = r
		item = i.Next()
	}

	return acc, true
}

// FoldFirst folds over the Iterator, using its first element as the accumulator initial value.
func (i *Iterator) FoldFirst(reducer func(acc, item interface{}) interface{}) interface{} {
	first := i.Next()
	if first == nil {
		return nil
	}

	return i.Fold(first, reducer)
}

// Count returns the number of elements in the Iterator.
func (i *Iterator) Count() uint {
	return i.Fold(uint(0), func(acc, item interface{}) interface{} {
		return acc.(uint) + 1
	}).(uint)
}

// Last returns the last element of the Iterator.
func (i *Iterator) Last() interface{} {
	return i.Fold(nil, func(acc, item interface{}) interface{} {
		return item
	})
}

// AdvanceBy calls Next n times.
// It returns an error if it reached the end of the Iterator
// before it finished to iterate.
func (i *Iterator) AdvanceBy(n uint) error {
	for k := uint(0); k < n; k++ {
		if i.Next() == nil {
			return &errAdvanceBy{}
		}
	}

	return nil
}

type errAdvanceBy struct{}

func (e *errAdvanceBy) Error() string {
	return "`AdvanceBy` reached the end of the iterator"
}

// Nth returns the nth element of the Iterator.
func (i *Iterator) Nth(n uint) interface{} {
	i.AdvanceBy(n)
	return i.Next()
}

// ForEach runs a callback for every element of the iterator.
func (i *Iterator) ForEach(callback func(item interface{})) {
	i.Fold(nil, func(acc, item interface{}) interface{} {
		callback(item)
		return nil
	})
}

// Collect returns a slice containing the elements of the Iterator.
func (i *Iterator) Collect() []interface{} {
	collected := []interface{}{}

	item := i.Next()
	for item != nil {
		collected = append(collected, item)

		item = i.Next()
	}

	return collected
}

// All checks if all the elements of the Iterator validates a predicate.
func (i *Iterator) All(predicate func(item interface{}) bool) bool {
	_, ok := i.TryFold(nil, func(acc, item interface{}) (interface{}, bool) {
		return nil, predicate(item)
	})
	return ok
}

// Any checks if at least one element of the Iterator validates a predicate.
func (i *Iterator) Any(predicate func(item interface{}) bool) bool {
	_, ok := i.TryFold(nil, func(acc, item interface{}) (interface{}, bool) {
		return nil, !predicate(item)
	})
	return !ok
}

// Find returns the first element of the Iterator that validates a predicate.
func (i *Iterator) Find(predicate func(item interface{}) bool) interface{} {
	r, ok := i.TryFold(nil, func(acc, item interface{}) (interface{}, bool) {
		return item, !predicate(item)
	})

	if ok {
		return nil
	}

	return r
}

// Position returns the position of the first element of the Iterator that validates a predicate.
func (i *Iterator) Position(predicate func(item interface{}) bool) interface{} {
	r, ok := i.TryFold(uint(0), func(acc, item interface{}) (interface{}, bool) {
		if predicate(item) {
			return acc, false
		}

		return acc.(uint) + 1, true
	})

	if ok {
		return nil
	}

	return r
}

// SkipWhile skips the next elements until it reaches one which validates predicate.
func (i *Iterator) SkipWhile(predicate func(item interface{}) bool) *Iterator {
	i.Find(predicate)

	return i
}

// Skip the next n iterations.
func (i *Iterator) Skip(n uint) *Iterator {
	for k := uint(0); k < n && i.Next() != nil; k++ {
	}

	return i
}

// Map returns a new Iterator applying a mapper function to every element.
func (i *Iterator) Map(mapper func(item interface{}) interface{}) *Iterator {
	return &Iterator{iter: &mapIterable{mapper: mapper, iter: i}}
}

// Chain returns a new Iterator sequentially joining the two it was built on.
func (i *Iterator) Chain(iter *Iterator) *Iterator {
	return &Iterator{iter: &chain{first: i, second: iter, flag: false}}
}

// Filter returns a new Iterator yielding only elements validating a predicate.
func (i *Iterator) Filter(predicate func(item interface{}) bool) *Iterator {
	return &Iterator{iter: &filter{iter: i, predicate: predicate}}
}

// Enumerate returns a new Iterator yielding elements and their indices.
func (i *Iterator) Enumerate() *Iterator {
	return &Iterator{iter: &enumerate{iter: i}}
}

// Zip returns a new Iterator yielding pairs of elements of the two Iterators it was built from.
func (i *Iterator) Zip(iter *Iterator) *Iterator {
	return &Iterator{iter: &zip{first: i, second: iter}}
}

// TakeWhile returns a new Iterator yielding elements until predicate becomes false.
func (i *Iterator) TakeWhile(predicate func(item interface{}) bool) *Iterator {
	return &Iterator{iter: &takeWhile{iter: i, predicate: predicate, flag: false}}
}

// Take returns a new Iterator yielding only the n next elements.
func (i *Iterator) Take(n uint) *Iterator {
	return &Iterator{iter: &take{iter: i, count: 0, max: n, flag: false}}
}

var _ Iterable = &Iterator{}
