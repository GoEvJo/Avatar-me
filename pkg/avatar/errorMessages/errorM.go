package errorMessages

import (
	"errors"
)

var (
	NoFile     = errors.New("error with the file")
	Hashing    = errors.New("incorrect hash value")
	BadLength  = errors.New("non-positive length")
	NullString = errors.New("empty string")
)
