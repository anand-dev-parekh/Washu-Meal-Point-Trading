package models

type Offer struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"userID"`
	MealPoints uint16 `json:"mealPoints"`
	Email      string `json:"email"`
}
