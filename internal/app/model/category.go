package model

type Category struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Position   int    `json:"position"`
	ImageURL   string `json:"image_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
