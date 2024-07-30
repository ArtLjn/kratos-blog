/**
 * @author ljn
 * @date 2024-03-21 17:42:32
 */

package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	group := r.Group("/tool")
	{
		group.POST("/upload", GinUploadImg)
		group.GET("/getBingPhoto", GetBingPhoto)
		group.GET("/backup_sql", BackUpSql)
		group.GET("/export_md", ExportMD)
		group.GET("/backupAll", BackUpAllC)
		group.POST("/zipToRecover", ZipToRecover)
	}
}
