package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routerGroup struct {
	e *gin.Engine
}

var (
	// RouterGroup :
	RouterGroup *routerGroup
)

func init() {
	RouterGroup = &routerGroup{}
}

// SetupEngine
func (r *routerGroup) SetupEngine(e *gin.Engine) {
	r.e = e
	r.setupRouter()
}

func (r *routerGroup) setupRouter() {
	v1 := r.e.Group("/v1")
	{
		v1.POST("/login", r.loginEndpoint)
		v1.POST("/submit", r.submitEndpoint)
		v1.POST("/read", r.readEndpoint)
	}

	var v2 *gin.RouterGroup
	v2 = r.e.Group("/v2")
	{
		v2.POST("/login", r.loginEndpoint)
		v2.POST("/submit", r.submitEndpoint)
		v2.POST("/read", r.readEndpoint)
	}
}

func (r *routerGroup) loginEndpoint(c *gin.Context) {
	getParh := c.Request.URL.String()
	c.JSON(http.StatusOK, gin.H{
		"pathInfo": getParh,
	})
}

func (r *routerGroup) submitEndpoint(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(http.StatusOK, gin.H{
		"pathInfo": getPath,
	})
}

func (r *routerGroup) readEndpoint(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(http.StatusOK, gin.H{
		"pathInfo": getPath,
	})
}
