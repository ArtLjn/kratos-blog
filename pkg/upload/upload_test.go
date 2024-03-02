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
