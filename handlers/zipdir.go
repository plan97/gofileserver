package handlers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ZipDir provides a gin.HandlerFunc that will zip the contents of the directory.
func ZipDir(baseDir string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Bind Dir from URI.
		var uri URIDir
		if err := c.ShouldBindUri(&uri); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Convert '/' from Dir to path separator and join it with base directory.
		path := filepath.Join(baseDir, filepath.FromSlash(uri.Dir))

		// Verify validity of the path w.r.t. base directory.
		if !strings.HasPrefix(path, baseDir) {
			c.AbortWithError(http.StatusBadRequest,
				fmt.Errorf("directory '%s' is not permitted", uri.Dir))
			return
		}

		// Take the last directory in the tree as ZIP file name.
		var filename string
		if strings.Compare(baseDir, path) == 0 {
			filename = "Home"
		} else {
			filename = filepath.Base(path)
		}

		c.Header("Content-Disposition",
			fmt.Sprintf("attachment; filename=%s", strconv.Quote(fmt.Sprintf("%s.zip", filename))))
		c.Header("Content-Type", c.Writer.Header().Get("Content-Type"))

		zipWriter := zip.NewWriter(c.Writer)
		defer zipWriter.Close()

		// Walk the directory recursively and add all files to ZIP.
		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			relPath := strings.TrimPrefix(filepath.ToSlash(strings.TrimPrefix(filePath, path)), "/")

			zipFile, err := zipWriter.Create(relPath)
			if err != nil {
				return err
			}

			fsFile, err := os.Open(filePath)
			if err != nil {
				return err
			}

			_, err = io.Copy(zipFile, fsFile)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
