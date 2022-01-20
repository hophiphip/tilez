package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const testImage = "test.png"

const defaultHost = "0.0.0.0"
const defaultPort = "8080"

func main() {
	host, isHostSet := os.LookupEnv("TILEZ_HOST")
	if !isHostSet {
		host = defaultHost
	}

	port, isPortSet := os.LookupEnv("TILEZ_PORT")
	if !isPortSet {
		port = defaultPort
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("/img")
	{
		api.GET("/:image/:x/:y/:zoom", handleImage)
	}

	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatal(err)
	}
}
