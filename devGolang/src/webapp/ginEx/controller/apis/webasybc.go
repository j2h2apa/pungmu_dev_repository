package apis

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type asyncEx struct {
	e *gin.Engine
}

var (
	// AsyncRouter : 
	AsyncRouter *asyncEx
)

func init() {
	AsyncRouter = &asyncEx{}
}

func (a *asyncEx) SetupEngine(e *gin.Engine) {
	a.e = e
	a.setupRouter()
}

func (a *asyncEx) setupRouter() {
	a.e.GET("/long_async", longAsync)
	a.e.GET("/long_sync", longSync)
}

func longSync(c *gin.Context) {
	time.Sleep(5 * time.Second)
	log.Println("Don't in path " + c.Request.URL.Path)
}

func longAsync(c *gin.Context) {
	var cpctx *gin.Context = c.Copy()

	go func() {
		time.Sleep(5*time.Second)
		log.Println("async handler Don't in path : " + cpctx.Request.URL.Path)
	}()
}