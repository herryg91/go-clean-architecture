package author_profile

import (
	"fmt"

	author_datasource "github.com/herryg91/go-clean-architecture/examples/book-rest-api/drivers/datasource/mysql/author"
	book_datasource "github.com/herryg91/go-clean-architecture/examples/book-rest-api/drivers/datasource/mysql/book"
	book_author_datasource "github.com/herryg91/go-clean-architecture/examples/book-rest-api/drivers/datasource/mysql/book_author"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type repository struct {
	db           *gorm.DB
	bookds       *book_datasource.MysqlDatasource
	authords     *author_datasource.MysqlDatasource
	bookauthords *book_author_datasource.MysqlDatasource
}

func NewRepository(db *gorm.DB, bookds *book_datasource.MysqlDatasource, authords *author_datasource.MysqlDatasource, bookauthords *book_author_datasource.MysqlDatasource) Repository {
	return &repository{db, bookds, authords, bookauthords}
}

func (r *repository) GetAuthorProfile(id int) (*entity.AuthorProfile, error) {
	author, err := r.authords.Get(nil, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	bookIds, err := r.bookauthords.GetBookIdsByAuthorId(nil, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	bookDatas, err := r.bookds.MultiGet(nil, bookIds)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	books := []entity.Book{}
	for _, book := range bookDatas {
		books = append(books, *book.ToBookEntity())
	}
	return entity.AuthorProfile{}.New(*author.ToAuthorEntity(), books), nil
}

func (r *repository) Search(keyword string) ([]*entity.AuthorProfile, error) {
	datas, err := r.authords.SearchWithJoinAuthors(nil, keyword)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	authorProfilesMap := map[int]*entity.AuthorProfile{}
	for _, data := range datas {
		author, book := data.ToEntity()
		if _, ok := authorProfilesMap[data.Id]; !ok {
			authorProfilesMap[data.Id] = entity.AuthorProfile{}.New(*author, []entity.Book{})
		}
		authorProfilesMap[data.Id].Books = append(authorProfilesMap[data.Id].Books, *book)
	}

	out := []*entity.AuthorProfile{}
	for _, authorProfile := range authorProfilesMap {
		out = append(out, authorProfile)
	}
	return out, nil
}
