package domain

// model or entity ->existence
type Product struct {
	ID          int     `json:"id" db:"id"` //if small charac needed then `json:"id"` is the format.Just write it to the right side.
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imgUrl" db:"img_url"`
}
