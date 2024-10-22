package response

type RestaurantResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	TypeOfKitchen string `json:"type_of_kitchen"`
	OpeningHours  string `json:"opening_hours"`
	Capacity      int16  `json:"capacity"`
}
