package dbhelper

import (
	"gorm.io/gorm"
)

const dbConnKeyPrefix string = "DbHelperDbConn:"

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

// GetDb 获取db连接
func (hp *DbHelper) GetDb(ctx context.Context) *gorm.DB {
	v := ctx.Value(dbConnKeyPrefix + hp.getDbKey()) // 优先使用ctx中的db，保证在一个事务内
	db, ok := v.(*gorm.DB)
	if ok && db != nil {
		return db.WithContext(ctx)
	}

	return hp.db.WithContext(ctx)
}

func (hp *DbHelper) getDbKey() string {
	dbKey := ""
	if hp.opts != nil {
		dbKey = hp.opts.dbKey
	}
	return dbKey
}

func (hp *DbHelper) withDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbConnKeyPrefix+hp.getDbKey(), db) // set进ctx的时候，带上db的key，防止同个ctx内使用多个db的时候出现覆盖
}

func (hp *DbHelper) Reset(ctx context.Context) context.Context {
	return hp.withDB(ctx, nil)
}