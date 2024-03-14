// Package database
// Date: 2024/3/6 10:54
// Author: Amu
// Description:
package database

import "gorm.io/gorm"

type Option func(db *gorm.DB) *gorm.DB

func OptionDB(db *gorm.DB, options ...Option) *gorm.DB {
	for _, option := range options {
		db = option(db)
	}
	return db
}

func WithInIds(ids []*string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) == 0 {
			return db
		}
		db = db.Where("id IN (?)", ids)
		return db
	}
}

func WithById(id *string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if id == nil {
			return db
		}
		db = db.Where("id = ?", id)
		return db
	}
}

func WithByUserName(username *string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if username == nil {
			return db
		}
		db = db.Where("username = ?", username)
		return db
	}
}

func WithByName(name *string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if name == nil {
			return db
		}
		db = db.Where("name = ?", name)
		return db
	}
}

func WithInNames(names []*string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) == 0 {
			return db
		}
		db = db.Where("names IN (?)", names)
		return db
	}
}

func WithByStatus(status *int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if status == nil {
			return db
		}
		db = db.Where("status = ?", status)
		return db
	}
}

func WithOffset(offset int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if offset <= 0 {
			return db
		}
		db = db.Offset(offset)
		return db
	}
}

func WithLimit(limit int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			return db
		}
		db = db.Limit(limit)
		return db
	}
}

func OrderBy(value interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Order(value)
		return db
	}
}
