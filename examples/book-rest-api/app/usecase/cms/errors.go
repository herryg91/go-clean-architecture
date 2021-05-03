package cms_usecase

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrBookNotFound = errors.New("Book not found")
var ErrAuthorNotFound = errors.New("Author not found")
