package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type formPost struct {
	r *gin.Engine
}

var (
	// FormPost :
	FormPost *formPost
)

func init() {
	FormPost = &formPost{}
}

func (f *formPost) setupRouter() {
	FormPost.r.GET("/form_input", func(c *gin.Context) {
		c.HTML(200, "formdata.tmpl", nil)
	})
	FormPost.r.POST("/form_post", FormPost.formPostHandler)

	FormPost.r.GET("/form_post_with_querystring_in", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uformdata.tmpl", nil)
	})
	FormPost.r.POST("/form_post_with_querystring", FormPost.uformPostHandler)
}

// SetupEngine :
func (f *formPost) SetupEngine(e *gin.Engine) {
	FormPost.r = e
	FormPost.setupRouter()
}

func (f *formPost) formPostHandler(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	headerType := c.GetHeader("content-Type")

	c.JSON(http.StatusOK, gin.H{
		"status":              "post",
		"message":             message,
		"nick":                nick,
		"header-content-type": headerType,
	})
}

// QueryString의 parameter와 Form tag의 post 데이터를 같이 받으려면 header의 Content-Type은
// "application/x-www-form-urlencoded"여야 합니다.
func (f *formPost) uformPostHandler(c *gin.Context) {
	id := c.Query("id")
	strPage := c.DefaultQuery("page", "0")
	nPage, _ := strconv.Atoi(strPage)

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	headerType := c.GetHeader("content-Type")

	c.JSON(http.StatusOK, gin.H{
		"status":              "post",
		"message":             message,
		"nick":                nick,
		"header-content-type": headerType,
		"id":                  id,
		"page":                nPage,
	})
}
