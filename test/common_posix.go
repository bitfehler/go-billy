//go:build !windows && !wasip1 && !js && !wasp
// +build !windows,!wasip1,!js,!wasp

package test

import (
	"io/fs"

	"github.com/go-git/go-billy/v6"
	"github.com/go-git/go-billy/v6/memfs"
	"github.com/go-git/go-billy/v6/osfs"
)

var (
	customMode            fs.FileMode = 0o755
	expectedSymlinkTarget             = "/dir/file"
)

func allFS(tempDir func() string) []billy.Filesystem {
	return []billy.Filesystem{
		osfs.New(tempDir(), osfs.WithChrootOS()),
		memfs.New(),
	}
}
