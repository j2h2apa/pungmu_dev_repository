package apis

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//ProfileForm :
type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

//ProfilesForm :
type ProfilesForm struct {
	FirstName  string                  `form:"first_name" binding:"required"`
	FamilyName string                  `form:"family_name" binding:"required"`
	Avatar     []*multipart.FileHeader `form:"file" binding:"required"`
}

type fileUploader struct {
	e *gin.Engine
}

var (
	// FileUploader upload struct
	FileUploader *fileUploader
)

func init() {
	FileUploader = &fileUploader{}
}

// SetupEngine :
func (f *fileUploader) SetupEngine(e *gin.Engine) {
	f.e = e
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	f.e.MaxMultipartMemory = 8 << 20 // 8MiB
	f.e.Static("/files", "/upload")
	f.setupRouter()
}

func (f *fileUploader) setupRouter() {
	// single file upload
	f.e.GET("/uploadpage", func(c *gin.Context) {
		title := "upload single file"
		c.HTML(http.StatusOK, "uploadfile.html", gin.H{
			"page": title,
		})
	})
	f.e.POST("/upload", uploadSingle)

	// multy file upload
	f.e.GET("uploadpages", func(c *gin.Context) {
		title := "upload single files"
		c.HTML(http.StatusOK, "uploadfiles.tmpl", gin.H{
			"page": title,
		})
	})
	f.e.POST("/multiupload", uploadMultifile)

	// model binding of single file form
	f.e.POST("uploadbind", uploadbindSingle)
	// model binding of multifile upload
	f.e.POST("multiuploadbind", uploadbindMulty)
}

func uploadSingle(c *gin.Context) {
	// single file
	var file *multipart.FileHeader
	var err error

	file, err = c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err : %s", err.Error()))
		return
	}

	log.Println(file.Filename)

	// upload the file to specific dst.
	filename := filepath.Base(file.Filename)
	uploadPath := "upload/" + filename
	log.Println(uploadPath)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err : %s", err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"status":              "posted",
		"file name":           file.Filename,
		"header-content-type": c.GetHeader("content-Type"),
	})
}

func uploadMultifile(c *gin.Context) {
	firstName := c.PostForm("first_name")
	familyName := c.PostForm("family_name")

	// Multipart form
	var form *multipart.Form
	var err error
	form, err = c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form error : %s", err.Error()))
		return
	}

	var files []*multipart.FileHeader
	files = form.File["file"]
	log.Println(form)

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		filename := filepath.Base(file.Filename) // filename to except path
		uploadPath := "./upload/" + filename

		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file error : %s", err.Error()))
			return
		}
	}

	c.JSON(200, gin.H{
		"status":              "posted",
		"file list":           files,
		"file count":          len(files),
		"first name":          firstName,
		"family name":         familyName,
		"header-content-type": c.GetHeader("content-Type"),
	})
}

func uploadbindSingle(c *gin.Context) {
	// you can bind multipart form with explicit binding declaration:
	// c.ShouldBindWith(&form, binding.Form)
	// or you can simply use autobinding with ShouldBind method:
	var form ProfileForm
	// in this case proper binding will be automatically selected
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	filename := filepath.Base(form.Avatar.Filename)
	uploadPath := "upload/" + filename

	err := c.SaveUploadedFile(form.Avatar, uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":             "성공",
		"status":              "posted",
		"file name":           form.Avatar.Filename,
		"name":                form.Name,
		"header-content-type": c.GetHeader("content-Type"),
	})
}

func uploadbindMulty(c *gin.Context) {
	var form ProfilesForm

	if err := c.ShouldBindWith(&form, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		log.Fatal(err)
		return
	}

	for _, file := range form.Avatar {
		log.Println(file.Filename)
		// Upload the file to specific dst.
		filename := filepath.Base(file.Filename)
		uploadPath := "upload/" + filename

		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Println(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "posted",
		"file list":    form.Avatar,
		"file count":   len(form.Avatar),
		"first_name":   form.FirstName,
		"family_name":  form.FamilyName,
		"content-Type": c.GetHeader("content-Type"),
	})
}
