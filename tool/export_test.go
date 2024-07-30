/*
@Time : 2024/7/11 下午1:17
@Author : ljn
@File : export_test
@Software: GoLand
*/

package main

import (
	"testing"
)

func TestExportData(t *testing.T) {
	exportData(NewExportData(WithExportPath("")))
}
