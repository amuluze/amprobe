// Package service
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package service

import (
	"log/slog"
	"time"
	
	"amprobe/pkg/database"
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

func InitAdapter(db *database.DB) *gormAdapter.Adapter {
	adapter, err := gormAdapter.NewAdapterByDBWithCustomTable(db.DB, &CasbinRule{})
	if err != nil {
		slog.Error("init adapter error", "error", err)
	}
	return adapter
}

func InitCasbin(modelPath ModeConf, adapter *gormAdapter.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	cleanFunc := func() {}
	
	e, err := casbin.NewSyncedEnforcer(string(modelPath))
	if err != nil {
		slog.Error("init casbin error", "error", err)
		return nil, cleanFunc, err
	}
	e.EnableLog(true)
	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	if err != nil {
		slog.Error("init casbin error", "error", err)
		return nil, cleanFunc, err
	}
	e.EnableEnforce(true)
	
	e.StartAutoLoadPolicy(time.Duration(60) * time.Second)
	cleanFunc = func() {
		e.StopAutoLoadPolicy()
	}
	
	return e, cleanFunc, nil
}
