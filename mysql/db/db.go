package db

import (
	dbhelper "github.com/baisiyi/gorm_helper/db_helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option func(db *gorm.DB) *gorm.DB

func GetDbHelper(name string, opts ...Option) *dbhelper.DbHelper {
	db, err := newDb(name)
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	for _, opt := range opts {
		db = opt(db)
	}
	return dbhelper.NewDbHelper(db, dbhelper.WithDbKey(name))
}

var DBMap = make(map[string]*gorm.DB)

func newDb(dsn string) (db *gorm.DB, err error) {
	db, ok := DBMap[dsn]
	if !ok {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		DBMap[dsn] = db
	}
	return db, nil
}
