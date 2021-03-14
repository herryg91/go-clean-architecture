package author_datasource

import (
	"github.com/jinzhu/gorm"
)

// This join under author_datasource because authors table is the primray one
func (r *MysqlDatasource) SearchWithJoinAuthors(transaction *gorm.DB, keyword string) ([]*AuthorBookJoinModel, error) {
	db := r.db
	if transaction != nil {
		db = transaction
	}

	result := []*AuthorBookJoinModel{}
	err := db.Raw(`
	SELECT 
		author.*,
		book.id as book_id, 
		book.title as book_title, 
		book.released_year as book_released_year
	FROM authors author
	INNER JOIN book_authors ba on author.id = ba.author_id 
	INNER JOIN books book on book.id = ba.book_id 
	where author.name like ?
	`, "%"+keyword+"%").Find(&result).Error
	if err != nil {
		return nil, err
	} else if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return result, nil
}
