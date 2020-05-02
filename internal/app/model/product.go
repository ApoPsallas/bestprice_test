package model

import "net/url"

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

func (p *Product) Validate() url.Values {
	errs := url.Values{}

	if p.CategoryID == 0 {
		errs.Add("category_id", "The category_id field is required!")
	}

	if p.Title == "" {
		errs.Add("title", "The title field is required!")
	}

	if p.ImageURL == "" {
		errs.Add("image_url", "The image_url field is required!")
	}

	if p.Price == 0 {
		errs.Add("price", "The price field is required!")
	}

	if p.Description == "" {
		errs.Add("description", "The description field is required!")
	}

	return errs
}
