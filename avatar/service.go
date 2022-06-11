package avatar

import (
	"fmt"

	encoder "github.com/GoEvJo/Avatar-me/pkg/avatar/encoder"
	errorMessages "github.com/GoEvJo/Avatar-me/pkg/avatar/errorMessages"
	images "github.com/GoEvJo/Avatar-me/pkg/avatar/images"
)

const length = 60

// cryptoEncoder is someone who can encode information.
type cryptoEncoder interface {
	EncodeInformation() (encodedInformation []byte, err error)
}

// imageGenerator is someone who can make images.
type imageGenerator interface {
	IdenticonGenerator(string2convert string, myHash []byte, length int) error
}

// Generator contains functionalities related to avatar generation.
type Generator struct {
	myEncoder     cryptoEncoder
	myIdenticoner imageGenerator
}

// Information contains information (?)
var TheInfo string

func DefaultAvatarGeneration(strInformation string, hash []byte, length int) (*Generator, error) {
	if strInformation == "" {
		return nil, errorMessages.NullString
	}
	if len(hash) != 64 {
		return nil, errorMessages.Hashing
	}
	if length <= 0 {
		return nil, errorMessages.BadLength
	}
	return &Generator{
		myEncoder:     encoder.NewMyEncoder(strInformation),
		myIdenticoner: images.Builder(hash, length),
	}, nil
}

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
