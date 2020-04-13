package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	var r multitemplate.Renderer = multitemplate.NewRenderer()

	articleLayouts, err := filepath.Glob(templatesDir + "/layouts/article-base.html")
	if err != nil {
		panic(err.Error())
	}
	log.Println(articleLayouts)

	articles, err := filepath.Glob(templatesDir + "/articles/*.html")
	if err != nil {
		panic(err.Error())
	}
	log.Println(articles)
	log.Println("-----------------------------------------------------")
	//Generate our templates map
	for _, article := range articles {
		layoutCopy := make([]string, len(articleLayouts))
		copy(layoutCopy, articleLayouts)
		files := append(layoutCopy, article)
		r.AddFromFiles(filepath.Base(article), files...)
		log.Println(files)
	}

	adminLayouts, err := filepath.Glob(templatesDir + "/layouts/admin-base.html")
	if err != nil {
		panic(err.Error())
	}

	admins, err := filepath.Glob(templatesDir + "/admins/*.html")
	if err != nil {
		panic(err.Error())
	}

	//Generate templates map
	for _, admin := range admins {
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		files := append(layoutCopy, admin)
		r.AddFromFiles(filepath.Base(admin), files...)
	}

	return r
}

func main() {
	var router *gin.Engine = gin.Default()
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "Welcome!",
		})
	})

	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article.html", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	router.Run(":8080")

}
