package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

type headerBind struct {
	e *gin.Engine
}

//HeaderBind :
var HeaderBind *headerBind

func init() {
	HeaderBind = &headerBind{}
}

func (h *headerBind) SetupEngine(e *gin.Engine) {
	h.e = e
	h.setupRouter()
}

func (h *headerBind) setupRouter() {
	h.e.GET("/getHeader", func(c *gin.Context) {
		lh := testHeader{}

		if err := c.ShouldBindHeader(&lh); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Printf("%#v\n", lh)
		c.JSON(http.StatusOK, gin.H{
			"Rate":   lh.Rate,
			"Domain": lh.Domain,
		})
	})
}
