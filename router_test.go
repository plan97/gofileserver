package gofileserver

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/plan97/gofileserver/config"
)

var errAssertionFailed = errors.New("assertion failed")

func TestSetup(t *testing.T) {

	conf := config.New()
	conf.Fetch()
	conf.AllowCors = true

	type args struct {
		conf *config.Config
	}
	tests := []struct {
		name     string
		args     args
		testFunc func(*gin.Engine) error
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/noroute", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusTemporaryRedirect, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 2",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusOK, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 3",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/static/dist/go-file-server/favicon.ico", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusOK, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 4",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/dir_zip/dist", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusOK, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 5",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/dir_zip/unknown", nil)
				router.ServeHTTP(w, req)

				if !assert.IsEqual(http.StatusInternalServerError, w.Code) {
					return errAssertionFailed
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "test 6",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				body := strings.NewReader(`{"dir":[""]}`)
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
			name: "test 7",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
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
			name: "test 8",
			args: func() args {
				return args{
					conf,
				}
			}(),
			testFunc: func(router *gin.Engine) error {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/favicon.ico", nil)
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
			gotRouter, err := Setup(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup(Config) error = %v, wantErr %v", err, tt.wantErr)
			}
			if err = tt.testFunc(gotRouter); (err != nil) != tt.wantErr {
				t.Errorf("testFunc(Engine) error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
