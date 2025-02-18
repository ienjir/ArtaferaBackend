package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var GinMode int

func SetGinMode() {
	switch os.Getenv("MODE") {
	case "debug":
		gin.SetMode("debug")
		GinMode = 0
		fmt.Println("Running in debug mode")

	case "test":
		gin.SetMode("test")
		GinMode = 1
		fmt.Println("Running in test mode")

	case "release":
		gin.SetMode("release")
		GinMode = 2
		fmt.Println("Running in release mode")

	default:
		fmt.Println("No application mode was provided")
		fmt.Println("Please set 'MODE' in the .env file to one of the following 3 options:")
		fmt.Println("debug")
		fmt.Println("test")
		fmt.Println("release")
		os.Exit(1)
	}

}
