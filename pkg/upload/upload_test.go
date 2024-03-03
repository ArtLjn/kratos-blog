package upload

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestUpload(t *testing.T) {
	po := UploadRepo{
		UploadPath: "./file",
		Domain:     "http://localhost:8080",
		MaxSize:    "3 MB",
	}
	router := gin.Default()
	router.POST("/util/upload", po.GinUploadImg)
	router.Run(":9080")
}

func TestGetBingPhoto(t *testing.T) {
	po := UploadRepo{
		UploadPath: "./file",
		Domain:     "http://localhost:8080",
		MaxSize:    "3 MB",
		Url:        "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1",
	}
	var uri string
	err := po.BingPhoto(&uri)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uri)
}
