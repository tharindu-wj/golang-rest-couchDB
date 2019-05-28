package model

type Company struct {
	RestaurantID   string `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
	CurrencyCode   string `json:"currency_code"`
	BranchID       string `json:"branch_id"`
	BranchName     string `json:"branch_name"`
}