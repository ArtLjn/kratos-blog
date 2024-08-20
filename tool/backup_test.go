/*
@Time : 2024/7/11 下午1:24
@Author : ljn
@File : backup_test
@Software: GoLand
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetSqlList(t *testing.T) {
	err := filepath.Walk(filepath.Join(Origin.U.UploadPath, Origin.Prefix[1]), visit)
	if err != nil {
		fmt.Println(err)
	}
}

// visit 函数用于对每个文件进行操作
func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !info.IsDir() {
		fmt.Println(path) // 打印文件路径
	}
	return nil
}
