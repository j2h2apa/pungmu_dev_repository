package apis

import "github.com/gin-gonic/gin"

type queryString struct {
	r *gin.Engine
}

var (
	// QryStr :
	QryStr *queryString
)

func init() {
	QryStr = &queryString{}
}

func (q *queryString) setupRouter() {
	q.r.GET("/welcome", q.welcomeQueryString)
}

// handler
func (q *queryString) welcomeQueryString(c *gin.Context) {
	// QueryString 의 firstname 자체가 값이 없을 경우 Guest 셋팅
	firstname := c.DefaultQuery("firstname", "Guest")
	// shortcut for c.Request.URL.Query().Get("lastname")
	lastname := c.Query("lastname")

	c.JSON(200, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}

// SetupEngine :
func (q *queryString) SetupEngine(e *gin.Engine) {
	q.r = e
	q.setupRouter()
}
