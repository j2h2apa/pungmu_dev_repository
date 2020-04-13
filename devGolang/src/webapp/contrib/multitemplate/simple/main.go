package main

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	var r multitemplate.Renderer = multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	var router *gin.Engine = gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})

	router.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	router.Run(":8080")
}
