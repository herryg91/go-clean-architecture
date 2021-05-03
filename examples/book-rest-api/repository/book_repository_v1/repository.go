package book_repository_v1

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

func New(db *gorm.DB) irepository.BookRepository {
	return &repository{db}
}

func (r *repository) Get(id int) (*entity.Book, error) {
	bookData := &Book{}
	err := r.db.Table("books").Where("id = ?", id).First(&bookData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrBookNotFound
		}
		return nil, err
	}
	bookAuthorDatas := []*BookAuthor{}
	err = r.db.Table("book_authors").Where("book_id = ?", id).Find(&bookAuthorDatas).Error
	if err != nil {
		return nil, err
	}

	resp := bookData.ToBookEntity()
	for _, bookAuthor := range bookAuthorDatas {
		resp.AddAuthor(*bookAuthor.ToBookAuthorEntity())
	}
	return resp, nil
}

func (r *repository) GetAll() ([]*entity.Book, error) {
	bookDatas := []*Book{}
	err := r.db.Table("books").Find(&bookDatas).Error
	if err != nil {
		return nil, err
	}

	bookIds := []int{}
	for _, b := range bookDatas {
		bookIds = append(bookIds, b.Id)
	}
	mapBookAuthorDatas := map[int][]*BookAuthor{}
	if len(bookIds) > 0 {
		bookAuthorDatas := []*BookAuthor{}
		err = r.db.Table("book_authors").Where("book_id in (?)", bookIds).Find(&bookAuthorDatas).Error
		if err != nil {
			return nil, err
		}
		for _, ba := range bookAuthorDatas {
			if _, ok := mapBookAuthorDatas[ba.BookId]; !ok {
				mapBookAuthorDatas[ba.BookId] = []*BookAuthor{}
			}
			mapBookAuthorDatas[ba.BookId] = append(mapBookAuthorDatas[ba.BookId], ba)
		}
	}

	resp := []*entity.Book{}
	for _, b := range bookDatas {
		tmpResp := b.ToBookEntity()
		if authoDatas, ok := mapBookAuthorDatas[b.Id]; ok {
			for _, authorData := range authoDatas {
				tmpResp.AddAuthor(*authorData.ToBookAuthorEntity())
			}
		}
		resp = append(resp, tmpResp)
	}
	return resp, nil
}

func (r *repository) Create(in entity.Book) (*entity.Book, error) {
	bookModel := Book{}.FromBookEntity(&in)

	timeNow := time.Now()
	bookModel.CreatedAt = &timeNow
	bookModel.UpdatedAt = &timeNow

	tx := r.db.Begin()

	err := tx.Table("books").Create(&bookModel).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	in.Id = bookModel.Id

	for _, author := range in.Authors {
		bookAuthorModel := BookAuthor{}.FromBookAuthorEntity(bookModel.Id, &author)
		err = tx.Table("book_authors").Create(&bookAuthorModel).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &in, nil
}
