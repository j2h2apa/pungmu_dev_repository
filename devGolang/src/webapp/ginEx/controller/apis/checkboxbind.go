package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]" binding:"required"`
}

type checkBoxBind struct {
	e *gin.Engine
}

var (
	// CheckBoxBinder :
	CheckBoxBinder *checkBoxBind
)

func init() {
	CheckBoxBinder = &checkBoxBind{}
}

func (cb *checkBoxBind) SetupEngine(e *gin.Engine) {
	cb.e = e
	cb.setupRouter()
}

func (cb *checkBoxBind) setupRouter() {
	cb.e.GET("/checkboxBindIn", func(c *gin.Context) {
		c.HTML(http.StatusOK, "checkbox_bind.tmpl", gin.H{
			"page": "binding check box",
		})
	})

	cb.e.POST("/checkboxBind", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(http.StatusOK, gin.H{
			"color": fakeForm.Colors,
		})
	})
}
