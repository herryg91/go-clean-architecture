package book_author_datasource

import (
	"log"
	"time"

	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDatasource struct {
	db        *gorm.DB
	tableName string
}

func NewMysqlDatasource(db *gorm.DB) *MysqlDatasource {
	return &MysqlDatasource{db, "book_authors"}
}

func (r *MysqlDatasource) GetAuthorIdsByBookId(transaction *gorm.DB, bookId int) ([]int, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	datas := []*BookAuthorModel{}
	err := db.Table(r.tableName).Where("book_id = ?", bookId).Find(&datas).Error
	if err != nil {
		return []int{}, err
	}
	result := []int{}
	for _, data := range datas {
		result = append(result, data.AuthorId)
	}
	return result, err
}

func (r *MysqlDatasource) GetBookIdsByAuthorId(transaction *gorm.DB, authorId int) ([]int, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	datas := []*BookAuthorModel{}
	err := db.Table(r.tableName).Where("author_id = ?", authorId).Find(&datas).Error
	if err != nil {
		return []int{}, err
	}
	result := []int{}
	for _, data := range datas {
		result = append(result, data.BookId)
	}
	return result, err
}

func (r *MysqlDatasource) Create(transaction *gorm.DB, in BookAuthorModel) (*BookAuthorModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}
	timeNow := time.Now()
	in.CreatedAt = &timeNow
	in.UpdatedAt = &timeNow

	err := db.Table(r.tableName).Create(&in).Error
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (r *MysqlDatasource) CreateBulk(transaction *gorm.DB, bookId int, authorIds []int) error {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	timeNow := time.Now()
	sqlStr := "INSERT INTO " + r.tableName + "(book_id, author_id, created_at, updated_at) VALUES "
	vals := []interface{}{}
	for _, authorId := range authorIds {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, bookId, authorId, &timeNow, &timeNow)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	log.Println(sqlStr)
	err := db.Exec(sqlStr, vals...).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MysqlDatasource) Delsert(transaction *gorm.DB, bookId int, authorIds []int) error {
	db := r.db
	if transaction != nil {
		db = transaction
	}
	err := r.DeleteByBookId(db, bookId)
	if err != nil {
		return err
	}
	err = r.CreateBulk(db, bookId, authorIds)
	if err != nil {
		return err
	}
	return nil
}

func (r *MysqlDatasource) Delete(transaction *gorm.DB, bookId, authorId int) error {
	db := r.db
	if transaction != nil {
		db = transaction
	}
	return db.Table(r.tableName).Delete(&entity.Author{}, "book_id = ? AND author_id = ? ", bookId, authorId).Error
}

func (r *MysqlDatasource) DeleteByBookId(transaction *gorm.DB, bookId int) error {
	db := r.db
	if transaction != nil {
		db = transaction
	}
	return db.Table(r.tableName).Delete(&entity.Author{}, "book_id = ? ", bookId).Error

}
