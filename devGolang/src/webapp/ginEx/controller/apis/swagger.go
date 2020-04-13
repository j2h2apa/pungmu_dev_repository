package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

type swago struct {
	e *gin.Engine
}

var (
	// Swago :
	Swago *swago
)

func init() {
	Swago = &swago{}
}

func (s *swago) SetupEngine(e *gin.Engine) {
	s.e = e
	s.setupRouter()
}

func (s *swago) setupRouter() {
	s.e.GET("/welcome/:name", s.welcomePathParam)
}

// Welcome godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @Description 자세한 설명은 이곳에 적습니다. - 2
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /welcome/{name} [get]
// @Success 200 {object} welcomeModel
func (s *swago) welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	message := name + " is very handsome!"
	println(message)
	welcomeMessage := &welcomeModel{1, message}

	c.JSON(http.StatusOK, gin.H{
		"message": welcomeMessage,
	})
}
