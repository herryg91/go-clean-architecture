package entity

type Book struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	ReleasedYear int     `json:"released_year"`
	Rating       float64 `json:"rating"`
}
