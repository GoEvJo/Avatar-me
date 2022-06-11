package images

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/GoEvJo/Avatar-me/pkg/avatar/errorMessages"
)

type generatorStaff struct {
	hash   []byte
	length int
}

func Builder(hash []byte, length int) (generatorStaff, error) {
	if length <= 0 {
		return generatorStaff{
			hash:   nil,
			length: 0,
		}, errorMessages.BadLength
	}
	if len(hash) != 64 {
		return generatorStaff{
			hash:   nil,
			length: 0,
		}, errorMessages.Hashing
	}
	return generatorStaff{
		hash:   hash,
		length: length,
	}, nil
}

func (meth *generatorStaff) IdenticonGenerator(string2convert string, myHash []byte, length int) error {
	if string2convert == "" {
		return errorMessages.NullString
	}
	if length <= 0 {
		return errorMessages.BadLength
	}
	if len(myHash) != 64 {
		return errorMessages.Hashing
	}

	imgColor := myHash[0:6]

	m := image.NewRGBA(image.Rect(0, 0, length, length))

	oddColor := color.RGBA{imgColor[0], imgColor[1], imgColor[2], 255}
	evenColor := color.RGBA{imgColor[3], imgColor[4], imgColor[5], 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{oddColor}, image.ZP, draw.Src)

	posX, posY, index := 0, 0, 0

	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {

			if x+y == 0 || (x == 0 && y == 5) || (x == 5 && y == 0) || (x == 5 && y == 5) {
				go drawRect(m, oddColor, posX, posY, posX+10, posY+10)
			} else if myHash[index]%2 == 0 {
				go drawRect(m, evenColor, posX, posY, posX+10, posY+10)
				index++
			} else {
				go drawRect(m, oddColor, posX, posY, posX+10, posY+10)
				index++
			}
			posX += 10
		}
		posX = 0
		posY += 10
	}

	f, err := os.OpenFile(string2convert+".png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, m)
	return nil
}

func drawRect(mainImage draw.Image, colorObj color.RGBA, x1 int, y1 int, x2 int, y2 int) {
	temp := image.NewRGBA(image.Rect(x1, y1, x2, y2))
	draw.Draw(mainImage, temp.Bounds(), &image.Uniform{colorObj}, image.ZP, draw.Src)
}
