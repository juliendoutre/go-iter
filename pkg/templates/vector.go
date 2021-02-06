package templates

// VectorOfElement builds an Iterator from a slice.
func VectorOfElement(slice []Element) IteratorForElement {
	return IteratorForElement{
		iter: &vectorForElement{slice: slice, cursor: 0},
	}
}

type vectorForElement struct {
	slice  []Element
	cursor uint
}

func (v *vectorForElement) Next() OptionForElement {
	if v.cursor >= uint(len(v.slice)) {
		return NoneElement()
	}

	item := v.slice[v.cursor]
	v.cursor++

	return SomeElement(item)
}

var _ IterableForElement = &vectorForElement{}
