package views

import (
	"crud-golanfg/src/controller/model/request"
	"crud-golanfg/src/controller/model/response"
)

func NewResponseRequest(req request.RestaurantRequest) response.RestaurantResponse {

	return response.RestaurantResponse{
		Name: req.Name,
		Address: req.Address.Street + ", " + req.Address.Number + ", " +
			req.Address.City + "/" + req.Address.State,
		Telephone:     req.Telephone,
		TypeOfKitchen: req.TypeOfKitchen,
		OpeningHours:  req.OpeningHours,
		Capacity:      req.Capacity,
	}
}
