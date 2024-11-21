package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func static(c *gin.Context) {
	staticServer := http.FileServer(http.FS(frontend))
	staticServer.ServeHTTP(c.Writer, c.Request)
}

func index(c *gin.Context) {
	file, err := frontend.ReadFile("static/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading index.html")
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", file)
}

func noRoute(c *gin.Context) {
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/api/") {
		notFound(c)
		return
	}

	c.Redirect(http.StatusFound, "/")
}
