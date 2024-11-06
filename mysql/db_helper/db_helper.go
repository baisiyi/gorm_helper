package dbhelper

import (
	"gorm.io/gorm"
)

type options struct {
	dbKey string
}

type DbHelper struct {
	db   *gorm.DB
	opts *options
}

type Option func(opts *options)

var db *gorm.DB

func NewDbHelper(db *gorm.DB, opts ...Option) *DbHelper {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}
	return &DbHelper{
		db: db,
	}
}

func WithDbKey(name string) func(o *options) {
	return func(o *options) {
		o.dbKey = name
	}
}
