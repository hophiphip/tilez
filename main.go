package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

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

	router.GET("/:image/:zoom/:x/:y", handleImage)

	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatal(err)
	}
}
