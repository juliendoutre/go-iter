package iter

// Vector builds an Iterator from a slice.
func Vector(slice []interface{}) *Iterator {
	return &Iterator{
		iter: &vector{slice: slice, cursor: 0},
	}
}

type vector struct {
	slice  []interface{}
	cursor uint
}

func (v *vector) Next() interface{} {
	if v.cursor >= uint(len(v.slice)) {
		return nil
	}

	item := v.slice[v.cursor]
	v.cursor++

	return item
}

var _ Iterable = &vector{}
