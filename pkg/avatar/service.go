package avatar

/*import (
	"fmt"
	//"pkg/avatar/encoder"
	//"D:\Documentos\Taller Golang\Integrador\pkg\avatar\images"
)*/
import "fmt"

const (
	defaultWidth  = 200
	defaultLength = 200
	strider       = 10
)

// cryptoEncoder is someone who can encode information.
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

// imageGenerator is someone who can make images.
type imageGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}

// Generator contains functionalities related to avatar generation.
type Generator struct {
	encoder   cryptoEncoder
	Generator imageGenerator
}

func DefaultAvatarGeneration() *Generator {
	return &Generator{
		encoder:   encoder.NewMyEncoder(),
		Generator: images.NewDrawer(images.NewDefaultColorEngine(), defaultWidth, defaultLength, strider),
	}
}

// Information contains information (?)
type Information struct {
	// here goes all the information you want to encode
	Name string
}

func (genOne *Generator) GenerateAndSaveAvatar(information Information) error {
	// here will be all the logic
	encodedBytes, err := genOne.encoder.EncodeInformation(information.Name)
	if err != nil {
		return err
	}
	err = genOne.Generator.BuildAndSaveImage(encodedBytes)
	if err != nil {
		return err
	}
	fmt.Printf("the avatar related to %s name is created and saved", information.Name)
	return nil
}
