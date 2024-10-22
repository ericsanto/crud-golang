package main

import (
	"crud-golanfg/src/controller/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api := router.Group("/api/v1")

	routes.Routes(api)

	router.Run(":8080")

}
