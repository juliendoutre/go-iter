package templates

import "github.com/cheekybits/genny/generic"

// Element is the type of the elements in Iterators.
type Element generic.Type

// IterableForElement describes a struct that can be iterated over.
type IterableForElement interface {
	Next() OptionForElement
}

// IteratorForElement embeds an Iterable and provides util functions for it.
type IteratorForElement struct {
	iter IterableForElement
}

// Iterator implements Iterable.
var _ IterableForElement = IteratorForElement{}

// Next returns the next element of the Iterator.
func (i IteratorForElement) Next() OptionForElement {
	return i.iter.Next()
}

// AdvanceBy calls Next n times.
// It returns an error if it reached the end of the Iterator
// before it finished to iterate.
func (i IteratorForElement) AdvanceBy(n uint) error {
	for k := uint(0); k < n; k++ {
		if i.Next().IsNone() {
			return &errAdvanceBy{}
		}
	}

	return nil
}

// Nth returns the nth element of the Iterator.
func (i IteratorForElement) Nth(n uint) OptionForElement {
	i.AdvanceBy(n)
	return i.Next()
}

// Skip the next n iterations.
func (i IteratorForElement) Skip(n uint) IteratorForElement {
	i.AdvanceBy(n)
	return i
}

// Collect returns a slice containing the elements of the Iterator.
func (i IteratorForElement) Collect() []Element {
	collected := []Element{}

	item := i.Next()
	for item.IsSome() {
		collected = append(collected, item.Unwrap())

		item = i.Next()
	}

	return collected
}

// FoldFirst folds over the Iterator, using its first element as the accumulator initial value.
func (i IteratorForElement) FoldFirst(reducer func(acc, item Element) Element) OptionForElement {
	first := i.Next()
	if first.IsNone() {
		return NoneElement()
	}

	return SomeElement(i.FoldForElement(first.Unwrap(), reducer))
}

// Count returns the number of elements in the Iterator.
func (i IteratorForElement) Count() uint {
	return i.FoldForUint(uint(0), func(acc uint, item Element) uint {
		return acc + 1
	})
}

// Last returns the last element of the Iterator.
func (i IteratorForElement) Last() OptionForElement {
	return i.FoldForOptionForElement(NoneElement(), func(acc OptionForElement, item Element) OptionForElement {
		return SomeElement(item)
	})
}

// ForEach runs a callback for every element of the iterator.
func (i IteratorForElement) ForEach(callback func(item Element)) {
	i.FoldForEmpty(Empty{}, func(acc Empty, item Element) Empty {
		callback(item)
		return acc
	})
}

// All checks if all the elements of the Iterator validates a predicate.
func (i IteratorForElement) All(predicate func(item Element) bool) bool {
	_, ok := i.TryFoldForEmpty(Empty{}, func(acc Empty, item Element) (Empty, bool) {
		return acc, predicate(item)
	})
	return ok
}

// Any checks if at least one element of the Iterator validates a predicate.
func (i IteratorForElement) Any(predicate func(item Element) bool) bool {
	_, ok := i.TryFoldForEmpty(Empty{}, func(acc Empty, item Element) (Empty, bool) {
		return acc, !predicate(item)
	})
	return !ok
}

// Find returns the first element of the Iterator that validates a predicate.
func (i IteratorForElement) Find(predicate func(item Element) bool) OptionForElement {
	r, ok := i.TryFoldForOptionForElement(NoneElement(), func(acc OptionForElement, item Element) (OptionForElement, bool) {
		return SomeElement(item), !predicate(item)
	})

	if ok {
		return NoneElement()
	}

	return r
}

// Position returns the position of the first element of the Iterator that validates a predicate.
func (i IteratorForElement) Position(predicate func(item Element) bool) OptionForUint {
	r, ok := i.TryFoldForUint(uint(0), func(acc uint, item Element) (uint, bool) {
		if predicate(item) {
			return acc, false
		}

		return acc + 1, true
	})

	if ok {
		return NoneUint()
	}

	return SomeUint(r)
}

// SkipWhile skips the next elements until it reaches one which validates predicate.
func (i IteratorForElement) SkipWhile(predicate func(item Element) bool) IteratorForElement {
	i.Find(predicate)

	return i
}

// Map returns a new Iterator applying a mapper function to every element.
func (i IteratorForElement) Map(mapper func(item Element) Element) IteratorForElement {
	return IteratorForElement{iter: &mapIterableForElement{mapper: mapper, iter: i.iter}}
}

// Chain returns a new Iterator sequentially joining the two it was built on.
func (i IteratorForElement) Chain(iter IteratorForElement) IteratorForElement {
	return IteratorForElement{iter: &chainForElement{first: i.iter, second: iter.iter, flag: false}}
}

// TakeWhile returns a new Iterator yielding elements until predicate becomes false.
func (i IteratorForElement) TakeWhile(predicate func(item Element) bool) IteratorForElement {
	return IteratorForElement{iter: &takeWhileForElement{iter: i.iter, predicate: predicate, flag: false}}
}

// Take returns a new Iterator yielding only the n next elements.
func (i IteratorForElement) Take(n uint) IteratorForElement {
	return IteratorForElement{iter: &takeForElement{iter: i.iter, count: 0, max: n, flag: false}}
}

// Filter returns a new Iterator yielding only elements validating a predicate.
func (i IteratorForElement) Filter(predicate func(item Element) bool) IteratorForElement {
	return IteratorForElement{iter: &filterForElement{iter: i, predicate: predicate}}
}
