/*
@Time : 2024/7/30 上午11:53
@Author : ljn
@File : zip_test
@Software: GoLand
*/

package main

import (
	"os"
	"testing"
)

func TestZipToMemory(t *testing.T) {
	buf, err := ZipToMemory("assets")
	if err != nil {
		t.Error(err)
		return
	}
	os.WriteFile("assets.zip", buf.Bytes(), 0644)
}

func TestUnzip(t *testing.T) {
}
