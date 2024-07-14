package main

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.String(200, "Hello! This is Route 3")
}

func main() {
	router := gin.Default()

	router.GET("/route3", Test)

	router.Run(":8080")
}
