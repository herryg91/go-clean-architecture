package entity

type BookInfo struct {
	Book
	Authors []*BookAuthor `json:"authors"`
}

func (BookInfo) New(book Book, authors []*BookAuthor) *BookInfo {
	return &BookInfo{
		Book:    book,
		Authors: authors,
	}
}

type BookAuthor struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (BookAuthor) FromAuthor(author *Author) *BookAuthor {
	return &BookAuthor{
		Id:   author.Id,
		Name: author.Name,
	}
}
