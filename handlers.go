package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hophiphip/tilez/models"
)

// TODO: Refactor

func handleImage(ctx *gin.Context) {

	// parse args
	parsedImage, err := models.New(
		ctx.Param("image"),
		ctx.Param("x"),
		ctx.Param("y"),
		ctx.Param("zoom"))

	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	reader, err := os.Open(parsedImage.ImagePath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("os.Open: %s", err))
		return
	}

	defer func() {
		if err = reader.Close(); err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("reader.Close: %s", err))
			return
		}
	}()

	img, _, err := image.Decode(reader)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("image.Decode: %s", err))
		return
	}

	bounds := img.Bounds()

	// NOTE: For now we suppose the image is square
	newSize := (bounds.Max.X - bounds.Min.X) / parsedImage.Zoom

	newBounds := image.Rectangle{
		Min: image.Point{
			X: newSize * parsedImage.X,
			Y: newSize * parsedImage.Y,
		},
		Max: image.Point{
			X: newSize * (parsedImage.X + 1),
			Y: newSize * (parsedImage.Y + 1),
		},
	}

	newImage := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(newBounds.Min.X, newBounds.Min.Y, newBounds.Max.X, newBounds.Max.Y))

	imgBuffer := new(bytes.Buffer)

	if err = png.Encode(imgBuffer, newImage); err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("png.Encode: %s", err))
		return
	}

	ctx.Data(http.StatusOK, "image/png", imgBuffer.Bytes())
}
