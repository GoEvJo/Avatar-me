package encoder

import (
	"crypto/sha512"

	"github.com/GoEvJo/Avatar-me/avatar/errorMessages"
)

// myEncoder is a struct wich contains the information (not-null string) to be hashed.
type myEncoder struct {
	strInformation string
}

// EncodeInformation is the function in charge of making the hash from a string. It returns hashed info (Sha512) and an error, if it occurs.
func (e *myEncoder) EncodeInformation() (encodedInformation []byte, err error) {
	//init the hash
	sha512 := sha512.New()

	//pass the string
	sha512.Write([]byte(e.strInformation))

	//64bit hash code
	returningValue := sha512.Sum(nil)
	if len(returningValue) != 64 {
		err = errorMessages.Hashing
		return nil, err
	}
	return returningValue, nil
}

// Function that provides a pointer to a new myEncoder struct from a string.
func NewMyEncoder(strInformation string) *myEncoder {
	NewEncoder := myEncoder{
		strInformation: strInformation,
	}
	return &NewEncoder
}
