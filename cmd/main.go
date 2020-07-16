package main

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/LEE-WAN/RABUMS/cmd/v1"
)

func main() {
	router := gin.Default()

	version1 := router.Group("/v1")
	{
		version1.GET("/get", v1.Get)
	}

	router.Run(":8080")
}
