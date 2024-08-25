// Package main
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package main

import (
	"amvector/pkg/compose"
	"amvector/pkg/config"
	"path/filepath"
)

func runSetup() error {
	if err := setupConfig(prefix, configFile); err != nil {
		return err
	}
	
	if err := setupComposeConfig(prefix); err != nil {
		return err
	}
	
	return nil
}

func setupConfig(prefix, configFile string) error {
	cfg, err := config.Create(prefix, configFile)
	if err != nil {
		return err
	}
	if err := cfg.Save(); err != nil {
		return err
	}
	return nil
}

func setupComposeConfig(prefix string) error {
	filePath := filepath.Join(prefix, "docker-compose.yml")
	err := compose.GenerateDockerCompose(filePath)
	if err != nil {
		return err
	}
	return nil
}
