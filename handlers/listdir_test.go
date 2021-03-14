package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var errAssertionFailed = errors.New("assertion failed")

func TestListDirFiles(t *testing.T) {
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
				api := router.Group("/api")
				api.POST("list_dir_files", fn)

				w := httptest.NewRecorder()
				body := strings.NewReader(`{"dir":[]}`)
				req, _ := http.NewRequest(http.MethodPost, "/api/list_dir_files", body)
				req.Header.Set("Content-Type", "application/json")
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
				api := router.Group("/api")
				api.POST("list_dir_files", fn)

				w := httptest.NewRecorder()
				body := strings.NewReader(`{"dir":[".."]}`)
				req, _ := http.NewRequest(http.MethodPost, "/api/list_dir_files", body)
				req.Header.Set("Content-Type", "application/json")
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
				api := router.Group("/api")
				api.POST("list_dir_files", fn)

				w := httptest.NewRecorder()
				body := strings.NewReader(`{"dir":["non_existing_dir"]}`)
				req, _ := http.NewRequest(http.MethodPost, "/api/list_dir_files", body)
				req.Header.Set("Content-Type", "application/json")
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
				api := router.Group("/api")
				api.POST("list_dir_files", fn)

				w := httptest.NewRecorder()
				body := strings.NewReader(`[]`)
				req, _ := http.NewRequest(http.MethodPost, "/api/list_dir_files", body)
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusBadRequest, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListDirFiles(tt.args.baseDir)
			if err := tt.testFunc(got); (err != nil) != tt.wantErr {
				t.Errorf("testFunc(HandlerFunc) error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
