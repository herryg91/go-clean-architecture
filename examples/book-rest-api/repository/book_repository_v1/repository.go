package book_repository_v1

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	irepository "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repository struct {
	db      *gorm.DB
	rdsPool *redis.Pool
	rdsTtl  int
}

func New(db *gorm.DB, rdsPool *redis.Pool) irepository.BookRepository {
	return &repository{db, rdsPool, 60}
}

func (r *repository) Get(id int) (*entity.Book, error) {
	bookData := &BookWithRating{}

	rdsKey := fmt.Sprintf("book:%d", id)
	// Get From Cache if Possible
	cachedData, err := r.getBookFromCache(rdsKey)
	if err != nil {
		logrus.Warn(err)
	} else if cachedData != nil {
		return cachedData, nil
	}

	// Fallback if there is not cache
	err = r.db.Raw(`SELECT b.*, AVG(br.rating) as rating
		FROM books b LEFT JOIN book_rating br on b.id = br.book_id
		WHERE b.id = ?
		GROUP BY b.id
	`, id).First(&bookData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrBookNotFound
		}
		return nil, err
	}

	// Set To Cache
	out := bookData.ToBookEntity()
	err = r.setBookToCache(rdsKey, out, r.rdsTtl)
	if err != nil {
		logrus.Warn(err)
	}

	return out, nil
}

func (r *repository) getBookFromCache(key string) (out *entity.Book, err error) {
	out = nil
	rdsConn := r.rdsPool.Get()
	defer rdsConn.Close()

	cachedData, errRds := redis.String(rdsConn.Do("GET", key))
	if errRds != nil {
		if !errors.Is(errRds, redis.ErrNil) {
			err = errors.New("[21001] redis error: " + errRds.Error())
			return
		}
		err = nil
		return
	}
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &out)
		if err != nil {
			return
		}
	}
	return
}
func (r *repository) setBookToCache(key string, data interface{}, ttl int) error {
	rdsConn := r.rdsPool.Get()
	defer rdsConn.Close()

	marshalledData, err := json.Marshal(&data)
	if err != nil {
		return errors.New("[21002] marshal error: " + err.Error())
	}
	_, err = rdsConn.Do("SETEX", key, ttl, string(marshalledData))
	if err != nil {
		return errors.New("[21003] redis error: " + err.Error())
	}

	return nil
}

func (r *repository) Search(keyword string) ([]*entity.Book, error) {
	bookDatas := []*BookWithRating{}
	err := r.db.Raw(`SELECT b.*, AVG(br.rating) as rating
		FROM books b LEFT JOIN book_rating br on b.id = br.book_id
		WHERE b.title like ?
		GROUP BY b.id
	`, "%"+keyword+"%").First(&bookDatas).Error
	if err != nil {
		return nil, err
	}

	resp := []*entity.Book{}
	for _, b := range bookDatas {
		resp = append(resp, b.ToBookEntity())
	}
	return resp, nil
}
