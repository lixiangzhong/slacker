package scope

import (
	"github.com/jinzhu/gorm"
)

type Wheres []func(*gorm.DB) *gorm.DB

func OffsetLimit(offset, limit uint64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if offset > 0 {
			db = db.Offset(offset)
		}
		if limit > 0 {
			db = db.Limit(limit)
		}
		return db
	}
}

func Where(query interface{}, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func Preload(column string, conditions ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(column, conditions...)
	}
}

func Order(column string, reorder ...bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(column, reorder...)
	}
}
