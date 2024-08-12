// Package model
// Date: 2024/06/10 18:44:53
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
	Labels       string    `json:"labels"`
}

func (d *Container) TableName() string {
	return "s_container"
}

type Docker struct {
	gorm.Model
	Timestamp     time.Time `json:"timestamp"`
	DockerVersion string    `json:"docker_version"`
	APIVersion    string    `json:"api_version"`
	MinAPIVersion string    `json:"min_api_version"`
	GitCommit     string    `json:"git_commit"`
	GoVersion     string    `json:"go_version"`
	Os            string    `json:"os"`
	Arch          string    `json:"arch"`
}

func (d *Docker) TableName() string {
	return "s_docker"
}

type Images []Image

type Image struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	ImageID   string    `json:"image_id"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
	Created   string    `json:"created"`
	Size      string    `json:"size"`
	Number    int       `json:"number"`
}

func (i *Image) TableName() string { return "s_image" }

type Networks []Network

type Network struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	NetworkID string    `json:"network_id"`
	Name      string    `json:"name"`
	Driver    string    `json:"driver"`
	Scope     string    `json:"scope"`
	Created   string    `json:"created"`
	Internal  bool      `json:"internal"`
	Subnet    string    `json:"subnet"`
	Gateway   string    `json:"gateway"`
	Labels    string    `json:"labels"`
}

func (n *Network) TableName() string { return "s_network" }
