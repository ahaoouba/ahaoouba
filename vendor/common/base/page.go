package base

type Page struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func NewPage(total, offset, limit int) *Page {
	return &Page{
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}
}
