package main

import (
	"fmt"

	server "github.com/plan97/gofileserver"
	"github.com/plan97/gofileserver/config"
)

func main() {
	c := config.New()
	err := c.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = server.Serve(c)
	if err != nil {
		fmt.Println(err)
		return
	}
}
