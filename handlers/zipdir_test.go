package handlers

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestZipDir(t *testing.T) {
	type args struct {
		baseDir string
	}
	tests := []struct {
		name     string
		args     args
		testFunc func(gin.HandlerFunc) error
		wantErr  bool
	}{
		{
			name: "test 1",
			args: args{baseDir: func() string {
				dir, err := filepath.Abs("")
				if err != nil {
					return ""
				}
				return dir
			}()},
			testFunc: func(fn gin.HandlerFunc) error {
				gin.SetMode(gin.ReleaseMode)
				router := gin.New()
				router.Use(gin.Recovery())
				zipDir := router.Group("/dir_zip")
				zipDir.GET("*dir", fn)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/dir_zip/", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusOK, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 2",
			args: args{baseDir: func() string {
				dir, err := filepath.Abs("")
				if err != nil {
					return ""
				}
				return dir
			}()},
			testFunc: func(fn gin.HandlerFunc) error {
				gin.SetMode(gin.ReleaseMode)
				router := gin.New()
				router.Use(gin.Recovery())
				zipDir := router.Group("/dir_zip")
				zipDir.GET("*dir", fn)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/dir_zip/..", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusBadRequest, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 3",
			args: args{baseDir: func() string {
				dir, err := filepath.Abs("")
				if err != nil {
					return ""
				}
				return dir
			}()},
			testFunc: func(fn gin.HandlerFunc) error {
				gin.SetMode(gin.ReleaseMode)
				router := gin.New()
				router.Use(gin.Recovery())
				zipDir := router.Group("/dir_zip")
				zipDir.GET("*dir", fn)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/dir_zip/unknown_dir", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusInternalServerError, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 4",
			args: args{baseDir: func() string {
				dir, err := filepath.Abs("./..")
				if err != nil {
					return ""
				}
				return dir
			}()},
			testFunc: func(fn gin.HandlerFunc) error {
				gin.SetMode(gin.ReleaseMode)
				router := gin.New()
				router.Use(gin.Recovery())
				zipDir := router.Group("/dir_zip")
				zipDir.GET("*dir", fn)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/dir_zip/dist/", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusOK, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZipDir(tt.args.baseDir)
			if err := tt.testFunc(got); (err != nil) != tt.wantErr {
				t.Errorf("testFunc(HandlerFunc) error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
