// Package config
// Date: 2024/7/2 14:38
// Author: Amu
// Description: amvector 的配置文件 config.yml 的生成
package config

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"amvector/pkg/resources"

	"amvector/pkg/utils"

	"gopkg.in/yaml.v2"
)

func Create(prefix, filename string) (*Config, error) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	cfg := new(Config)
	cfg.filename = filename

	if err := cfg.loadDefault(prefix); err != nil {
		return nil, err
	}

	if data, err := io.ReadAll(fp); err != nil {
		return nil, err
	} else if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	cfg.loadVariables()
	return cfg, nil
}

type Config struct {
	filename string
	// 配置项
	Log struct {
		Output   string `yaml:"output"`
		Level    string `yaml:"level"`
		Rotation int    `yaml:"rotation"`
		MaxAge   int    `yaml:"max_age"`
	} `yaml:"log"`

	Variables struct {
		ImageTag        string `yaml:"image_tag"`
		HostPrefix      string `yaml:"host_prefix"`
		ContainerPrefix string `yaml:"container_prefix"`
		// generated variables
		HostResourceDir      string `yaml:"-"`
		HostLogsDir          string `yaml:"-"`
		ContainerResourceDir string `yaml:"-"`
		ContainerLogsDir     string `yaml:"-"`
	} `yaml:"variables"`
}

func (c *Config) loadDefault(prefix string) error {
	c.Variables.ImageTag = "latest"
	c.Variables.HostPrefix = prefix
	c.Variables.ContainerPrefix = "/"
	c.loadVariables()

	c.Log.Output = filepath.Join(c.Variables.HostLogsDir, resources.AmvectorFolder, "vector.log")
	c.Log.Level = "info"
	c.Log.Rotation = 1
	c.Log.MaxAge = 7

	return nil
}

func (c *Config) loadVariables() {
	c.Variables.HostResourceDir = filepath.Join(c.Variables.HostPrefix, resources.RootPath)
	c.Variables.HostLogsDir = filepath.Join(c.Variables.HostPrefix, "logs")

	_ = utils.EnsureDirExists(c.Variables.HostResourceDir)
	_ = utils.EnsureDirExists(c.Variables.HostLogsDir)
	_ = utils.EnsureDirExists(filepath.Join(c.Variables.HostResourceDir, resources.AmvectorSockFolder))

	c.Variables.ContainerResourceDir = filepath.Join(c.Variables.ContainerPrefix, resources.RootPath)
	c.Variables.ContainerLogsDir = filepath.Join(c.Variables.ContainerPrefix, "logs")
}

func (c *Config) Save() error {
	if c.filename == "" {
		return errors.New("empty filename")
	}
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	if err := os.WriteFile(c.filename, out, 0644); err != nil {
		return err
	}
	return nil
}
