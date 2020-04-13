package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//Logger : customer handler function
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//set example variable
		c.Set("example", "1233333")

		//before request

		c.Next()

		//after request
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	r.Run(":8080")
}
