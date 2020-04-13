package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"webapp/ginEx/controller/apis"
	_ "webapp/ginEx/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func getting(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func main() {

	// setup log file
	gin.DisableConsoleColor() // don't write on console
	var t time.Time = time.Now()
	startTime := t.Format("2006-01-02 15-04-05")
	logFile := "./log/ngleLog-" + startTime

	// Logging to a file
	path, e := os.Getwd()
	if e != nil {
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> %s\n", path)
	} else {
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>> %s\n", path)
	}

	f, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
	defer f.Close()
	// end of logging

	var r *gin.Engine = gin.Default()
	//sample of custom template function
	r.Delims("{[{", "}]}") // custom delimeter
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": apis.FormatAsDate,
	})
	r.LoadHTMLGlob("views/**/*")
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	// end of router log format

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/index", func(c *gin.Context) {
		log.Println("/index")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apis.Qstart.SetupEngine(r)
	apis.Prminpath.SetupEngine(r)
	apis.QryStr.SetupEngine(r)
	apis.FormPost.SetupEngine(r)
	apis.ParmConvMap.SetupEngine(r)
	apis.FileUploader.SetupEngine(r)
	apis.RouterGroup.SetupEngine(r)
	apis.Swago.SetupEngine(r)
	apis.ModelBind.SetupEngine(r)
	apis.URIBindRouter.SetupEngine(r)
	apis.AsyncRouter.SetupEngine(r)
	apis.HeaderBind.SetupEngine(r)
	apis.CheckBoxBinder.SetupEngine(r)
	apis.RenderBinder.SetupEngine(r)
	apis.ServingRouter.SetupEngine(r)
	apis.TemplateRouter.SetupEngine(r)
	apis.RedirectsRouter.SetupEngine(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
