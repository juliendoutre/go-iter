package iter

// Option can hold a value or not.
type Option interface {
	IsSome() bool
	IsNone() bool
	Expect(msg string) interface{}
	ExpectNone(msg string)
	Unwrap() interface{}
	UnwrapOr(defaultValue interface{}) interface{}
	UnwrapOrElse(f func() interface{}) interface{}
	UnwrapNone()
}

// Some returns an Option holding a value.
func Some(value interface{}) Option {
	return &some{value: value}
}

type some struct {
	value interface{}
}

func (s *some) IsSome() bool {
	return true
}

func (s *some) IsNone() bool {
	return false
}

func (s *some) Expect(msg string) interface{} {
	return s.value
}

func (s *some) ExpectNone(msg string) {
	panic(msg)
}

func (s *some) Unwrap() interface{} {
	return s.value
}

func (s *some) UnwrapOr(defaultValue interface{}) interface{} {
	return s.value
}

func (s *some) UnwrapOrElse(f func() interface{}) interface{} {
	return s.value
}

func (s *some) UnwrapNone() {
	panic("Called `UnwrapNone` on a `Some` value")
}

var _ Option = &some{}

// None is an empty Option.
var None Option = &none{}

type none struct{}

func (n *none) IsSome() bool {
	return false
}

func (n *none) IsNone() bool {
	return true
}

func (n *none) Expect(msg string) interface{} {
	panic(msg)
}

func (n *none) ExpectNone(msg string) {}

func (n *none) Unwrap() interface{} {
	panic("Called `Unwrap` on a `None` value")
}

func (n *none) UnwrapOr(defaultValue interface{}) interface{} {
	return defaultValue
}

func (n *none) UnwrapOrElse(f func() interface{}) interface{} {
	return f()
}

func (n *none) UnwrapNone() {}
