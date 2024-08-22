// Package schema
// Date: 2024/06/11 00:14:42
// Author: Amu
// Description:
package schema

type VersionArgs struct {
}

type ContainerQueryArgs struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type QueryCountArgs struct{}

type QueryCountReply struct {
	Count int `json:"count"`
}

type ContainerCreateArgs struct {
	ContainerName string            `json:"container_name"`
	ImageName     string            `json:"image_name"`
	NetworkName   string            `json:"network_name"`
	Ports         []string          `json:"ports"`
	Volumes       []string          `json:"volumes"`
	Environments  []string          `json:"environments"`
	Commands      []string          `json:"commands"`
	Labels        map[string]string `json:"labels"`
}

type ContainerCreateReply struct {
	ContainerID string `json:"container_id"`
}

type ContainerUpdateArgs struct {
	ContainerID   string            `json:"container_id"`
	ContainerName string            `json:"container_name"`
	ImageName     string            `json:"image_name"`
	NetworkName   string            `json:"network_name"`
	Ports         []string          `json:"ports"`
	Volumes       []string          `json:"volumes"`
	Environments  []string          `json:"environments"`
	Commands      []string          `json:"commands"`
	Labels        map[string]string `json:"labels"`
}

type ContainerUpdateReply struct {
	ContainerID string `json:"container_id"`
}

type ContainerDeleteArgs struct {
	ContainerID string `json:"container_id"`
}

type ContainerDeleteReply struct {
}

type ContainerStartArgs struct {
	ContainerID string `json:"container_id"`
}

type ContainerStartReply struct {
}

type ContainerStopArgs struct {
	ContainerID string `json:"container_id"`
}

type ContainerStopReply struct {
}

type ContainerRestartArgs struct {
	ContainerID string `json:"container_id"`
}

type ContainerRestartReply struct {
}

type ImageQueryArgs struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type ImagePullArgs struct {
	ImageName string `json:"image_name"`
}

type ImagePullReply struct {
}

type ImageTagArgs struct {
	OldTag string `json:"old_tag"`
	NewTag string `json:"new_tag"`
}

type ImageTagReply struct {
}

type ImageCountArgs struct{}

type ImageCountReply struct {
	Count int `json:"count"`
}

type ImageDeleteArgs struct {
	ImageID string `json:"image_id"`
}

type ImageDeleteReply struct {
}

type ImageImportArgs struct {
	SourceFile string `json:"source_file"`
}

type ImageImportReply struct{}

type ImageExportArgs struct {
	ImageIDs   []string `json:"image_ids"`
	TargetFile string   `json:"target_file"`
}

type ImageExportReply struct{}

type NetworkCreateArgs struct {
	Name    string            `json:"name"`
	Driver  string            `json:"driver"`
	Subnet  string            `json:"subnet"`
	Gateway string            `json:"gateway"`
	Labels  map[string]string `json:"labels"`
}

type NetworkCreateReply struct {
	NetworkID string `json:"network_id"`
}

type NetworkQueryArgs struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type NetworkCountArgs struct{}

type NetworkCountReply struct {
	Count int `json:"count"`
}

type NetworkDeleteArgs struct {
	NetworkID string `json:"network_id"`
}

type NetworkDeleteReply struct {
}
