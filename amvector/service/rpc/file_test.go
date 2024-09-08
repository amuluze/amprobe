// Package rpc
// Date: 2024/06/25 00:25:29
// Author: Amu
// Description:
package rpc

import (
	"common/rpc"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/amuluze/docker"
)

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func TestFilesSearch(t *testing.T) {

	args := &rpc.FilesSearchArgs{
		Path: "/Users/amu",
	}
	files, err := os.ReadDir(args.Path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			size, _ := DirSize(filepath.Join(args.Path, file.Name()))
			fmt.Println(size)
		}
	}
}

func TestFileDelete(t *testing.T) {
	manager, err := docker.NewManager()
	if err != nil {
		t.Fatalf("new manager error: %#v", err)
	}
	service := NewService(manager)
	args := rpc.FileDeleteArgs{
		Filepath: "/wget-log.3",
	}
	var reply rpc.FileDeleteReply
	err = service.FileDelete(context.Background(), args, &reply)
	if err != nil {
		t.Fatalf("file delete error: %#v", err)
	}
	t.Log("delete successful")
}
