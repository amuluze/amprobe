// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import "time"

type Docker struct {
	Timestamp     time.Time `json:"timestamp"`
	DockerVersion string    `json:"docker_version"`
	APIVersion    string    `json:"api_version"`
	MinAPIVersion string    `json:"min_api_version"`
	GitCommit     string    `json:"git_commit"`
	GoVersion     string    `json:"go_version"`
	Os            string    `json:"os"`
	Arch          string    `json:"arch"`
}

type DockerSummaryArgs struct {
	Timestamp time.Time `json:"timestamp"`
}

type DockerSummaryReply struct {
	Data Docker `json:"data"`
}

type Container struct {
	Timestamp    time.Time `json:"timestamp"`
	ContainerID  string    `json:"container_id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Network      string    `json:"network"`
	IP           string    `json:"ip"`
	Ports        string    `json:"ports"`
	State        string    `json:"state"`
	Uptime       string    `json:"uptime"`
	Volumes      string    `json:"volumes"`
	Environments string    `json:"environments"`
	Labels       string    `json:"labels"`
	CPUPercent   float64   `json:"cpu_percent"`
	MemPercent   float64   `json:"mem_percent"`
	MemUsage     float64   `json:"mem_usage"`
	MemLimit     float64   `json:"mem_limit"`
}

type ContainerSummaryArgs struct {
	Timestamp time.Time `json:"timestamp"`
}

type ContainerSummaryReply struct {
	Data []Container `json:"data"`
}

type Image struct {
	Timestamp time.Time `json:"timestamp"`
	ImageID   string    `json:"image_id"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
	Created   string    `json:"created"`
	Size      string    `json:"size"`
}

type ImageSummaryArgs struct {
	Timestamp time.Time `json:"timestamp"`
}

type ImageSummaryReply struct {
	Data []Image `json:"data"`
}

type Network struct {
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

type NetworkSummaryArgs struct {
	Timestamp time.Time `json:"timestamp"`
}

type NetworkSummaryReply struct {
	Data []Network `json:"data"`
}
