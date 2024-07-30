/**
 * @author ljn
 * @date 2024-03-21 17:41:16
 * @dev 资源服务器上传文件使用
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gopkg.in/ini.v1"
	"kratos-blog/pkg/m_logs"
	"kratos-blog/pkg/server"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	Host               string
	Port               string
	u                  *UploadRepo
	FileSavaPath       string
	Current            CurrentBing
	Mail               MailS
	BackUpSqlSendEmail bool
)

func main() {
	m_logs.InitGinLog("tool")
	r := gin.Default()
	r.StaticFS("/img", http.Dir(FileSavaPath))
	UpdatePhoto()
	InitRouter(r)
	err := r.Run(fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		return
	}
}

func UpdatePhoto() {
	go func() {
		Current.mu.Lock()
		defer Current.mu.Unlock()
		Current.lastUpdateTime = time.Now()
		c := make(chan string, 1)
		err := BingPhoto(c)
		if err != nil {
			log.Println(err)
			return
		}
		Current.Url = <-c
		log.Println(Current.Url)
	}()
}

type ConfigBuilder interface {
	BuilderUpload() ConfigBuilder
	BuilderEmail() ConfigBuilder
	BuilderMysql() ConfigBuilder
	BuilderFilter() ConfigBuilder
	BuilderServer() ConfigBuilder
	BuilderCronTask() ConfigBuilder
	BuilderPath() ConfigBuilder
}

type BuilderConf struct {
	f *ini.File
}

func NewBuilderConf(f *ini.File) *BuilderConf {
	return &BuilderConf{f: f}
}

func (b *BuilderConf) BuilderUpload() ConfigBuilder {
	u = NewUploadRepo(
		b.f.Section("upload").Key("path").String(),
		b.f.Section("upload").Key("domain").String(),
		b.f.Section("upload").Key("maxsize").String(),
		b.f.Section("upload").Key("uri").String(),
	)
	return b
}

func (b *BuilderConf) BuilderEmail() ConfigBuilder {
	mailKey := b.f.Section("mail")
	port, _ := mailKey.Key("port").Int()
	Mail = NewMail(
		mailKey.Key("username").String(),
		mailKey.Key("host").String(),
		port,
		mailKey.Key("password").String(),
	)
	BackUpSqlSendEmail = b.f.Section("mysql").Key("backup_email").MustBool(false)
	return b
}

func (b *BuilderConf) BuilderMysql() ConfigBuilder {
	redSql := b.f.Section("mysql")
	NewDB(
		redSql.Key("username").String(),
		redSql.Key("password").String(),
		redSql.Key("host").String(),
		redSql.Key("port").String(),
		redSql.Key("database").String(),
		redSql.Key("charset").String(),
	)

	return b
}

func (b *BuilderConf) BuilderFilter() ConfigBuilder {
	F = NewFilter(
		strings.Split(b.f.Section("filter").Key("domain").String(), ","),
		b.f.Section("filter").Key("open").MustBool(false),
		strings.Split(b.f.Section("filter").Key("whiteList").String(), ","),
		b.f.Section("filter").Key("key").String(),
	)
	return b
}

func (b *BuilderConf) BuilderServer() ConfigBuilder {
	Host = b.f.Section("server").Key("host").String()
	Port = b.f.Section("server").Key("port").String()
	return b
}

func (b *BuilderConf) BuilderCronTask() ConfigBuilder {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {
		go UpdatePhoto()
	})
	OutPath = filepath.Join(b.f.Section("upload").Key("path").String(), "sql")
	backupCycle := b.f.Section("mysql").Key("backup_cycle").String()
	_, err = c.AddFunc("0 0 */"+backupCycle+" * *", func() {
		m_logs.CleanOldFile(0, OutPath)
		InitBackUp(Dns, OutPath)
	})
	if err != nil {
		return nil
	}
	c.Start()
	return b
}

func (b *BuilderConf) BuilderPath() ConfigBuilder {
	FileSavaPath = filepath.Join(u.UploadPath, "img")
	InitPath(u.UploadPath, FileSavaPath)
	return b
}

func init() {
	cfg, err := ini.Load(server.TooPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	NewBuilderConf(cfg).
		BuilderUpload().
		BuilderEmail().
		BuilderMysql().
		BuilderFilter().
		BuilderServer().
		BuilderCronTask().
		BuilderPath()
}
