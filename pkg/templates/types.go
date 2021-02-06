package templates

// Empty struct.
type Empty struct{}

type errAdvanceBy struct{}

func (e *errAdvanceBy) Error() string {
	return "`AdvanceBy` reached the end of the iterator"
}
