// Package model
// Date: 2024/3/6 11:06
// Author: Amu
// Description:
package model

import (
	"gorm.io/gorm"
	"time"
)

type Containers []Container

type Container struct {
	gorm.Model
	Timestamp   time.Time
	ContainerID string
	Name        string
	Image       string
	IP          string
	State       string
	Uptime      string
	CPUPercent  float64
	MemPercent  float64
	MemUsage    float64
	MemLimit    float64
}

func (d *Container) TableName() string {
	return "s_container"
}

func (d *Container) AfterCreate(tx *gorm.DB) error {
	// 数据清理
	tx.Unscoped().Where("timestamp < ?", time.Now().Add(-time.Minute*5)).Delete(&Container{})
	return nil
}
