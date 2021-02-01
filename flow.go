package iter

// ControlFlow returns a value along a control flow indication.
type ControlFlow interface {
	ShouldContinue() bool
	ShouldBreak() bool
	Unwrap() interface{}
}

// Continue indicates a control flow should continue.
func Continue(value interface{}) ControlFlow {
	return &continueControlFlow{v: value}
}

type continueControlFlow struct {
	v interface{}
}

func (c *continueControlFlow) ShouldContinue() bool {
	return true
}

func (c *continueControlFlow) ShouldBreak() bool {
	return false
}

func (c *continueControlFlow) Unwrap() interface{} {
	return c.v
}

var _ ControlFlow = &continueControlFlow{}

// Break indicates a control flow should stop.
func Break(value interface{}) ControlFlow {
	return &breakControlFlow{v: value}
}

type breakControlFlow struct {
	v interface{}
}

func (b *breakControlFlow) ShouldContinue() bool {
	return false
}

func (b *breakControlFlow) ShouldBreak() bool {
	return true
}

func (b *breakControlFlow) Unwrap() interface{} {
	return b.v
}

var _ ControlFlow = &breakControlFlow{}
