/*
@Time : 2024/7/11 下午1:24
@Author : ljn
@File : backup_test
@Software: GoLand
*/

package main

import "testing"

func TestBackUp(t *testing.T) {
	InitBackUp(Dns, OutPath)
}
