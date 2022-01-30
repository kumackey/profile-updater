package domain

type line struct {
	value         string
	shouldReplace bool
}

func (l line) String() string {
	return l.value
}
