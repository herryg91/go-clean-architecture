package crud_books

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

func (r *repository) Get(id int) (*entity.BookInfo, error) {
	book, err := r.bookds.Get(nil, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	authorIds, err := r.bookauthords.GetAuthorIdsByBookId(nil, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	authorDatas, err := r.authords.MultiGet(nil, authorIds)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	authors := []*entity.BookAuthor{}
	for _, author := range authorDatas {
		authors = append(authors, author.ToBookAuthorEntity())
	}

	return entity.BookInfo{}.New(*book.ToBookEntity(), authors), nil
}

func (r *repository) GetAll() ([]*entity.BookInfo, error) {
	datas, err := r.bookds.JoinWithAuthorsAll(nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	bookInfoMap := map[int]*entity.BookInfo{}
	for _, data := range datas {
		book, author := data.ToEntity()
		if _, ok := bookInfoMap[data.Id]; !ok {
			bookInfoMap[data.Id] = &entity.BookInfo{Book: *book, Authors: []*entity.BookAuthor{}}
		}
		bookInfoMap[data.Id].Authors = append(bookInfoMap[data.Id].Authors, author.ToBookAuthor())
	}

	out := []*entity.BookInfo{}
	for _, bookInfo := range bookInfoMap {
		out = append(out, bookInfo)
	}

	return out, err
}

func (r *repository) Create(in entity.Book, authorIds []int) (int, error) {
	out := 0
	err := r.db.Transaction(func(tx *gorm.DB) error {
		book, err := r.bookds.Create(tx, *book_datasource.BookModel{}.FromBookEntity(&in))
		if err != nil {
			return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}

		err = r.bookauthords.CreateBulk(tx, book.Id, authorIds)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}
		out = book.Id
		return nil
	})
	if err != nil {
		return 0, err
	}
	return out, err
}

func (r *repository) Update(in entity.Book, authorIds []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		book, err := r.bookds.Update(tx, *book_datasource.BookModel{}.FromBookEntity(&in))
		if err != nil {
			return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}

		err = r.bookauthords.Delsert(tx, book.Id, authorIds)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}

		return nil
	})
}
