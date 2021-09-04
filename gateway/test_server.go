package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/abb", func(c *gin.Context) {
		c.String(200, "url abb")
	})
	r.GET("/bba", func(c *gin.Context) {
		c.String(200, "url bba")
	})
	r.Run(":2003")
}
