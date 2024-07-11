/**
 * @author ljn
 * @date 2024-03-21 17:42:32
 */

package main

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	group := r.Group("/tool")
	{
		group.POST("/upload", GinUploadImg)
		group.GET("/getBingPhoto", GetBingPhoto)
		group.GET("/backup_sql", func(context *gin.Context) {
			InitBackUp(Dns, OutPath)
			context.JSON(200, gin.H{
				"code": 200,
				"msg":  "success",
			})
		})
		group.GET("/export_md", func(context *gin.Context) {
			exportData(NewExportData())
			context.JSON(200, gin.H{
				"code": 200,
				"msg":  "success",
			})
		})
	}
}
