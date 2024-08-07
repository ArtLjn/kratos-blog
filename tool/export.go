/*
@Time : 2024/7/11 下午12:49
@Author : ljn
@File : export 导出博客为md文件
@Software: GoLand
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"os"
	"path/filepath"
)

type ExportData struct {
	ExportPath string
}

type BlogMD struct {
	Title   string
	Preface string
	Content string
}

type ExportOption func(*ExportData)

func WithExportPath(path string) ExportOption {
	return func(data *ExportData) {
		data.ExportPath = path
	}
}

func NewExportData(options ...ExportOption) *ExportData {
	export := new(ExportData)
	for _, option := range options {
		option(export)
	}
	return export
}

func exportData(data *ExportData) {
	// 创建上传目录
	if len(data.ExportPath) == 0 {
		data.ExportPath = filepath.Join(Origin.U.UploadPath, "md")
	}
	if !directoryExists(data.ExportPath) {
		err := os.MkdirAll(data.ExportPath, os.ModePerm)
		if err != nil {
			return
		}
	}
	var blogs []BlogMD
	if err := GormDB.Model(BlogMD{}).Table("person_table").Find(&blogs).Error; err != nil {
		log.Errorf("query error %s", err)
		return
	}

	if len(blogs) != 0 {
		for _, blog := range blogs {
			savePath := filepath.Join(data.ExportPath, fmt.Sprintf("%s.%s", blog.Title, "md"))
			// 创建上传文件
			f, es := os.Create(savePath)
			if es != nil {
				log.Errorf("create file error %s", es)
				return
			}
			saveContent := fmt.Sprintf("%s\n%s", fmt.Sprintf("%s %s", "###", blog.Preface), blog.Content)
			_, err := io.Copy(f, bytes.NewBuffer([]byte(saveContent)))
			if err != nil {
				log.Errorf("save content error: %s", err)
				return
			}
		}
	}
}
