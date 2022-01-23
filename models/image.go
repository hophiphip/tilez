package models

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Image struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	Zoom int `json:"zoom"`

	ImagePath string `json:"image_path"`

	// TODO: Add the image itself (imge.Image) max/min x,y,zoom and improve parser/constructor
}

const imageFolderPrefix = "img"

func New(image, x, y, zoom string) (Image, error) {
	img := Image{}

	if err := img.initZoom(zoom); err != nil {
		return Image{}, err
	}

	if err := img.initX(x); err != nil {
		return Image{}, err
	}

	if err := img.initY(y); err != nil {
		return Image{}, err
	}

	if err := img.initImagePath(image); err != nil {
		return Image{}, err
	}

	return img, nil
}

// initZoom check if provided Zoom value is correct and if it is then set it
func (img *Image) initZoom(_zoom string) error {
	zoom, err := strconv.Atoi(_zoom)

	if err != nil {
		return err
	}

	if zoom > 3 {
		return fmt.Errorf("zoom value can not be grater than %d for now", 3)
	}

	img.Zoom = zoom

	return nil
}

// initX check if provided X value is correct and if it is then set it
func (img *Image) initX(_x string) error {
	x, err := strconv.Atoi(_x)

	if err != nil {
		return err
	}

	if x < 0 || x >= img.ZoomAsPowOf2() {
		return fmt.Errorf("x is not in [0..%d]", img.ZoomAsPowOf2()-1)
	}

	img.X = x

	return nil
}

// initY check if provided Y value is correct and if it is then set it
func (img *Image) initY(_y string) error {
	y, err := strconv.Atoi(_y)

	if err != nil {
		return err
	}

	if y < 0 || y >= img.ZoomAsPowOf2() {
		return fmt.Errorf("y is not in [0..%d]", img.ZoomAsPowOf2()-1)
	}

	img.Y = y

	return nil
}

// TODO: Add support for other extensions ? or mb. use vector images
// initImagePath check if provided file does exist and if it is then set it
func (img *Image) initImagePath(_imagePath string) error {
	imagePath := fmt.Sprintf("%s/%s.png", imageFolderPrefix, _imagePath)

	if _, err := os.Stat(imagePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return err
		} else {
			return fmt.Errorf("could not access file")
		}
	}

	img.ImagePath = imagePath

	return nil
}

// ZoomAsPowOf2 return zoom as a power of 2, for example [0,1,2] => [1,2,4]
func (img *Image) ZoomAsPowOf2() int {
	return int(math.Pow(2, float64(img.Zoom)))
}
