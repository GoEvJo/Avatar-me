package avatar

import (
	"fmt"

	"github.com/GoEvJo/Avatar-me/avatar/encoder"
	"github.com/GoEvJo/Avatar-me/avatar/errorMessages"
	"github.com/GoEvJo/Avatar-me/avatar/images"
)

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

type imageGenerator interface {
	IdenticonGenerator(string2convert string, myHash []byte) error
}

// Generator contains functionalities related to identicon generation: myEncoder wich allows to encode a string using Sha512
// and myIdenticoner wich creates and saves one identicon.
type Generator struct {
	myEncoder     cryptoEncoder
	myIdenticoner imageGenerator
}

// Information to be hashed, cannot be null.
var TheInfo string

// Function that gives me a Generator struct.
func DefaultFeaturesGeneration() (*Generator, error) {

	if TheInfo == "" {
		return nil, errorMessages.NullString
	}

	theCryptoEncoder := encoder.NewMyEncoder()
	theImageGenerator := images.Builder()
	NewGenerator := Generator{
		myEncoder:     &theCryptoEncoder,
		myIdenticoner: &theImageGenerator,
	}

	return &NewGenerator, nil
}

// Generator method in charge of making and save the identicon.
func (genOne *Generator) GenerateAndSaveAvatar() error {

	encodedBytes, err := genOne.myEncoder.EncodeInformation(TheInfo)
	if err != nil {
		return err
	}
	err = genOne.myIdenticoner.IdenticonGenerator(TheInfo, encodedBytes)
	if err != nil {
		return err
	}
	fmt.Printf("The identicon related to %s name was created and saved", TheInfo)
	return nil
}
