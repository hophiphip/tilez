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
	// TODO: Add the image itself (a file) max/min x,y,zoom and improve parser/constructor
}

const minX = 0
const maxX = 3
const minY = 0
const maxY = 3

func New(image, x, y, zoom string) (Image, error) {

	// parse Zoom
	intZoom, err := strconv.Atoi(zoom)
	if err != nil {
		return Image{}, err
	}

	switch intZoom {
	case 1:
	case 2:
	case 4:
		{
			/* DO NOTHING */
		}
	default:
		return Image{}, fmt.Errorf("zoom is not in [1, 2, 4]")
	}

	// parse X
	intX, err := strconv.Atoi(x)
	if err != nil {
		return Image{}, err
	}

	if intX < minX && intX > maxX {
		return Image{}, fmt.Errorf("x is not in [%d..%d]", minX, maxX)
	}

	if intX >= intZoom {
		return Image{}, fmt.Errorf("x: %d can not be greater than zoom: %d", intX, intZoom)
	}

	// parse Y
	intY, err := strconv.Atoi(y)
	if err != nil {
		return Image{}, err
	}

	if intY < minY && intY > maxY {
		return Image{}, fmt.Errorf("y is not in [%d..%d]", minY, maxY)
	}

	if intY >= intZoom {
		return Image{}, fmt.Errorf("y: %d can not be greater than zoom: %d", intY, intZoom)
	}

	// check if image exists
	// TODO: Only .png supported for now

	imagePath := fmt.Sprintf("img/%s.png", image)

	if _, err := os.Stat(imagePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Image{}, err
		} else {
			return Image{}, fmt.Errorf("could not access file")
		}
	}

	return Image{
		X:         intX,
		Y:         intY,
		Zoom:      intZoom,
		ImagePath: imagePath,
	}, nil
}
