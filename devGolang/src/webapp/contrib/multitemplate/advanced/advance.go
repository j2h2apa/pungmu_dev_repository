package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome!",
		})
	})

	router.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article.html", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	router.Run(":8080")
}

func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templateDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templateDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	//generate template map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}
