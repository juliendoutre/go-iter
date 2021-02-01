package iter

type mapIterable struct {
	iter   *Iterator
	mapper func(item interface{}) interface{}
}

func (m *mapIterable) Next() Option {
	item := m.iter.Next()
	if item.IsNone() {
		return None
	}

	return Some(m.mapper(item.Unwrap()))
}

var _ Iterable = &mapIterable{}

// Enumeration stores an Iterator element and its index.
type Enumeration struct {
	Index   uint
	Element interface{}
}

type enumerate struct {
	iter  *Iterator
	index uint
}

func (e *enumerate) Next() Option {
	item := e.iter.Next()
	if item.IsNone() {
		return None
	}

	o := Some(Enumeration{Index: e.index, Element: item.Unwrap()})
	e.index++
	return o
}

var _ Iterable = &enumerate{}

type chain struct {
	first  *Iterator
	second *Iterator
	flag   bool
}

func (c *chain) Next() Option {
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

var _ Iterable = &chain{}

type filter struct {
	iter      *Iterator
	predicate func(item interface{}) bool
}

func (f *filter) Next() Option {
	return f.iter.Find(f.predicate)
}

var _ Iterable = &filter{}

// Pair is a 2-tuple.
type Pair struct {
	First  interface{}
	Second interface{}
}

type zip struct {
	first  *Iterator
	second *Iterator
}

func (z *zip) Next() Option {
	f := z.first.Next()
	if f.IsNone() {
		return None
	}

	s := z.second.Next()
	if s.IsNone() {
		return None
	}

	return Some(Pair{First: f.Unwrap(), Second: s.Unwrap()})
}

var _ Iterable = &zip{}

type takeWhile struct {
	iter      *Iterator
	predicate func(item interface{}) bool
	flag      bool
}

func (t *takeWhile) Next() Option {
	if t.flag {
		return None
	}

	item := t.iter.Next()
	if item.IsNone() {
		t.flag = true
		return None
	}

	if !t.predicate(item.Unwrap()) {
		t.flag = true
		return None
	}

	return item
}

var _ Iterable = &takeWhile{}

type take struct {
	iter  *Iterator
	max   uint
	count uint
	flag  bool
}

func (t *take) Next() Option {
	if t.flag {
		return None
	}

	item := t.iter.Next()
	if item.IsNone() {
		t.flag = true
		return None
	}

	if t.count >= t.max {
		t.flag = true
		return None
	}

	t.count++

	return item
}

var _ Iterable = &take{}
