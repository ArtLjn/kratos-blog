/**
 * @author ljn
 * @date 2024-03-21 17:43:12
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thedevsaddam/gojsonq"
	"gopkg.in/ini.v1"
	"io"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	"log"
	h2 "net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var u *UploadRepo

func NewUploadRepo(data ...string) *UploadRepo {
	return &UploadRepo{UploadPath: data[0], Domain: data[1], MaxSize: data[2], Url: data[3]}
}

func init() {
	cfg, err := ini.Load("tool/config.ini")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	u = NewUploadRepo(
		cfg.Section("upload").Key("path").String(),
		cfg.Section("upload").Key("domain").String(),
		cfg.Section("upload").Key("maxsize").String(),
		cfg.Section("upload").Key("uri").String(),
	)
}

type UploadRepo struct {
	UploadPath string //上传路径
	Domain     string // 域名
	MaxSize    string // 最大文件大小（字节）
	Url        string // 外部uri
}

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
	savePath := filepath.Join(u.UploadPath, newFileName)
	getLastIndex(&u.Domain)
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(vo.BAD_REQUEST, res.SetCode(vo.BAD_REQUEST).SetInfo("上传文件失败").GetCodeInfo())
		return
	}
	uri := fmt.Sprintf("%s%s", u.Domain, newFileName)
	ctx.JSON(vo.SUCCESS_REQUEST, res.SetCode(vo.SUCCESS_REQUEST).SetInfo("上传成功").SetData(uri).GetCodeInfo())
}

// CurrentBing :dev 今日Bing图片
type CurrentBing struct {
	Url            string
	lastUpdateTime time.Time
	mu             sync.Mutex
}

var Current CurrentBing

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
		newImgName := fmt.Sprintf("%s%s", uuid.New().String()[:8], ".jpg")
		savePath := filepath.Join(u.UploadPath, newImgName)
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

func CopyFile(savePath string, src io.Reader) error {
	// 创建上传目录
	if !directoryExists(u.UploadPath) {
		os.MkdirAll(u.UploadPath, os.ModePerm)
	}
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
	io.Copy(f, src)
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

	// Use regular expression to match the number and unit.
	re := regexp.MustCompile(`^(\d+)([A-Za-z]+)$`)
	matches := re.FindStringSubmatch(sizeStr)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid size format: %s", sizeStr)
	}

	// Parse the number as an int64.
	size, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, err
	}

	// Convert the unit to bytes.
	var bytes int64
	switch strings.ToUpper(matches[2]) {
	case "B":
		bytes = size
	case "KB":
		bytes = size * 1024
	case "MB":
		bytes = size * 1024 * 1024
	case "GB":
		bytes = size * 1024 * 1024 * 1024
	default:
		return 0, fmt.Errorf("unsupported size unit: %s", matches[2])
	}

	return bytes, nil
}
