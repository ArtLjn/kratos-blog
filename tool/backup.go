package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jarvanstack/mysqldump"
	"gopkg.in/gomail.v2"
	"os"
	"path/filepath"
	"time"
)

// InitBackUp 备份数据库
func InitBackUp(dns, outPath string) {
	_, e := os.Stat(outPath)
	if e != nil && os.IsNotExist(e) {
		err := os.MkdirAll(outPath, os.ModePerm)
		if err != nil {
			return
		}
	}
	name := fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), "-hongDou.sql")
	filePath := filepath.Join(outPath, name)
	f, err := os.Create(filePath)

	_ = mysqldump.Dump(
		dns,
		mysqldump.WithData(),
		mysqldump.WithAllTable(),
		mysqldump.WithWriter(f),
	)
	if err != nil {
		log.Error("backup error: %v", err)
	}
	go SendBackupMail(filePath)
}

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

var Mail MailS

func SendBackupMail(filepath string) {
	m := gomail.NewMessage()
	m.SetHeader("From", Mail.Username)
	m.SetHeader("To", Mail.Username)
	m.SetHeader("Subject", "sql备份")
	m.SetBody("text/plain", "sql备份")
	// 添加附件
	// 注意：这里的 "/path/to/your/sql_backup.sql" 需要替换为实际的文件路径
	m.Attach(filepath)
	d := gomail.NewDialer(Mail.Host, Mail.Port, Mail.Username, Mail.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Error("send mail error: %v", err)
	}
}
