package iter

// Range builds an Iterator from a range of integers.
// The range start is inclusive but its end is exclusive.
func Range(start, end, step int) IteratorForInt {
	return IteratorForInt{
		iter: &rangeIterable{index: start, end: end, step: step},
	}
}

type rangeIterable struct {
	index int
	end   int
	step  int
}

func (r *rangeIterable) Next() OptionForInt {
	if (r.index-r.end)*r.step >= 0 {
		return NoneInt()
	}

	item := r.index
	r.index += r.step

	return SomeInt(item)
}

var _ IterableForInt = &rangeIterable{}
