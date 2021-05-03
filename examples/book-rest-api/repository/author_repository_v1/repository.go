package author_repository_v1

import (
	"errors"
	"time"

	irepository "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) irepository.AuthorRepository {
	return &repository{db}
}

func (r *repository) Get(id int) (*entity.Author, error) {
	author := &Author{}
	err := r.db.Table("authors").Where("id = ?", id).First(&author).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrAuthorNotFound
		}
		return nil, err
	}
	return author.ToAuthorEntity(), nil
}
func (r *repository) GetAll() ([]*entity.Author, error) {
	authors := []*Author{}
	err := r.db.Table("authors").Find(&authors).Error
	if err != nil {
		return nil, err
	}
	resp := []*entity.Author{}
	for _, author := range authors {
		resp = append(resp, author.ToAuthorEntity())
	}
	return resp, nil

}
func (r *repository) Create(in entity.Author) (*entity.Author, error) {
	author := Author{}.FromAuthorEntity(in)

	timeNow := time.Now()
	author.CreatedAt = &timeNow
	author.UpdatedAt = &timeNow
	err := r.db.Table("authors").Create(&author).Error
	if err != nil {
		return nil, err
	}
	in.Id = author.Id
	return &in, nil
}
func (r *repository) Search(keyword string) ([]*entity.Author, error) {
	searchResult := []*Author{}
	err := r.db.Table("authors").Where("name like ?", "%"+keyword+"%").Find(&searchResult).Error
	if err != nil {
		return nil, err
	}

	resp := []*entity.Author{}
	for _, author := range searchResult {
		resp = append(resp, author.ToAuthorEntity())
	}
	return resp, nil
}

func (r *repository) MultiGet(ids []int) (map[int]*entity.Author, error) {
	authors := []*Author{}
	err := r.db.Table("authors").Where("id in ?", ids).Find(&authors).Error
	if err != nil {
		return nil, err
	}

	resp := map[int]*entity.Author{}
	for _, author := range authors {
		resp[author.Id] = author.ToAuthorEntity()
	}

	return resp, nil

}
