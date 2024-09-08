package main

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/server"
	"kernel/sql"
	"kernel/util"
)

func main() {

	util.Boot()
	_ = sql.InitDatabase(false)
	model.InitConf()
	server.Start()

	engine := gin.Default()
	engine.GET("/start", func(c *gin.Context) {
		server.Start()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.GET("/stop", func(c *gin.Context) {
		server.Stop()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.GET("/restart", func(c *gin.Context) {
		server.Restart()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.Run(":8080")

	a := make(chan string)
	a <- "exit"
}
