package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// ListDirFiles provides a gin.HandlerFunc that describes the directories/files in a directory.
func ListDirFiles(dir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DirFiles
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		reqDir := filepath.Join(append([]string{dir}, req.Dir...)...)
		if !strings.HasPrefix(reqDir, dir) {
			c.AbortWithError(http.StatusBadRequest,
				fmt.Errorf("directory '%s' is not permitted", req.Dir))
			return
		}

		fileInfo, err := ioutil.ReadDir(reqDir)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req.FileDescriptor = make([]FileDescriptor, len(fileInfo))
		for i, file := range fileInfo {
			req.FileDescriptor[i] = FileDescriptor{
				IsDir:   file.IsDir(),
				ModTime: file.ModTime().UTC(),
				Mode:    file.Mode().String(),
				Name:    file.Name(),
				Size:    file.Size(),
			}
		}
		c.JSON(http.StatusOK, req)
		return
	}
}
