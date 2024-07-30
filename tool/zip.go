/*
@Time : 2024/7/30 上午10:15
@Author : ljn
@File : zip
@Software: GoLand
*/

package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipToMemory(src string) (*bytes.Buffer, error) {
	// 创建一个内存缓冲区
	buf := new(bytes.Buffer)

	// 创建 zip.Writer
	zw := zip.NewWriter(buf)
	defer func() {
		if err := zw.Close(); err != nil {
			log.Error(err)
		}
	}()

	// 递归处理目录和文件
	err := filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) error {
		if errBack != nil {
			return errBack
		}

		// 通过文件信息，创建 zip 的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		// 替换文件信息中的文件名，使用 forward slashes for zip paths
		fh.Name = strings.TrimPrefix(strings.Replace(path, src, "", 1), string(filepath.Separator))

		// 如果是目录，名称后添加 /
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 Write 结构
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}

		// 如果不是标准文件就只写入头信息，不写入文件数据到 w
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		// 将文件内容拷贝到 zip.Writer
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func Unzip(src, out string) error {
	archive, err := zip.OpenReader(src)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("reader error %v", err)
	}
	defer archive.Close()
	for _, f := range archive.File {
		fp := filepath.Join(out, f.Name)
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(fp, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
			log.Error(err)
			return err
		}

		dstFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Error(err)
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			log.Error(err)
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			log.Error(err)
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return nil

}

func UnUploadZip(reader io.ReaderAt, size int64, out string) error {
	archive, err := zip.NewReader(reader, size)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("reader error %v", err)
	}

	for _, f := range archive.File {
		fp := filepath.Join(out, f.Name)
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(fp, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
			log.Error(err)
			return err
		}

		dstFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Error(err)
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			log.Error(err)
			return err
		}
		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			log.Error(err)
			return err
		}
		dstFile.Close()
		fileInArchive.Close()

	}
	return nil
}
