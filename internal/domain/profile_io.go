package domain

type ProfileIO interface {
	Scan() (*Profile, error)
	Write(*Profile) error
}
