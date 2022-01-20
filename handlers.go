package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hophiphip/tilez/models"
	"image"
	"image/png"
	_ "image/png"
	"net/http"
	"os"
)

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

	//bounds := img.Bounds()
	//var newBounds image.Rectangle

	imgBuffer := new(bytes.Buffer)

	if err = png.Encode(imgBuffer, img); err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("png.Encode: %s", err))
		return
	}

	ctx.Data(http.StatusOK, "image/png", imgBuffer.Bytes())
}

//reader, err := os.Open(parsedImage.ImagePath)
//if err != nil {
//ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
//return
//}
//
//defer func() {
//	if err = reader.Close(); err != nil {
//		ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
//		return
//	}
//}()
//
//imageType, _, err := image.Decode(reader)
//if err != nil {
//ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
//return
//}
//
//// read an img
//imageBytes, err := ioutil.ReadFile(parsedImage.ImagePath)
//if err != nil {
//ctx.String(http.StatusInternalServerError, fmt.Sprint(err))
//return
//}
//
//ctx.Data(http.StatusOK, "img/png", imageBytes)
