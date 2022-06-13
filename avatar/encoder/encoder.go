package encoder

import (
	"crypto/sha512"

	"github.com/GoEvJo/Avatar-me/avatar/errorMessages"
)

type myEncoder struct {
}

// Function that provides a pointer to a new myEncoder struct capable of using the interface.
func NewMyEncoder() myEncoder {
	NewEncoder := myEncoder{}
	return NewEncoder
}

// EncodeInformation is the function in charge of making the hash from a string. It returns hashed info (Sha512) and an error, if it occurs.
func (e *myEncoder) EncodeInformation(strInformation string) (encodedInformation []byte, err error) {

	if strInformation == "" {
		return nil, errorMessages.NullString
	}

	sha512 := sha512.New()

	sha512.Write([]byte(strInformation))
	returningValue := sha512.Sum(nil)
	if len(returningValue) != 64 {
		err = errorMessages.Hashing
		return nil, err
	}
	return returningValue, nil
}
