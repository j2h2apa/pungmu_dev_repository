package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MemberURIBind struct of uri bind testing
type MemberURIBind struct {
	Name string `format:"string" uri:"uriname" binding:"required"`
	Age  uint16 `format:"uint16" uri:"uriage" binding:"required"`
	ID   string `formag:"string" uri:"uriid" binding:"required"`
}

type uRIBindRouter struct {
	e *gin.Engine
}

var (
	// URIBindRouter :
	URIBindRouter *uRIBindRouter
)

func init() {
	URIBindRouter = &uRIBindRouter{}
}

// SetupGin setup web framework
func (u *uRIBindRouter) SetupEngine(e *gin.Engine) {
	u.e = e
	u.setupRouter()
}

func (u *uRIBindRouter) setupRouter() {
	u.e.GET("/name/:uriname/age/:uriage/id/:uriid", func(c *gin.Context) {
		var user MemberURIBind
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":         user.Name,
			"age":          user.Age,
			"id":           user.ID,
			"content-Type": c.GetHeader("content-Type"),
		})
	})
}
