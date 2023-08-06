package errorx

import "errors"

var (
	ErrNoPermission        = errors.New("no permission")
	ErrRecordAlreadyExists = errors.New("record already exists")
)
