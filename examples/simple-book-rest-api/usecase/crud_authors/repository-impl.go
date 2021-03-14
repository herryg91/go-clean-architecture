package crud_authors

import (
	"fmt"

	author_datasource "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/drivers/datasource/mysql/author"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type repository struct {
	db *gorm.DB
	ds *author_datasource.MysqlDatasource
}

func NewRepository(db *gorm.DB, ds *author_datasource.MysqlDatasource) Repository {
	return &repository{db, ds}
}

func (r *repository) Get(id int) (*entity.Author, error) {
	out, err := r.ds.Get(nil, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}

	return out.ToAuthorEntity(), err
}

func (r *repository) GetAll() ([]*entity.Author, error) {
	datas, err := r.ds.GetAll(nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	out := []*entity.Author{}
	for _, data := range datas {
		out = append(out, data.ToAuthorEntity())
	}
	return out, err
}

func (r *repository) Create(in entity.Author) (*entity.Author, error) {
	out, err := r.ds.Create(nil, *author_datasource.AuthorModel{}.FromAuthorEntity(in))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out.ToAuthorEntity(), err
}

func (r *repository) Update(in entity.Author) (*entity.Author, error) {
	out, err := r.ds.Update(nil, *author_datasource.AuthorModel{}.FromAuthorEntity(in))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
	}
	return out.ToAuthorEntity(), err
}
