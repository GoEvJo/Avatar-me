package encoder

import (
	"crypto/sha512"

	"github.com/GoEvJo/Avatar-me/pkg/avatar/errorMessages"
)

type MyEncoder struct {
	strInformation string
}

func (e *MyEncoder) EncodeInformation() (encodedInformation []byte, err error) {
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

func NewMyEncoder(strInformation string) MyEncoder {
	NewEncoder := MyEncoder{
		strInformation: strInformation,
	}
	return NewEncoder
}
