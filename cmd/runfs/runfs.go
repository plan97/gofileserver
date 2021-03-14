package main

import (
	server "github.com/plan97/gofileserver"
	"github.com/plan97/gofileserver/config"
)

func runfs() error {
	conf := config.New()
	if err := conf.Fetch(); err != nil {
		return err
	}

	router, err := server.Setup(conf)
	if err != nil {
		return err
	}

	if conf.HTTPS {
		err = router.RunTLS(conf.Addr, conf.SSLCertFile, conf.SSLKeyFile)
	} else {
		err = router.Run(conf.Addr)
	}
	return err
}
