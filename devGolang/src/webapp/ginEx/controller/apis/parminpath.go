package apis

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	// Prminpath :
	Prminpath *Paraminpath
)

func init() {
	Prminpath = &Paraminpath{}
}

// Paraminpath :
type Paraminpath struct {
	r *gin.Engine
}

// SetupEngine :
func (p *Paraminpath) SetupEngine(r *gin.Engine) {
	p.r = r
	p.setuprouter()
}

func (p *Paraminpath) setuprouter() {
	p.r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	p.r.GET("/user/:name/age/:old", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    c.Param("name"),
			"action":  c.Param("old"),
			"message": c.Param("name") + " is " + c.Param("old") + " years old.",
		})
	})

	p.r.GET("/color/:color/*fruits", func(c *gin.Context) {
		color := c.Param("color")
		fruits := c.Param("fruits")
		var fruitArray []string = strings.Split(fruits, "/")
		fruitArray = append(fruitArray[:0], fruitArray[1:]...)
		c.JSON(http.StatusOK, gin.H{
			"color":  color,
			"fruits": fruitArray,
		})
	})

	p.r.POST("/user/:name/*action", func(c *gin.Context) {
		same := c.FullPath() == "/user/:name/*action"
		strsame := strconv.FormatBool(same)
		c.JSON(http.StatusOK, gin.H{
			"FullPath": c.FullPath(),
			"is_same":  strsame,
		})
	})
}
