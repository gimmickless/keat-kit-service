package domain

type Ingredient struct {
	ID         string  `json:"id"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Unit       string  `json:"unit"`
	Size       string  `json:"size"`
	ImgPath    string  `json:"imgPath"`
	UnitEnergy float64 `json:"unitEnergy"` // in kcal
}
