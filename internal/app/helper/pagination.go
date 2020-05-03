package helper

import "strconv"

type Pagination struct {
	Page    int
	PerPage int
	Order   string
}

//func NewPagination(p int, pp int, ob string, o string) *Pagination {
func NewPagination(p map[string][]string) *Pagination {
	pn := Pagination{Page: 1, PerPage: 10, Order: "asc"}

	value, found := p["page"]
	if found {
		page, err := strconv.Atoi(value[0])
		if err == nil {
			pn.Page = page
		}
	}

	value, found = p["per_page"]
	if found {
		perPage, err := strconv.Atoi(value[0])
		if err == nil {
			pn.PerPage = perPage
		}
	}

	value, found = p["order"]
	if found {
		if value[0] == "asc" || value[0] == "desc" {
			pn.Order = value[0]
		}
	}

	return &pn
}
