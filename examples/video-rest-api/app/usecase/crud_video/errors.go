package crud_video

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrVideoNotFound = errors.New("Video not found")
