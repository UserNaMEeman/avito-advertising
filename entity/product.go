package entity

type Product struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL1        string `json:"url1,omitempty"`
	URL2        string `json:"url2,omitempty"`
	URL3        string `json:"url3,omitempty"`
	Price       string `json:"price,omitempty"`
	// DateCreate  time.Time `json:"date_create"`
}

type SortProduct struct {
	PriceSort bool `json:"price_sort,omitempty"`
	DateSort  bool `json:"date_sort,omitempty"`
	Up        bool `json:"up,omitempty"`
}

type SortProductDB struct {
	Field string
	Type  string
}

type ResponseList struct {
	Title string
	URL1  string
	Price string
}
