package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type redirectsRouter struct {
	r *gin.Engine
}

var (
	//RedirectsRouter :
	RedirectsRouter *redirectsRouter
)

func init() {
	RedirectsRouter = &redirectsRouter{}
}

func (e *redirectsRouter) SetupEngine(r *gin.Engine) {
	e.r = r
	e.setupRouter()
}

func (e *redirectsRouter) setupRouter() {
	//case 1
	e.r.GET("/redirectCase1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com")
	})

	//case 2 issueing a router redirect
	e.r.GET("redirectCase2", func(c *gin.Context) {
		c.Request.URL.Path = "/redirectCase21"
		e.r.HandleContext(c)
	})
	e.r.GET("/redirectCase21", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
}
