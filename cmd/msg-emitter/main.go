package main

import (
	e "Med/internal/emitter"
	c "Med/pkg/config"
)

func main() {
	var conf c.Config
	c.LoadConfig(&conf, "emitter", false)
	e.Emit(&conf.Emitter)
	//fmt.Println(conf.Emitter.Connection.Host)

}
