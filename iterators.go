package iter

type mapIterable struct {
	iter   *Iterator
	mapper func(item interface{}) interface{}
}

func (m *mapIterable) Next() interface{} {
	item := m.iter.Next()
	if item == nil {
		return nil
	}

	return m.mapper(item)
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

func (e *enumerate) Next() interface{} {
	item := e.iter.Next()
	if item == nil {
		return nil
	}

	o := Enumeration{Index: e.index, Element: item}
	e.index++
	return o
}

var _ Iterable = &enumerate{}

type chain struct {
	first  *Iterator
	second *Iterator
	flag   bool
}

func (c *chain) Next() interface{} {
	if c.flag {
		return c.second.Next()
	}

	item := c.first.Next()
	if item == nil {
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

func (f *filter) Next() interface{} {
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

func (z *zip) Next() interface{} {
	f := z.first.Next()
	if f == nil {
		return nil
	}

	s := z.second.Next()
	if s == nil {
		return nil
	}

	return Pair{First: f, Second: s}
}

var _ Iterable = &zip{}

type takeWhile struct {
	iter      *Iterator
	predicate func(item interface{}) bool
	flag      bool
}

func (t *takeWhile) Next() interface{} {
	if t.flag {
		return nil
	}

	item := t.iter.Next()
	if item == nil {
		t.flag = true
		return nil
	}

	if !t.predicate(item) {
		t.flag = true
		return nil
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

func (t *take) Next() interface{} {
	if t.flag {
		return nil
	}

	item := t.iter.Next()
	if item == nil {
		t.flag = true
		return nil
	}

	if t.count >= t.max {
		t.flag = true
		return nil
	}

	t.count++

	return item
}

var _ Iterable = &take{}
