package main

import (
	"hydrachat/configurator"
	"hydrachat/core"
)

func main() {
	config, _ := configurator.GetConfiguration("chat.conf")
	core.Run(config.RemoteAddr)
}
