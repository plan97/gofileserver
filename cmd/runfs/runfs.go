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

	e := make(chan error, 1)

	go func() {
		if conf.HTTPS {
			e <- srv.ListenAndServeTLS(conf.SSLCertFile, conf.SSLKeyFile)
		} else {
			e <- srv.ListenAndServe()
		}
	}()

	select {
	case err = <-e:
		return err
	case <-ctx.Done():
		return srv.Shutdown(ctx)
	}
}
