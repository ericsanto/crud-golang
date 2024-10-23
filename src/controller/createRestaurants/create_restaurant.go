package createRestaurants

import (
	"crud-golanfg/src/controller/model/request"
	"crud-golanfg/src/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(c *gin.Context) {

	var restaurantRequest request.RestaurantRequest

	if err := c.ShouldBindJSON(&restaurantRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "A solicitação não pôde ser processada devido a dados inválidos.",
			"details": err.Error(),
		})
		return
	}

	response := views.NewResponseRequest(restaurantRequest)

	c.JSON(http.StatusOK, response)
}
