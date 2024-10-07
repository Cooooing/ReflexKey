package main

import (
	"kernel/conf"
	"kernel/server"
	"kernel/sql"
	"kernel/util"
)

func main() {

	util.Boot()
	_ = sql.InitDatabase(false)
	conf.InitConf()
	server.Start()

	a := make(chan string)
	a <- "exit"
}
