/*
@Time : 2024/7/30 上午10:21
@Author : ljn
@File : util
@Software: GoLand
*/

package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func InitPath(paths ...string) {
	for _, path := range paths {
		// 创建上传目录
		if !directoryExists(path) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		if e := checkPathAvailability(path); e != nil {
			panic(fmt.Errorf("Path is not available: %v\n", e))
		}
	}
}

func CopyFile(savePath string, src io.Reader) error {
	// 创建上传文件
	f, es := os.Create(savePath)
	if es != nil {
		return fmt.Errorf("create file error: %v\n", es)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	_, err := io.Copy(f, src)
	if err != nil {
		return err
	}
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
