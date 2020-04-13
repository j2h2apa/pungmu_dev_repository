package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// ParmConvMap ParmConvMap
	ParmConvMap *parmConvMap
)

type parmConvMap struct {
	e *gin.Engine
}

func init() {
	ParmConvMap = &parmConvMap{}
}

// SetupEngine releation gin engine of main
func (p *parmConvMap) SetupEngine(e *gin.Engine) {
	p.e = e
	p.setupRouter()
}

func (p *parmConvMap) setupRouter() {
	title := "map type query string"
	p.e.GET("/note", func(c *gin.Context) {
		c.HTML(http.StatusOK, "parm_conv_map.tmpl", gin.H{
			"page": title,
		})
	})

	p.e.POST("/postMap", p.PostMap)
}

// PostMap web handler of /postMap
func (p *parmConvMap) PostMap(c *gin.Context) {
	var ids map[string]string = c.QueryMap("ids")
	var names map[string]string = c.PostFormMap("names")

	headerType := c.GetHeader("content-Type")

	c.JSON(http.StatusOK, gin.H{
		"status":            "posted",
		"ids":               ids,
		"names":             names,
		"head-content-type": headerType,
	})
}
