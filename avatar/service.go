package avatar

import (
	"fmt"

	"github.com/GoEvJo/Avatar-me/avatar/encoder"
	"github.com/GoEvJo/Avatar-me/avatar/errorMessages"
	"github.com/GoEvJo/Avatar-me/avatar/images"
)

const length = 60

// cryptoEncoder is someone who can encode information.
type cryptoEncoder interface {
	EncodeInformation() (encodedInformation []byte, err error)
}

// imageGenerator is someone who can make and save images.
type imageGenerator interface {
	IdenticonGenerator(string2convert string, myHash []byte, length int) error
}

// Generator contains functionalities related to identicon generation, myEncoder wich allows to encode a string using Sha512
// and myIdenticoner wich creates and saves one identicon.
type Generator struct {
	myEncoder     cryptoEncoder
	myIdenticoner imageGenerator
}

// Information to be hashed, cannot be null.
var TheInfo string

// Function that gives me a Generator struct. Needs a not-null string, Sha512 info and positive length.
func DefaultFeaturesGeneration(strInformation string, hash []byte, length int) (*Generator, error) {
	if strInformation == "" {
		return nil, errorMessages.NullString
	}
	if len(hash) != 64 {
		return nil, errorMessages.Hashing
	}
	if length <= 0 {
		return nil, errorMessages.BadLength
	}
	theCryptoEncoder := encoder.NewMyEncoder(strInformation)
	myHash, err := theCryptoEncoder.EncodeInformation()
	if err != nil {
		return nil, err
	}
	theImageGenerator := images.Builder(myHash, length)
	return &Generator{
		myEncoder:     theCryptoEncoder,
		myIdenticoner: theImageGenerator,
	}, nil
}

// Generator method in charge of making and save the identicon.
func (genOne *Generator) GenerateAndSaveAvatar() error {
	// here will be all the logic

	encodedBytes, err := genOne.myEncoder.EncodeInformation()
	if err != nil {
		return err
	}
	err = genOne.myIdenticoner.IdenticonGenerator(TheInfo, encodedBytes, length)
	if err != nil {
		return err
	}
	fmt.Printf("the avatar related to %s name is created and saved", TheInfo)
	return nil
}
