package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hophiphip/tilez/models"
	"image"
	"io/ioutil"
	"net/http"
	"os"
)

func handleImage(ctx *gin.Context) {
	img := ctx.Param("img")
	x := ctx.Param("x")
	y := ctx.Param("y")
	zoom := ctx.Param("zoom")

	// parse args
	parsedImage, err := models.New(img, x, y, zoom)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	reader, err := os.Open(parsedImage.ImagePath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}

	defer func() {
		if err = reader.Close(); err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}()

	imageType, _, err := image.Decode(reader)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}

	// read an img
	imageBytes, err := ioutil.ReadFile(parsedImage.ImagePath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}

	ctx.Data(http.StatusOK, "img/png", imageBytes)
}
