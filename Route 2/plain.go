package main

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.String(200, "Hello! This is Route 2")
}

func main() {
	router := gin.Default()

	router.GET("/route2", Test)

	router.Run(":8080")
}
