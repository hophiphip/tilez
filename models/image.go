package models

import (
	"errors"
	"fmt"
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

// TODO: These values must be initialized during image border parsing
const zoom1 = 1
const zoom2 = 2
const zoom3 = 4

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

	switch zoom {

	case zoom1:
		fallthrough
	case zoom2:
		fallthrough
	case zoom3:
		{
			img.Zoom = zoom
			return nil
		}

	default:
		return fmt.Errorf("zoom is not in [1, 2, 4]")
	}
}

// initX check if provided X value is correct and if it is then set it
func (img *Image) initX(_x string) error {
	x, err := strconv.Atoi(_x)

	if err != nil {
		return err
	}

	if x < 0 || x >= img.Zoom {
		return fmt.Errorf("x is not in [0..%d]", img.Zoom-1)
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

	if y < 0 || y >= img.Zoom {
		return fmt.Errorf("y is not in [0..%d]", img.Zoom-1)
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
