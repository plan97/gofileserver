package main

import (
	"context"
	"net/http"

	server "github.com/plan97/gofileserver"
	"github.com/plan97/gofileserver/config"
)

func runfs(ctx context.Context) error {
	conf := config.New()
	if err := conf.Fetch(); err != nil {
		return err
	}

	router, err := server.Setup(conf)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    conf.Addr,
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		srv.Shutdown(ctx)
	}()

	if conf.HTTPS {
		return srv.ListenAndServeTLS(conf.SSLCertFile, conf.SSLKeyFile)
	}
	return srv.ListenAndServe()
}
