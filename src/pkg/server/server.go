package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()

	r.GET("/", index)
	r.Run()
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "foo")
}