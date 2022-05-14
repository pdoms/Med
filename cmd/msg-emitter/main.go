package main

import (
	e "Med/internal/emitter"
	u "Med/internal/utils"
	"Med/pkg/config"
)

func main() {
	var conf config.Config
	err := config.LoadConfig(&conf, "emitter", false)
	u.HandleError(err)
	e.Emit(&conf.Emitter)
}
