package gofileserver

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/plan97/gofileserver/config"
	"github.com/plan97/gofileserver/handlers"
)

//go:embed dist/go-file-server/*
var content embed.FS

// Setup the router using the provided configuration.
func Setup(conf *config.Config) (router *gin.Engine, err error) {
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()
	if conf.AllowCors {
		router.Use(cors.Default())
	}

	subContent, err := fs.Sub(content, "dist/go-file-server")
	if err != nil {
		return nil, err
	}

	router.Use(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
			c.Next()
			return
		}

		reqPath := strings.TrimPrefix(c.Request.URL.Path, "/")
		if reqPath == "" {
			file, err := subContent.Open("index.html")
			if err != nil {
				c.Next()
				return
			}

			page, err := io.ReadAll(file)
			if err != nil {
				c.Next()
				return
			}

			c.Status(http.StatusOK)
			c.Writer.Write(page)
			c.Abort()
		}

		if _, err = subContent.Open(reqPath); err != nil {
			c.Next()
			return
		}

		http.FileServer(http.FS(subContent)).ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	static := router.Group("/static")
	static.Static("", conf.BaseDir)

	zipDir := router.Group("/dir_zip")
	zipDir.GET("*dir", handlers.ZipDir(conf.BaseDir))

	api := router.Group("/api")
	api.POST("list_dir_files", handlers.ListDirFiles(conf.BaseDir))

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})
	return
}
