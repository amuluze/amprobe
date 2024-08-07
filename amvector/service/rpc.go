// Package service
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package service

import (
	"github.com/amuluze/amprobe/amvector/pkg/resources"
	"github.com/amuluze/amprobe/amvector/service/rpc"
	"github.com/amuluze/amutool/database"
	"github.com/amuluze/docker"
	"github.com/smallnest/rpcx/server"
	"path/filepath"
)

type Server struct {
	address string
	server  *server.Server
}

func NewRPCServer(config *Config, db *database.DB) (*Server, error) {
	srv := server.NewServer()
	manager, err := docker.NewManager()
	if err != nil {
		return nil, err
	}
	s := rpc.NewService(db, manager)

	err = srv.Register(s, "")
	if err != nil {
		return nil, err
	}

	return &Server{
		address: filepath.Join(string(config.prefix), resources.RootPath, resources.AmvectorSockFile),
		server:  srv,
	}, nil
}

func (s *Server) Start() error {
	return s.server.Serve("unix", s.address)
}

func (s *Server) Stop() error {
	return s.server.Close()
}
