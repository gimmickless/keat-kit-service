package domain

type Category struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	ImgPath string `json:"imagePath"`
}
