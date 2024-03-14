// Package database
// Date: 2024/3/6 10:54
// Author: Amu
// Description:
package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Debug        bool
	Type         string
	Host         string
	Port         string
	UserName     string
	Password     string
	Name         string
	TablePrefix  string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func (c *Config) Dial() gorm.Dialector {
	var dsn string
	var dialector gorm.Dialector
	switch c.Type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.UserName,
			c.Password,
			c.Host,
			c.Port,
			c.Name,
		)
		dialector = mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		})
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
			c.Host,
			c.Port,
			c.UserName,
			c.Name,
			c.Password,
		)
		dialector = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	default:
		dsn = fmt.Sprintf("%s.db", c.Name)
		dialector = sqlite.Open(dsn)
	}

	return dialector
}

type DB struct {
	*gorm.DB
}

func NewDB(cfg *Config) (*DB, error) {
	dial := cfg.Dial()
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 开启调试模式
	if cfg.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(cfg.MaxLifetime))

	return &DB{db}, nil
}

func (db *DB) Close() {
	if db != nil {
		db.Close()
	}
}

func (db *DB) RunInTransaction(fn func(tx *gorm.DB) error) error {
	return db.DB.Transaction(fn)
}

func (db *DB) Set(key string, value interface{}) {
	db.DB.Set(key, value)
}

func (db *DB) AutoMigrate(models ...interface{}) error {
	return db.DB.AutoMigrate(models...)
}

func (db *DB) GetModel(model interface{}) *gorm.DB {
	return db.DB.Model(model)
}
