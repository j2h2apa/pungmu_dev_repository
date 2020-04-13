package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// Qstart :
	Qstart *qstart
)

func init() {
	Qstart = &qstart{}
}

type qstart struct {
	r *gin.Engine
}

// SetupRouter :
func (q *qstart) setupRouter() {
	q.r.GET("/someGet", q.someMethod)
	q.r.POST("/somePost", q.someMethod)
	q.r.PUT("/somePut", q.someMethod)
	q.r.DELETE("/someDelete", q.someMethod)
	q.r.PATCH("/somePatch", q.someMethod)
	q.r.HEAD("/someHead", q.someMethod)
	q.r.OPTIONS("/someOptions", q.someMethod)
}

func (q *qstart) someMethod(c *gin.Context) {
	var httpMethod string = c.Request.Method
	c.JSON(http.StatusOK, gin.H{"status": "good", "sending": httpMethod})
}

// SetupEngine :
func (q *qstart) SetupEngine(r *gin.Engine) {
	q.r = r
	q.setupRouter()
}
