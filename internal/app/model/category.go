package model

import "net/url"

type Category struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Position   int    `json:"position"`
	ImageURL   string `json:"image_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func (c *Category) Validate() url.Values {
	errs := url.Values{}

	if c.Title == "" {
		errs.Add("title", "The title field is required!")
	}

	if c.Position == 0 {
		errs.Add("position", "The position field is required!")
	}

	if c.ImageURL == "" {
		errs.Add("image_url", "The image_url field is required!")
	}

	return errs
}
