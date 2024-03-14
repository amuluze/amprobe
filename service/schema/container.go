// Package schema
// Date: 2024/3/6 13:20
// Author: Amu
// Description:
package schema

type Container struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Image         string  `json:"image"`
	IP            string  `json:"ip"`
	State         string  `json:"state"`
	Uptime        string  `json:"uptime"`
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryPercent float64 `json:"memory_percent"`
	MemoryUsage   float64 `json:"memory_usage"`
	MemoryLimit   float64 `json:"memory_limit"`
}

type ContainerQueryRely struct {
	Containers []Container `json:"containers"`
}
