package domain

type Ingredient struct {
	ID       string  `json:"id"`
	Unit     string  `json:"unit"`
	Size     string  `json:"size"`
	ImageURL string  `json:"imageUrl"`
	Energy   float64 `json:"energy"` // in kcal
}
