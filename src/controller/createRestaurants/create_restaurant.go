package createRestaurants

import (
	"crud-golanfg/src/controller/model/request"
	"crud-golanfg/src/controller/model/response"
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

	restaurantResponse := response.RestaurantResponse{
		ID:            1,
		Name:          restaurantRequest.Name,
		Address:       restaurantRequest.Address.Street + ", " + restaurantRequest.Address.Number + ", " + restaurantRequest.Address.City + "/" + restaurantRequest.Address.State,
		Telephone:     restaurantRequest.Telephone,
		TypeOfKitchen: restaurantRequest.TypeOfKitchen,
		OpeningHours:  restaurantRequest.OpeningHours,
		Capacity:      restaurantRequest.Capacity,
	}

	c.JSON(http.StatusOK, restaurantResponse)
}
