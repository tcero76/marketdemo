package demo

type Product struct {
	ID                 int     `json:"id",omitempty`
	Title              string  `json:"title",omitempty`
	Description        string  `json:"description",omitempty`
	Price              float64 `json:"price",omitempty`
	DiscountPercentage float64 `json:"discountPercentage",omitempty`
	Rating             float64 `json:"rating",omitempty`
	Stock              int     `json:"stock",omitempty`
	Brand              string  `json:"brand",omitempty`
	Category           string  `json:"category",omitempty`
	Thumbnail          string  `json:"thumbnail",omitempty`
}
