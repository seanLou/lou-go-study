package main

import "github.com/gin-gonic/gin"

func test(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "test",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello 关杨",
		})
	})

	router.POST("/test", test)

	router.Run(":8080")
}
