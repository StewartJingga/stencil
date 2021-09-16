package models

import (
	"mime/multipart"

	"github.com/odpf/stencil/server/snapshot"
)

type FileDownloadRequest struct {
	Namespace string `uri:"namespace" binding:"required"`
	Name      string `uri:"name" binding:"required"`
	Version   string `uri:"version" binding:"required,version|eq=latest"`
	FullNames []string
}

// TODO: fix this, a tag who is latest and user passes actual version doesn't work correctly
func (f *FileDownloadRequest) IsLatest() *bool {
	var (
		isLatest *bool = nil
		trueBool = true
		falseBool = false
	)
	if f.Version == "latest" {
		isLatest = &trueBool
	} else if f.Version != "" {
		isLatest = &falseBool
	}
	return isLatest
}

// ToSnapshot creates snapshot
func (f *FileDownloadRequest) ToSnapshot() *snapshot.Snapshot {
	s := &snapshot.Snapshot{
		Namespace: f.Namespace,
		Name:      f.Name,
	}
	// TODO: fix this
	if f.Version == "latest" {
		s.Latest = true
	} else {
		s.Version = f.Version
	}
	return s
}

type DescriptorUploadRequest struct {
	Namespace string                `uri:"namespace" binding:"required"`
	Name      string                `form:"name" binding:"required"`
	Version   string                `form:"version" binding:"required,version"`
	File      *multipart.FileHeader `form:"file" binding:"required"`
	Latest    bool                  `form:"latest"`
	SkipRules []string              `form:"skiprules"`
	DryRun    bool                  `form:"dryrun"`
}

// ToSnapshot creates sanpshot
func (d *DescriptorUploadRequest) ToSnapshot() *snapshot.Snapshot {
	return &snapshot.Snapshot{
		Namespace: d.Namespace,
		Name:      d.Name,
		Version:   d.Version,
		Latest:    d.Latest,
	}
}
