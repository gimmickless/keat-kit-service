package domain

import (
	"time"
)

type Kit struct {
	ID            string            `json:"id"`
	CategoryIDs   []string          `json:"categoryIds"`
	Name          string            `json:"name"`
	Version       string            `json:"version"`
	Description   string            `json:"description"`
	Status        string            `json:"status"`
	AuthorID      string            `json:"authorId"` // username of author
	Recipe        []PreparationStep `json:"recipe"`
	Energy        float64           `json:"energy"` // in kcal
	Portion       float64           `json:"portion"`
	PrepTime      int               `json:"prepTime"` // in minutes
	CreatedAt     time.Time         `json:"createdAt"`
	LastUpdatedAt time.Time         `json:"lastUpdatedAt"`
	Prices        []Price           `json:"prices"`
}

type PreparationStep struct {
	Ingredient
	Quantity float64 `json:"quantity"`
	Action   string  `json:"action"` // add, pour, etc.
}

type Price struct {
	KitID    string  `json:"kitID"`
	Amount   float64 `json:"amount"`
	Country  string  `json:"country"`  // ISO 3166-1 alpha-2 code
	Currency string  `json:"currency"` // ISO 4217 code of currency
}
