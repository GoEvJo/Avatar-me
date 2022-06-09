package encoder

import (
	"crypto/sha1"
)

type myEncoder struct{}

func (e *myEncoder) EncodeInformation(strInformation string) (encodedInformation []byte, err error) {
	strInformation = "sha1 this string"

	h := sha1.New()

	_, err = h.Write([]byte(strInformation))
	if err != nil {
		return nil, err
	}

	theHash := h.Sum(nil)

	return theHash, err
}
