package iter

// Range builds an Iterator from a range of integers.
// The range start is inclusive but its end is exclusive.
func Range(start, end, step int) *Iterator {
	return &Iterator{
		iter: &rangeIterable{index: start, end: end, step: step},
	}
}

type rangeIterable struct {
	index int
	end   int
	step  int
}

func (r *rangeIterable) Next() interface{} {
	if (r.index-r.end)*r.step >= 0 {
		return nil
	}

	item := r.index
	r.index += r.step

	return item
}

var _ Iterable = &rangeIterable{}
