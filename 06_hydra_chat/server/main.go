package main

import (
	"hydrachat/configurator"
	"hydrachat/core"
)

func main() {
	configuration := configurator.Configuration{}
	err := configuration.GetConfiguration("chat.conf")
	if err != nil {
		panic(err)
	}
	core.Run(configuration.RemoteAddr)
}
