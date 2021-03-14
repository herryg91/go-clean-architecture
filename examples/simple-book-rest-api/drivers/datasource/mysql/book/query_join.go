package book_datasource

import (
	"github.com/jinzhu/gorm"
)

// This join under book_datasource because books table is the primray one
func (r *MysqlDatasource) JoinWithAuthors(transaction *gorm.DB, id int) ([]*BookAuthorJoinModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	result := []*BookAuthorJoinModel{}
	err := db.Raw(`
	SELECT 
		book.*,
		author.id as author_id, 
		author.name as author_name, 
		author.birthdate as author_birthdate
	FROM books book
	INNER JOIN book_authors ba on book.id = ba.book_id 
	INNER JOIN authors author on author.id = ba.author_id 
	where book.id = ?
	`, id).Find(&result).Error
	if err != nil {
		return nil, err
	} else if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return result, nil
}

func (r *MysqlDatasource) JoinWithAuthorsAll(transaction *gorm.DB) ([]*BookAuthorJoinModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	result := []*BookAuthorJoinModel{}
	err := db.Raw(`
	SELECT 
		book.*,
		author.id as author_id, 
		author.name as author_name, 
		author.birthdate as author_birthdate
	FROM books book
	INNER JOIN book_authors ba on book.id = ba.book_id 
	INNER JOIN authors author on author.id = ba.author_id 
	where book.id
	`).Find(&result).Error
	if err != nil {
		return nil, err
	} else if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return result, nil
}
