package book_page

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrBookNotFound = errors.New("ErrBookNotFound")
