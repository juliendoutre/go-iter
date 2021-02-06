package templates

// OptionForElement can hold an Element value or not.
type OptionForElement struct {
	value  Element
	isNone bool
}

// SomeElement returns an Option holding an Element value.
func SomeElement(value Element) OptionForElement {
	return OptionForElement{value: value, isNone: false}
}

// NoneElement returns an Option holding no Element value.
func NoneElement() OptionForElement {
	return OptionForElement{isNone: true}
}

func (o OptionForElement) IsSome() bool {
	return !o.isNone
}

func (o OptionForElement) IsNone() bool {
	return o.isNone
}

func (o OptionForElement) Expect(msg string) Element {
	if o.isNone {
		panic(msg)
	}

	return o.value
}

func (o OptionForElement) ExpectNone(msg string) {
	if !o.isNone {
		panic(msg)
	}
}

func (o OptionForElement) Unwrap() Element {
	return o.Expect("Called `Unwrap` on a `None` Option.")
}

func (o OptionForElement) UnwrapOr(defaultValue Element) Element {
	if o.isNone {
		return defaultValue
	}

	return o.value
}

func (o OptionForElement) UnwrapOrElse(f func() Element) Element {
	if o.isNone {
		return f()
	}

	return o.value
}

func (o OptionForElement) UnwrapNone() {
	o.ExpectNone("Called `UnwrapNone` on a `Some` value")
}
