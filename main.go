package main

import (
	"crud-golanfg/database"
	"crud-golanfg/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("/home/eric/crud-golang/database/.env"); err != nil {
		panic(err)
	}

	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := gin.Default()

	api := router.Group("/api/v1")

	routes.Routes(api)

	router.Run(":8080")

}
