package request

type RestaurantRequest struct {
	Name          string `json:"name" binding:"required,min=10,max=200"`
	Password      string `json:"password" binding:"required,min=8,containsany=!@#$%&*"`
	Email         string `json:"email" binding:"required,email"`
	Address       Address
	Telephone     string `json:"telephone" binding:"required,min=11"`
	TypeOfKitchen string `json:"type_of_kitchen"`
	OpeningHours  string `json:"opening_hours"`
	Capacity      int16  `json:"capacity"`
}

type Address struct {
	Street string `json:"street" binding:"required"`
	Number string `json:"number" binding:"required"`
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required,min=2"`
}
