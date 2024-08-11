/*
@Time : 2024/8/7 上午11:10
@Author : ljn
@File : setting_test
@Software: GoLand
*/

package main

import (
	"fmt"
	"kratos-blog/pkg/conf"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	CreateConfig(conf.Config{
		// ... 配置初始化
		Version: "1.0",
		Open:    false,
		Upload: struct {
			Path    string
			Domain  string
			MaxSize string
			BingUrl string
		}{Path: Origin.U.UploadPath, Domain: Origin.U.Domain, MaxSize: Origin.U.MaxSize, BingUrl: Origin.U.Url},
		BackUp: struct {
			Cycle     int
			OpenEmail bool
		}{Cycle: Origin.BackUp.BackUpCycle, OpenEmail: Origin.BackUp.BackUpSqlSendEmail},
		QQEmail: struct {
			Username string
			Password string
			Host     string
			Port     int
		}{Username: Origin.Mail.Username, Password: Origin.Mail.Password, Host: Origin.Mail.Host, Port: Origin.Mail.Port},
	})
}

func TestFindAllConfig(t *testing.T) {
	fmt.Print(FindAllConfig())
	_, err := FindConfig("1.0")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(HasVersion("1.0"))
}

func TestDeleteConfigByVersion(t *testing.T) {
	DeleteConfigByVersion("1.0")
}

func TestUpdateConfigByVersion(t *testing.T) {
	UpdateConfigByVersion(conf.Config{
		Version: "1.1",
		Open:    false,
		Upload: struct {
			Path    string
			Domain  string
			MaxSize string
			BingUrl string
		}{Path: "test", Domain: "test", MaxSize: "test", BingUrl: "test"},
		BackUp: struct {
			Cycle     int
			OpenEmail bool
		}{Cycle: 1, OpenEmail: false},
		QQEmail: struct {
			Username string
			Password string
			Host     string
			Port     int
		}{
			Username: "test",
			Password: "test",
			Host:     "test",
			Port:     1,
		},
	})
}

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("1.2")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetCurrentConfig(t *testing.T) {
	fmt.Print(GetCurrentConfig())
}
