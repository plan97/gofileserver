package main

import (
	"fmt"

	server "github.com/plan97/gofileserver"
	"github.com/plan97/gofileserver/config"
)

func main() {
	conf := config.New()
	err := conf.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}

	router, err := server.Setup(conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	if conf.HTTPS {
		err = router.RunTLS(conf.Addr, conf.SSLCertFile, conf.SSLKeyFile)
	} else {
		err = router.Run(conf.Addr)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
