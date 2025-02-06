// Package service
// Date: 2024/06/10 18:42:59
// Author: Amu
// Description:
package service

import (
	"amvector/service/model"
	"common/database"
)

func NewDB(config *Config, models *model.Models) (*database.DB, error) {
	dbConfig := config.DB
	db, err := database.NewDB(
		database.WithDebug(true),
		database.WithType(dbConfig.DBType),
		database.WithHost(dbConfig.Host),
		database.WithPort(dbConfig.Port),
		database.WithUsername(dbConfig.User),
		database.WithPassword(dbConfig.Password),
		database.WithDBName(dbConfig.DBName),
		database.WithMaxLifetime(7200),
		database.WithMaxOpenConns(150),
		database.WithMaxIdleConns(50),
	)
	if err != nil {
		return nil, err
	}
	// 启用WAL模式
	db.Exec("PRAGMA journal_mode=WAL;")

	err = db.AutoMigrate(models.GetAllModels()...)
	if err != nil {
		return nil, err
	}
	return db, nil
}
