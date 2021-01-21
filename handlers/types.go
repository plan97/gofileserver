package handlers

import (
	"time"
)

// DirFiles is used to list the files in a directory.
type DirFiles struct {
	Dir            []string         `json:"dir"`
	FileDescriptor []FileDescriptor `json:"file_descriptor,omitempty"`
}

// FileDescriptor provides the directory/file description.
type FileDescriptor struct {
	IsDir   bool      `json:"is_dir"`
	ModTime time.Time `json:"mod_time,omitempty"`
	Mode    string    `json:"mode,omitempty"`
	Name    string    `json:"name,omitempty"`
	Size    int64     `json:"size,omitempty"`
}
