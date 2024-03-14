// Package service
// Date: 2024/3/6 11:08
// Author: Amu
// Description:
package service

import (
	"github.com/amuluze/amprobe/pkg/database"
	"github.com/amuluze/amprobe/service/model"
	"strconv"
	"strings"
)

func NewDB(config *Config, models *model.Models) (*database.DB, error) {
	if config.Gorm.GenDoc {
		return nil, nil
	}
	gormConfig := config.Gorm
	dbConfig := config.DB
	db, err := database.NewDB(&database.Config{
		Debug:        gormConfig.Debug,
		Type:         gormConfig.DBType,
		Host:         dbConfig.Host,
		Port:         strconv.Itoa(dbConfig.Port),
		UserName:     dbConfig.User,
		Password:     dbConfig.Password,
		Name:         dbConfig.DBName,
		TablePrefix:  gormConfig.TablePrefix,
		MaxIdleConns: gormConfig.MaxIdleConns,
		MaxLifetime:  gormConfig.MaxLifetime,
		MaxOpenConns: gormConfig.MaxOpenConns,
	})
	if err != nil {
		return nil, err
	}
	if gormConfig.EnableAutoMigrate {
		if dbType := gormConfig.DBType; strings.ToLower(dbType) == "mysql" {
			db.Set("gorm:table_options", "ENGINE=InnoDB")
		}
		err := db.AutoMigrate(models.GetAllModels()...)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
