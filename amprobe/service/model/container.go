// Package model
// Date: 2024/3/6 11:06
// Author: Amu
// Description:
package model

import (
	"time"
	
	"gorm.io/gorm"
)

type Containers []Container

type Container struct {
	gorm.Model
	Timestamp    time.Time `json:"timestamp"`
	ContainerID  string    `json:"container_id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Network      string    `json:"network"`
	IP           string    `json:"ip"`
	Ports        string    `json:"ports"`
	State        string    `json:"state"`
	Uptime       string    `json:"uptime"`
	CPUPercent   float64   `json:"cpu_percent"`
	MemPercent   float64   `json:"mem_percent"`
	MemUsage     float64   `json:"mem_usage"`
	MemLimit     float64   `json:"mem_limit"`
	Volumes      string    `json:"volumes"`
	Environments string    `json:"environments"`
	Commands     string    `json:"commands"`
	Labels       string    `json:"labels"`
}

func (d *Container) TableName() string {
	return "s_container"
}

// func (d *Container) AfterCreate(tx *gorm.DB) error {
// 	// 数据清理
// 	tx.Unscoped().Where("timestamp < ?", time.Now().Add(-time.Minute*5)).Delete(&Container{})
// 	return nil
// }

type Docker struct {
	gorm.Model
	Timestamp     time.Time
	DockerVersion string
	APIVersion    string
	MinAPIVersion string
	GitCommit     string
	GoVersion     string
	Os            string
	Arch          string
}

func (d *Docker) TableName() string {
	return "s_docker"
}

// func (d *Docker) AfterCreate(tx *gorm.DB) error {
// 	tx.Unscoped().Where("timestamp < ?", time.Now().Add(-time.Minute*5)).Delete(&Docker{})
// 	return nil
// }

type Images []Image

type Image struct {
	gorm.Model
	Timestamp time.Time
	ImageID   string
	Name      string
	Tag       string
	Created   string
	Size      string
	Number    int
}

func (i *Image) TableName() string { return "s_image" }

// func (i *Image) AfterCreate(tx *gorm.DB) error {
// 	tx.Unscoped().Where("timestamp < ?", time.Now().Add(-time.Minute*5)).Delete(&Image{})
// 	return nil
// }

type Networks []Network

type Network struct {
	gorm.Model
	Timestamp time.Time
	NetworkID string
	Name      string
	Driver    string
	Scope     string
	Created   string
	Internal  bool
	Subnet    string
	Gateway   string
	Labels    string
}

func (n *Network) TableName() string { return "s_network" }
