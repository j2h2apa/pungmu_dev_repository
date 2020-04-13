package apis

import (
	"net/http"
	"time"

	// _ "webapp/ginEx/docs"

	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

/*
form post로 넘어왔을 경우와 body json으로 넘어 왔을때 그리고 xml 데이터로
왔을때 binding되는 key를 적어 줍니다.
그리고 binding: "required"일 경우 request되는 데이터에 반드시 넘어와야 되며
없을 경우 error가 리턴됩니다.
*/

// User User model for request binding
type User struct {
	User     string `format:"string" form:"user" json:"user" xml:"user" binding:"required"`
	Password string `format:"string" form:"password" json:"password" xml:"password" binding:"required"`
	Nick     string `format:"string" form:"nick" json:"nick" xml:"nick" example:"nickname"`
	Age      uint16 `format:"uint16" form:"age" json:"age" xml:"age" example:"46" binding:"required,min=13,max=36"`
}

// UseMan model binding for request Querystring variable
type UseMan struct {
	User     string    `format:"string" form:"user" json:"user" xml:"user" binding:"required"`
	Password string    `format:"string" form:"password" json:"password" xml:"password" binding:"required"`
	Birthay  time.Time `format:"time.Time" form:"birthday" json:"birthday" xml:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type modelbind struct {
	e *gin.Engine
}

var (
	// ModelBind :
	ModelBind *modelbind
)

func init() {
	ModelBind = &modelbind{}
}

func (m *modelbind) SetupEngine(e *gin.Engine) {
	m.e = e
	m.setupRouter()
}

func (m *modelbind) setupRouter() {
	m.e.POST("/binding/json", m.bindingTestJSON)
	// Any 모든 method listen (handler는 한개)
	m.e.Any("/binding/form", m.bindingTestFORM)
	m.e.Any("/binding/shouldbindquery", m.bindingTestQUERY)
}

// @Summary querystring binding sample for request info
// @Description querystring binding model.
// @Accept  json
// @Produce  json
// @info get-string-by-string
// @Param body body UseMan true "user password birthday"
// @Router /binding/shouldbindquery [get]
// @Success 200 {object} UseMan
func (m *modelbind) bindingTestQUERY(c *gin.Context) {
	var user UseMan
	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "you are logged in",
		"info":         user,
		"content-Type": c.GetHeader("content-Type"),
	})
}

// @Summary json binding sample for request info
// @Description JSON binding model.
// @Accept  json
// @Produce  json
// @info get-string-by-string
// @Param body body User true "user password nick age"
// @Router /binding/json [post]
// @Success 200 {object} User
func (m *modelbind) bindingTestJSON(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if json.User != "j2h2apa" || json.Password != "j2h2s2apa" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              "you are logged in",
		"info":                json,
		"header-content-Type": c.GetHeader("content-Type"),
	})
}

/*
	binding은 method에 따라 달라집니다.
	ShouldBind() 함수는 GET일때와 POST일때가 다릅니다.
	GET method일 때는 Query String이 bind로 오고,
	POST method일 경우 header의 content-type를 확인해 Json인지, XML인지를 확인합니다.
	만약 Json도 XML도 아니라면 Form 데이터를 받습니다.
*/
// @Summary form binding sample for request info
// @Description FORM binding model.
// @Accept  json
// @Produce  json
// @info get-string-by-string
// @Param body body User true "user password nick age"
// @Router /binding/form [post]
// @Success 200 {object} User
func (m *modelbind) bindingTestFORM(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if form.User != "j2h2apa" && form.Password != "j2h2s2apa" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              "you are logged in form",
		"info":                form,
		"header-content-Type": c.GetHeader("content-Type"),
	})
}
