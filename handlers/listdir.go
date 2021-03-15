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
func ListDirFiles(baseDir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DirFiles
		if err := c.BindJSON(&req); err != nil {
			if e := c.AbortWithError(http.StatusBadRequest, err); e.Err == nil {
				fmt.Println(err)
			}
			return
		}

		reqDir := filepath.Join(append([]string{baseDir}, req.Dir...)...)
		if !strings.HasPrefix(reqDir, baseDir) {
			er := fmt.Errorf("access to directory '%s' is not allowed", reqDir)
			if e := c.AbortWithError(http.StatusBadRequest, er); e.Err == nil {
				fmt.Println(er)
			}
			return
		}

		fileInfo, err := ioutil.ReadDir(reqDir)
		if err != nil {
			if e := c.AbortWithError(http.StatusInternalServerError, err); e.Err == nil {
				fmt.Println(err)
			}
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
	}
}
