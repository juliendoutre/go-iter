package templates

type mapIterableForElement struct {
	iter   IterableForElement
	mapper func(item Element) Element
}

func (m *mapIterableForElement) Next() OptionForElement {
	item := m.iter.Next()
	if item.IsNone() {
		return NoneElement()
	}

	return SomeElement(m.mapper(item.Unwrap()))
}

var _ IterableForElement = &mapIterableForElement{}

type chainForElement struct {
	first  IterableForElement
	second IterableForElement
	flag   bool
}

func (c *chainForElement) Next() OptionForElement {
	if c.flag {
		return c.second.Next()
	}

	item := c.first.Next()
	if item.IsNone() {
		c.flag = true
		return c.second.Next()
	}

	return item
}

var _ IterableForElement = &chainForElement{}

// PairForElement is a 2-tuple.
type PairForElement struct {
	First  Element
	Second Element
}

type takeWhileForElement struct {
	iter      IterableForElement
	predicate func(item Element) bool
	flag      bool
}

func (t *takeWhileForElement) Next() OptionForElement {
	if t.flag {
		return NoneElement()
	}

	item := t.iter.Next()
	if item.IsNone() {
		t.flag = true
		return NoneElement()
	}

	if !t.predicate(item.Unwrap()) {
		t.flag = true
		return NoneElement()
	}

	return item
}

var _ IterableForElement = &takeWhileForElement{}

type takeForElement struct {
	iter  IterableForElement
	max   uint
	count uint
	flag  bool
}

func (t *takeForElement) Next() OptionForElement {
	if t.flag {
		return NoneElement()
	}

	item := t.iter.Next()
	if item.IsNone() {
		t.flag = true
		return NoneElement()
	}

	if t.count >= t.max {
		t.flag = true
		return NoneElement()
	}

	t.count++

	return item
}

var _ IterableForElement = &takeForElement{}

type filterForElement struct {
	iter      *IteratorForElement
	predicate func(item Element) bool
}

func (f *filterForElement) Next() OptionForElement {
	return f.iter.Find(f.predicate)
}

var _ IterableForElement = &filterForElement{}
