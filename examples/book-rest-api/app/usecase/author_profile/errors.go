package author_profile_usecase

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrProfileNotFound = errors.New("Author profile not found")
