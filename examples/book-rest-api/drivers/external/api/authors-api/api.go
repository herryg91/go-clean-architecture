package authors_api

/*
	put implementation of all dependencies to rest server in:
		drivers/external/api/

	this is just example
*/
type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AuthorsCli interface {
	GetAll() ([]Author, error)
}

type client struct{}

func NewClient() AuthorsCli {
	return &client{}
}

func (s *client) GetAll() ([]Author, error) {
	// url := "https://www.localhost.com:8081/authors"
	// logic implementation of GetAll()
	return []Author{}, nil
}
