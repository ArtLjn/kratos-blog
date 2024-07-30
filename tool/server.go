/**
 * @author ljn
 * @date 2024-03-21 17:43:12
 */

package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/thedevsaddam/gojsonq"
	"io"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	h2 "net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// GinUploadImg :dev 使用gin框架进行文件上传
func GinUploadImg(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	size, _ := ParseSize(u.MaxSize)
	res := vo.Response{}
	if err != nil {
		ctx.JSON(vo.BAD_REQUEST, res.SetCode(vo.BAD_REQUEST).SetInfo(vo.REQUIRE_FILE_ERROR).GetCodeInfo())
		return
	} else if file.Size > size {
		info := fmt.Sprintf("上传文件大小不的大于%s", u.MaxSize)
		ctx.JSON(vo.BAD_REQUEST, res.SetCode(vo.BAD_REQUEST).SetInfo(info).GetCodeInfo())
		return
	}
	originName := file.Filename
	lastDotIndex := strings.LastIndex(originName, ".")
	if lastDotIndex == -1 {
		ctx.JSON(vo.BAD_REQUEST, res.SetCode(vo.BAD_REQUEST).SetInfo("文件后缀获取失败").GetCodeInfo())
		return
	}
	// 提取后缀名
	extension := originName[lastDotIndex+1:]
	newFileName := fmt.Sprintf("%s%s%s", uuid.New().String()[:8], ".", extension)
	savePath := filepath.Join(FileSavaPath, newFileName)
	getLastIndex(&u.Domain)
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(vo.BAD_REQUEST, res.SetCode(vo.BAD_REQUEST).SetInfo("上传文件失败").GetCodeInfo())
		return
	}
	uri := fmt.Sprintf("%s%s", u.Domain, newFileName)
	ctx.JSON(vo.SUCCESS_REQUEST, res.SetCode(vo.SUCCESS_REQUEST).SetInfo("上传成功").SetData(uri).GetCodeInfo())
}

// GetBingPhoto :dev GetBingPhoto 获取Bing每日随机图片
func GetBingPhoto(ctx *gin.Context) {
	res := vo.Response{}
	ctx.JSON(vo.SUCCESS_REQUEST, res.SetCode(vo.SUCCESS_REQUEST).SetInfo(vo.QUERY_SUCCESS).SetData(Current.Url).GetCodeInfo())
}

// BingPhoto :dev 获取bing每日图片接口
func BingPhoto(uri chan string) error {
	body, err := util.Request(h2.MethodGet, u.Url, nil)
	if err != nil {
		return err
	}
	images, ok := gojsonq.New().JSONString(body).Find("images").([]interface{})
	if len(images) == 0 || !ok {
		return fmt.Errorf("require images failed \n")
	}
	imageUri := images[0].(map[string]interface{})
	img, ok := imageUri["url"].(string)
	if ok {
		fullUri := fmt.Sprintf("%s%s", "https://cn.bing.com", img)
		res, e := h2.Get(fullUri)
		if e != nil {
			return e
		} else if res.StatusCode != h2.StatusOK {
			return fmt.Errorf("The request failure status code is %d\n", res.StatusCode)
		}
		defer res.Body.Close()
		newImgName := fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), "_bing.jpg")
		savePath := filepath.Join(FileSavaPath, newImgName)
		if ue := CopyFile(savePath, res.Body); ue != nil {
			return ue
		}
		getLastIndex(&u.Domain)
		uri <- fmt.Sprintf("%s%s", u.Domain, newImgName)
		close(uri)
		return nil
	}
	return fmt.Errorf("The image URI parsing failed \n")
}

func BackUpSql(ctx *gin.Context) {
	InitBackUp(Dns, OutPath)
	if GetNeedDownload(ctx) {
		HttpDownload(ctx, filepath.Join(u.UploadPath, "sql"), "sql")
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func ExportMD(ctx *gin.Context) {
	exportData(NewExportData())
	if GetNeedDownload(ctx) {
		HttpDownload(ctx, filepath.Join(u.UploadPath, "md"), "md")
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func BackUpAllC(ctx *gin.Context) {
	BackUpAll()
	if GetNeedDownload(ctx) {
		HttpDownload(ctx, u.UploadPath, "backup")
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetNeedDownload(ctx *gin.Context) bool {
	needDownload := ctx.Query("download")
	need, err := strconv.ParseBool(needDownload)
	if err != nil {
		log.Error("needDownload must be bool")
		return false
	}
	return need
}

func HttpDownload(ctx *gin.Context, path, zipName string) {
	buf, err := ZipToMemory(path)
	if err != nil {
		ctx.String(h2.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
		return
	}

	// 设置响应头
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", zipName))
	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Length", fmt.Sprintf("%d", buf.Len()))

	// 将缓冲区中的数据写入响应
	_, err = io.Copy(ctx.Writer, buf)
	if err != nil {
		ctx.String(h2.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
		return
	}
}

func ZipToRecover(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(h2.StatusBadRequest, gin.H{"code": 400, "msg": "No file is received"})
		return
	}
	filePrefix := strings.Split(file.Filename, ".")
	if len(filePrefix) < 2 {
		ctx.JSON(h2.StatusBadRequest, gin.H{"code": 400, "msg": "No file is received"})
		return
	}
	if filePrefix[len(filePrefix)-1] != "zip" {
		ctx.JSON(h2.StatusBadRequest, gin.H{"code": 400, "msg": "Unsupported file types"})
		return
	}
	uploadedFile, err := file.Open()
	if err != nil {
		ctx.JSON(h2.StatusInternalServerError, gin.H{"code": 400, "msg": "Failed to open the uploaded file"})
		return
	}
	defer uploadedFile.Close()

	// Create a buffer to read the uploaded file into
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, uploadedFile); err != nil {
		ctx.JSON(h2.StatusInternalServerError, gin.H{"code": 400, "msg": "Failed to read the uploaded file"})
		return
	}

	// Unzip the file directly from the buffer
	if err := UnUploadZip(bytes.NewReader(buf.Bytes()), int64(buf.Len()), u.UploadPath); err != nil {
		ctx.JSON(h2.StatusInternalServerError, gin.H{"code": 400, "msg": "Failed to unzip the file"})
		return
	}

	ctx.JSON(h2.StatusOK, gin.H{"code": 200, "msg": "File unzipped successfully"})
}
