package main

import (
	"fmt"
)

func main() {
	err := runfs()
	if err != nil {
		fmt.Println(err)
	}
}
