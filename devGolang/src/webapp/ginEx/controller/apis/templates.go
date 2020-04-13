package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type templateRouter struct {
	r *gin.Engine
}

var (
	//TemplateRouter :
	TemplateRouter *templateRouter
)

func init() {
	TemplateRouter = &templateRouter{}
}

func (t *templateRouter) SetupEngine(r *gin.Engine) {
	t.r = r
	t.setupRouter()
}

func (t *templateRouter) setupRouter() {
	//using templates with same name in different directory
	endpoint := t.r.Group("/templates")
	endpoint.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "apis/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	endpoint.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	//end using templates with same name in different directory

	// router.LoadHTMLFiles("./testdata/template/raw.tmpl") // already import in main
	endpoint.GET("/raw", customTemplateFunc)
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func customTemplateFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "raw.tmpl", gin.H{
		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	})
}
