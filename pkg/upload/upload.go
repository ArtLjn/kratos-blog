package upload

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"io"
	"kratos-blog/pkg/vo"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type UploadRepo struct {
	UploadPath string //上传路径
	Domain     string // 域名
	MaxSize    string // 最大文件大小（字节）
}

// GinUploadImg :dev 使用gin框架进行文件上传
func (u *UploadRepo) GinUploadImg(ctx *gin.Context) {
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
	savePath := filepath.Join(u.UploadPath, newFileName)
	getLastIndex(&u.Domain)
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(vo.BAD_REQUEST, res.SetCode(vo.BAD_REQUEST).SetInfo("上传文件失败").GetCodeInfo())
		return
	}
	uri := fmt.Sprintf("%s%s", u.Domain, newFileName)
	ctx.JSON(vo.SUCCESS_REQUEST, res.SetCode(vo.SUCCESS_REQUEST).SetInfo("上传成功").SetData(uri).GetCodeInfo())
}

// UploadImg kratos 通用上传文件
func (u *UploadRepo) UploadImg(ctx context.Context, uri *string) error {
	req, ok := http.RequestFromServerContext(ctx)
	if !ok {
		return fmt.Errorf("%v\n", "解析context异常")
	} else if !strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") {
		return fmt.Errorf("Invalid Content-Type: %v\n", req.Header.Get("Content-Type"))
	}
	size, _ := ParseSize(u.MaxSize)
	file, handler, err := req.FormFile("file")
	if err != nil {
		return err
	} else if handler.Size > size {
		return fmt.Errorf("%v%s\n", "上传文件大小不的大于", u.MaxSize)
	}
	defer file.Close()
	originName := handler.Filename
	lastDotIndex := strings.LastIndex(originName, ".")
	if lastDotIndex == -1 {
		return fmt.Errorf("%v\n", "文件后缀获取失败")
	}
	// 提取后缀名
	extension := originName[lastDotIndex+1:]
	newFileName := fmt.Sprintf("%s%s%s", uuid.New().String()[:8], ".", extension)
	savePath := filepath.Join(u.UploadPath, newFileName)
	getLastIndex(&u.Domain)
	// 创建上传目录
	if !directoryExists(u.UploadPath) {
		os.MkdirAll(u.UploadPath, os.ModePerm)
	}
	if e := checkPathAvailability(u.UploadPath); e != nil {
		return fmt.Errorf("Path is not available: %v\n", e)
	}
	// 创建上传文件
	f, es := os.Create(savePath)
	if es != nil {
		return fmt.Errorf("create file error: %v\n", es)
	}
	defer f.Close()
	io.Copy(f, file)
	*uri = fmt.Sprintf("%s%s", u.Domain, newFileName)
	return nil
}

func getLastIndex(uri *string) {
	n := len(*uri)
	lastIndex := (*uri)[n-1 : n]
	if lastIndex != "/" {
		*uri = fmt.Sprintf("%s%s", *uri, "/")
	}
}

// 检查路径是否可用
func checkPathAvailability(path string) error {
	// 尝试在给定路径下创建一个临时文件
	f, err := os.CreateTemp(path, "temple.*")
	if err != nil {
		return err // 返回错误，路径不可用
	}
	defer os.Remove(f.Name()) // 删除临时文件
	f.Close()                 // 关闭文件句柄
	return nil
}

func directoryExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return fi.IsDir()
}

// ParseSize parses a size string like "10MB" and returns the number of bytes.
func ParseSize(sizeStr string) (int64, error) {
	// Remove any potential spaces from the string.
	sizeStr = strings.TrimSpace(sizeStr)

	// Split the number and the unit.
	parts := strings.Split(sizeStr, " ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid size format: %s", sizeStr)
	}

	// Parse the number as an int64.
	size, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	// Convert the unit to bytes.
	var bytes int64
	switch strings.ToUpper(parts[1]) {
	case "B":
		bytes = size
	case "KB":
		bytes = size * 1024
	case "MB":
		bytes = size * 1024 * 1024
	case "GB":
		bytes = size * 1024 * 1024 * 1024
	default:
		return 0, fmt.Errorf("unsupported size unit: %s", parts[1])
	}

	return bytes, nil
}
