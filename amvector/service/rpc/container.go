// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"common/rpc"
	"context"
	
	"github.com/amuluze/docker"
)

func (s *Service) ContainerCreate(ctx context.Context, args rpc.ContainerCreateArgs, reply *rpc.ContainerCreateReply) error {
	var containerID string
	var err error
	if containerID, err = s.Manager.CreateContainer(
		ctx,
		args.Name,
		args.Image,
		args.Network,
		args.Ports,
		args.Volumes,
		args.Environments,
		args.Commands,
		args.Labels,
	); err != nil {
		return err
	}
	reply.ContainerID = containerID
	return nil
}

func (s *Service) ContainerUpdate(ctx context.Context, args rpc.ContainerUpdateArgs, reply *rpc.ContainerUpdateReply) error {
	if exists, err := s.Manager.ContainerExists(ctx, args.ContainerID); err != nil {
		return err
	} else if exists {
		if err := s.Manager.DeleteContainer(ctx, args.ContainerID); err != nil {
			return err
		}
	}
	
	cID, err := s.Manager.CreateContainer(
		ctx,
		args.Name,
		args.Image,
		args.Network,
		args.Ports,
		args.Volumes,
		args.Environments,
		args.Commands,
		args.Labels,
	)
	if err != nil {
		return err
	}
	reply.ContainerID = cID
	return nil
}

func (s *Service) ContainerDelete(ctx context.Context, args rpc.ContainerDeleteArgs, reply *rpc.ContainerDeleteReply) error {
	if err := s.Manager.DeleteContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerStart(ctx context.Context, args rpc.ContainerStartArgs, reply *rpc.ContainerStartReply) error {
	if err := s.Manager.StartContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerStop(ctx context.Context, args rpc.ContainerStopArgs, reply *rpc.ContainerStopReply) error {
	if err := s.Manager.StopContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerRestart(ctx context.Context, args rpc.ContainerRestartArgs, reply *rpc.ContainerRestartReply) error {
	if err := s.Manager.RestartContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImagePull(ctx context.Context, args rpc.ImagePullArgs, reply *rpc.ImagePullReply) error {
	if err := s.Manager.PullImage(ctx, args.ImageName); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImageTag(ctx context.Context, args rpc.ImageTagArgs, reply *rpc.ImageTagReply) error {
	if err := s.Manager.TagImage(ctx, args.OldTag, args.NewTag); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImageDelete(ctx context.Context, args rpc.ImageDeleteArgs, reply *rpc.ImageDeleteReply) error {
	if err := s.Manager.DeleteImage(ctx, args.ImageID); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImagesPrune(ctx context.Context) error {
	return s.Manager.PruneImages(ctx)
}

func (s *Service) ImageImport(ctx context.Context, args rpc.ImageImportArgs, reply *rpc.ImageImportReply) error {
	if err := s.Manager.ImportImage(ctx, args.SourceFile); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImageExport(ctx context.Context, args rpc.ImageExportArgs, reply *rpc.ImageExportReply) error {
	if err := s.Manager.ExportImage(ctx, args.ImageIDs, args.TargetFile); err != nil {
		return err
	}
	return nil
}

func (s *Service) NetworkCreate(ctx context.Context, args rpc.NetworkCreateArgs, reply *rpc.NetworkCreateReply) error {
	args.Labels[docker.CreatedByProbe] = "true"
	if networkID, err := s.Manager.CreateNetwork(ctx, args.Name, args.Driver, args.Subnet, args.Gateway, args.Labels); err != nil {
		return err
	} else {
		reply.NetworkID = networkID
		return nil
	}
}

func (s *Service) NetworkUpdate(ctx context.Context, args rpc.NetworkUpdateArgs, reply *rpc.NetworkUpdateReply) error {
	return nil
}

func (s *Service) NetworkDelete(ctx context.Context, args rpc.NetworkDeleteArgs, reply *rpc.NetworkDeleteReply) error {
	if err := s.Manager.DeleteNetwork(ctx, args.NetworkID); err != nil {
		return err
	}
	return nil
}
