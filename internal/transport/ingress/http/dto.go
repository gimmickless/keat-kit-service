package http

import "time"

type categoryResp struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	ImageURL string `json:"imageUrl"`
}

type ingredientResp struct {
	ID       string  `json:"id"`
	Unit     string  `json:"unit"`
	Size     string  `json:"size"`
	ImageURL string  `json:"imageUrl"`
	Energy   float64 `json:"energy"` // in kcal
}

type kitResp struct {
	ID            string                `json:"id"`
	CategoryIDs   []string              `json:"categoryIds"`
	Name          string                `json:"name"`
	Version       string                `json:"version"`
	Description   string                `json:"description"`
	Status        string                `json:"status"`
	Author        string                `json:"author"` // username of author
	Recipe        []preparationStepResp `json:"recipe"`
	Energy        float64               `json:"energy"` // in kcal
	Portions      float64               `json:"portions"`
	PrepTime      int                   `json:"prepTime"` // in minutes
	CreatedAt     time.Time             `json:"createdAt"`
	LastUpdatedAt time.Time             `json:"lastUpdatedAt"`
	Price         priceResp             `json:"price"`
}

// Auxiliary
type preparationStepResp struct {
	Ingredient ingredientResp `json:"ingredient"`
	Quantity   float64        `json:"quantity"`
	Action     string         `json:"action"` // add, pour, etc.
}

type priceResp struct {
	KitID    string `json:"kitID"`
	Country  string `json:"country"`  // ISO 3166-1 alpha-2 code
	Currency string `json:"currency"` // ISO 4217 code of currency
}
