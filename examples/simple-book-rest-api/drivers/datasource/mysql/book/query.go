package book_datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDatasource struct {
	db        *gorm.DB
	tableName string
}

func NewMysqlDatasource(db *gorm.DB) *MysqlDatasource {
	return &MysqlDatasource{db, "books"}
}

func (r *MysqlDatasource) Get(transaction *gorm.DB, id int) (*BookModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	result := &BookModel{}
	err := db.Table(r.tableName).Where("id = ?", id).Scan(&result).Error
	return result, err
}

func (r *MysqlDatasource) MultiGet(transaction *gorm.DB, ids []int) ([]*BookModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}
	result := []*BookModel{}
	err := db.Table(r.tableName).Where("id in (?)", ids).Find(&result).Error
	return result, err
}

func (r *MysqlDatasource) GetAll(transaction *gorm.DB) ([]*BookModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	result := []*BookModel{}
	err := db.Table(r.tableName).Find(&result).Error
	return result, err
}

func (r *MysqlDatasource) Create(transaction *gorm.DB, in BookModel) (*BookModel, error) {
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

func (r *MysqlDatasource) Update(transaction *gorm.DB, in BookModel) (*BookModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	timeNow := time.Now()
	in.CreatedAt = nil
	in.UpdatedAt = &timeNow
	err := db.Table(r.tableName).Where("id = ?", in.Id).Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (r *MysqlDatasource) Delete(transaction *gorm.DB, id int) error {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	return db.Table(r.tableName).Delete(&BookModel{}, "id = ?", id).Error
}
