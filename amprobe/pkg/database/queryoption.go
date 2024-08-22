// Package database
// Date       : 2024/8/22 14:25
// Author     : Amu
// Description:
package database

type QueryOption func(db *DB) *DB

func OptionDB(db *DB, options ...QueryOption) *DB {
	for _, option := range options {
		db = option(db)
	}
	return db
}

func WithTable(tableName string) QueryOption {
	return func(db *DB) *DB {
		db.DB = db.DB.Table(tableName)
		return db
	}
}

func WithInIds(ids ...string) QueryOption {
	return func(db *DB) *DB {
		if len(ids) == 0 {
			return db
		}
		db.DB = db.DB.Where("id IN (?)", ids)
		return db
	}
}

func WithById(id string) QueryOption {

	return func(db *DB) *DB {
		if id == "" {
			return db
		}
		db.DB = db.DB.Where("id = ?", id)
		return db
	}
}

func WithOffset(offset int) QueryOption {
	if offset < 0 {
		return nil
	}
	return func(db *DB) *DB {
		db.DB = db.DB.Offset(offset)
		return db
	}
}

func WithLimit(limit int) QueryOption {
	if limit <= 0 {
		return nil
	}
	return func(db *DB) *DB {
		db.DB = db.DB.Limit(limit)
		return db
	}
}

func OrderBy(value interface{}) QueryOption {
	return func(db *DB) *DB {
		db.DB = db.DB.Order(value)
		return db
	}
}
