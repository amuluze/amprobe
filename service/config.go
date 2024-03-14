// Package service
// Date: 2024/3/6 11:08
// Author: Amu
// Description:
package service

import (
	"fmt"
	
	"github.com/amuluze/amprobe/pkg/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Fiber    Fiber
	Gorm     Gorm
	DB       DB
	Disk     Disk
	Task     Task
	Ethernet Ethernet
	Logger   Logger
}

// NewConfig Load config file (toml/json/yaml)
func NewConfig(configFile string) (*Config, error) {
	config := &Config{}
	
	viper.SetConfigFile(configFile)
	
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	
	switch config.Logger.Level {
	case "debug":
		logger.SetDebugLevel()
	case "info":
		logger.SetInfoLevel()
	case "warn":
		logger.SetWarnLevel()
	case "error":
		logger.SetErrorLevel()
	default:
		logger.Info("use default logger level")
	}
	
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err := viper.Unmarshal(config); err != nil {
			fmt.Printf("unmarshal config error when change, %v", err)
		}
		switch config.Logger.Level {
		case "debug":
			logger.SetDebugLevel()
		case "info":
			logger.SetInfoLevel()
		case "warn":
			logger.SetWarnLevel()
		case "error":
			logger.SetErrorLevel()
		default:
			logger.Info("use default logger level")
		}
	})
	return config, nil
}

type Fiber struct {
	Host            string
	Port            int
	ShutdownTimeout int
	SeverHeader     string
	AppName         string
	Prefork         bool
}

type Gorm struct {
	GenDoc            bool
	Debug             bool
	DBType            string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Disk struct {
	Devices []string
}

type Task struct {
	Interval int
}

type Ethernet struct {
	Names []string
}

type Logger struct {
	Level string
}
