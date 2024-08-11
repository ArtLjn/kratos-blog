/*
@Time : 2024/8/7 下午1:33
@Author : ljn
@File : mail
@Software: GoLand
*/

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/gomail.v2"
)

type MailS struct {
	Username string
	Host     string
	Port     int
	Password string
}

func NewMail(username, host string, port int, password string) MailS {
	return MailS{
		Username: username,
		Host:     host,
		Port:     port,
		Password: password,
	}
}

func SendBackupMail(filepath string) {
	m := gomail.NewMessage()
	m.SetHeader("From", Origin.Mail.Username)
	m.SetHeader("To", Origin.Mail.Username)
	m.SetHeader("Subject", "sql备份")
	m.SetBody("text/plain", "sql备份")
	// 添加附件
	// 注意：这里的 "/path/to/your/sql_backup.sql" 需要替换为实际的文件路径
	m.Attach(filepath)
	d := gomail.NewDialer(Origin.Mail.Host, Origin.Mail.Port, Origin.Mail.Username, Origin.Mail.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Errorf("send mail error: %v", err)
	}
}
