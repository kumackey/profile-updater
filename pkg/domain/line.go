package domain

type line struct {
	value      string
	isReplaced bool
}

func (l line) String() string {
	return l.value
}
