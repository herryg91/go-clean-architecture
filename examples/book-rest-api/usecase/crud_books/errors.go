package crud_books

import "errors"

var ErrDatabaseError = errors.New("Database Error")
var ErrRecordNotFound = errors.New("Record Not Found")
