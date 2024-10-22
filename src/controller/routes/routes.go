package routes

import (
	"crud-golanfg/src/controller/createRestaurants"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {

	r.GET("/getRestaurants")
	r.GET("/getRestaurants/:restaurantsId")
	r.POST("/insertRestaurant", createRestaurants.CreateRestaurant)
	r.PUT("/insertRestaurant/:restaurantsId")
	r.DELETE("/deleteRestaurant/:restaurantsId")
}
