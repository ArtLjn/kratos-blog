package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jarvanstack/mysqldump"
	"os"
	"path/filepath"
	"time"
)

// InitBackUp 备份数据库
func InitBackUp(dns, outPath string, sendEmail bool) {
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
		log.Errorf("backup error: %v", err)
		return
	}
	if sendEmail {
		go SendBackupMail(filePath)
	}
}

func BackUpAll() {
	InitBackUp(Origin.BackUp.Dns, filepath.Join(Origin.U.UploadPath, Origin.Prefix[0]), Origin.BackUp.BackUpSqlSendEmail)
	exportData(NewExportData())
}
