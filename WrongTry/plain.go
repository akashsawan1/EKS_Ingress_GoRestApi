package main

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.String(200, "Wrong Url! Bruh")
}

func main() {
	router := gin.Default()

	router.GET("/", Test)

	router.Run(":8080")
}
