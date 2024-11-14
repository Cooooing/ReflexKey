package main

import (
	"kernel/conf"
	"kernel/server"
	"kernel/sql"
)

func main() {

	conf.Boot()
	_ = sql.InitDatabase(false)
	conf.InitConf()
	server.Start()

	a := make(chan string)
	a <- "exit"
}
