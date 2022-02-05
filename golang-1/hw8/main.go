package main

import (
	"fmt"
	configuration "hw8/config/config"
)

func main() {
	conf := configuration.Configuration{}
	conf.Load()
	fmt.Print(conf)
}
