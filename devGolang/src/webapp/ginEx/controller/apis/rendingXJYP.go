package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type renderBind struct {
	e *gin.Engine
}

var (
	//RenderBinder :
	RenderBinder *renderBind
)

func init() {
	RenderBinder = &renderBind{}
}

func (r *renderBind) SetupEngine(e *gin.Engine) {
	r.e = e
	r.setupRouter()
}

func (r *renderBind) setupRouter() {
	r.e.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":      "hey",
			"status":       http.StatusOK,
			"content-Type": c.GetHeader("content-Type"),
		})
	})

	r.e.GET("/moreJSON", func(c *gin.Context) {
		// also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		msg.Name, msg.Message, msg.Number = "Lena", "hey", 123
		c.JSON(http.StatusOK, msg)
	})

	r.e.GET("/someXML", func(c *gin.Context) {
		msg := struct {
			Message     string `xml:"message" json:"message"`
			Status      int    `xml:"status" json:"status"`
			RendInt     []int  `xml:"datas" json:"datas"`
			ContentType string `xml:"content-Type" json:"content-Type"`
		}{
			Message:     "hey",
			Status:      http.StatusOK,
			RendInt:     []int{1, 2, 3, 4, 5},
			ContentType: c.GetHeader("content-Type"),
		}

		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK, "data": msg})
	})

	r.e.GET("/someYAML", func(c *gin.Context) {
		// var msg struct {
		// 	Message     string `xml:"message" json:"message" yaml:"message"`
		// 	Status      int    `xml:"status" json:"status" yaml:"status"`
		// 	RendInt     []int  `xml:"data" json:"data" yaml:"data"`
		// 	ContentType string `xml:"content-Type" json:"content-Type" yaml:"content-Type"`
		// }

		var cntn struct {
			HeaderTitle string `yaml:"root"`
			Msg         struct {
				Message     string `xml:"message" json:"message" yaml:"message"`
				Status      int    `xml:"status" json:"status" yaml:"status"`
				RendInt     []int  `xml:"data" json:"data" yaml:"data"`
				ContentType string `xml:"content-Type" json:"content-Type" yaml:"content-Type"`
			} `yaml:"body"`
		}
		cntn.HeaderTitle = "yamlsample"
		cntn.Msg.Message = "hey"
		cntn.Msg.Status = http.StatusOK
		cntn.Msg.RendInt = []int{1, 2, 3, 4, 5}
		cntn.Msg.ContentType = c.GetHeader("content-Type")

		log.Println(cntn)

		c.YAML(http.StatusOK, cntn)
	})

	r.e.GET("/someProtoBuf", func(c *gin.Context) {
		var cntn struct {
			HeaderTitle string `protobuf="name=root" yaml:"root"`
			Msg         struct {
				Message     string `protobuf="name=message" xml:"message" json:"message" yaml:"message"`
				Status      int    `protobuf="name=status" xml:"status" json:"status" yaml:"status"`
				RendInt     []int  `protobuf="name=data" xml:"data" json:"data" yaml:"data"`
				ContentType string `protobuf="name=content-Type" xml:"content-Type" json:"content-Type" yaml:"content-Type"`
			} `protobuf="name=body" yaml:"body"`
		}
		cntn.HeaderTitle = "protosample"
		cntn.Msg.Message = "hey"
		cntn.Msg.Status = http.StatusOK
		cntn.Msg.RendInt = []int{1, 2, 3, 4, 5}
		cntn.Msg.ContentType = c.GetHeader("content-Type")

		c.ProtoBuf(http.StatusOK, cntn)
	})

	r.e.GET("/someSecureJSON", handlerSecureJSON)
	r.e.GET("/nonPureJSON", handlerPureJSON)
	r.e.GET("/pureJSON", handlerPureJSON)
}

//PureJSON
//Normally, JSON replaces special HTML characters with their unicode entities,
//e.g. < becomes \u003c. If you want to encode such characters literally,
//you can use PureJSON instead. This feature is unavailable in Go 1.6 and lower.
func handlerPureJSON(c *gin.Context) {
	opt := c.Query("options")

	if opt == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "non option",
		})
		return
	}

	if opt == "N" {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	} else if opt == "P" {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, worldd!</b>",
		})
	}
}

// json 하이재킹을 방지하기 위해 SecureJSON를 사용합니다.
// 주어진 구조체가 배열인 경우, 기본적으로 "while(1)," 이 응답 본문에 포함 됩니다.
func handlerSecureJSON(c *gin.Context) {
	var names []string = []string{"재희", "한희", "솔희"}
	c.SecureJSON(http.StatusOK, gin.H{
		"data": names,
	})
}

/*
	xml: unsupported type: struct
	{
		message string "xml:\"message\"";
		status int "xml:\"status\"";
		rendInt []int "xml:\"datas\"";
		contentType string "xml:\"contentType\""
	}
*/
