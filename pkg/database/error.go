package database

import "errors"

var (
	ErrRequestNotFound     = errors.New("database: request infomation not found.")
	ErrDatabaseUnavailable = errors.New("database: database is not available.")
	ErrRequestCanNotBeNil  = errors.New("databse: request should not be nil.")
)
