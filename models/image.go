package models

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Image struct {
	X    int64 `json:"x"`
	Y    int64 `json:"y"`
	Zoom int64 `json:"zoom"`

	ImagePath string `json:"image_path"`
	// TODO: Add the image itself (a file) max/min x,y,zoom and improve parser/constructor
}

const minX = 0
const maxX = 3
const minY = 0
const maxY = 3

func New(image, x, y, zoom string) (Image, error) {

	// parse X
	intX, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		return Image{}, err
	}

	if intX < minX && intX > maxX {
		return Image{}, fmt.Errorf("x is not in [%d..%d]", minX, maxX)
	}

	// parse Y
	intY, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return Image{}, err
	}

	if intY < minY && intY > maxY {
		return Image{}, fmt.Errorf("y is not in [%d..%d]", minY, maxY)
	}

	// parse Zoom
	intZoom, err := strconv.ParseInt(zoom, 10, 64)
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
