package main

import (
	c "Med/pkg/config"
	s "Med/pkg/server"
)

func main() {
	var conf c.Config
	c.LoadConfig(&conf, "server", false)
	if conf.Server.Protocol == "mllp" {
		s.ServeAndReportMllp(&conf.Server)
	}
}
