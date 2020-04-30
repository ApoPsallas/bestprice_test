package model

type Product struct {
	ID          int     `json:"id"`
	CategoryID  int     `json:"category_id"`
	Title       string  `json:"title"`
	ImageURL    string  `json:"image_url"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Created_at  string  `json:"created_at"`
	Updated_at  string  `json:"updated_at"`
}
