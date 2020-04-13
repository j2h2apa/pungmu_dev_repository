package apis

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type servingRouter struct {
	router *gin.Engine
}

var (
	//ServingRouter :
	ServingRouter *servingRouter
)

func init() {
	ServingRouter = &servingRouter{}
}

func (s *servingRouter) SetupEngine(e *gin.Engine) {
	s.router = e
	s.setupRouter()
}

func (s *servingRouter) setupRouter() {
	// Serving static files
	s.router.Static("/assets", "./assets")
	s.router.StaticFS("/more_static", http.Dir("./my_file_system"))
	s.router.StaticFile("/solhee", "./static_file/solhee.jpg")

	// Serving data from file
	s.router.GET("/local/file", func(c *gin.Context) {
		c.File("./servingfile/file.go")
	})

	// Serving data from file case 2
	endpoint := s.router.Group("/diagnose")
	var fs http.FileSystem = http.Dir("./servingfile")
	endpoint.GET("/servingfilefs", func(c *gin.Context) {
		c.FileFromFS("file.go", fs)
	})

	// Serving data from reader
	endpoint.GET("/someDataFromReader", servingReader)
}

func servingReader(c *gin.Context) {
	log.Println("servingReader")
	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	var reader io.ReadCloser = response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	var extraHeaders map[string]string = map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
