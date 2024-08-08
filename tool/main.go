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
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"kratos-blog/pkg/db"
	"kratos-blog/pkg/m_logs"
	"kratos-blog/pkg/server"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type OriginConfig struct {
	Server struct {
		Host string
		Port string
	}
	BackUp struct {
		Dns                string
		BackUpCycle        int
		BackUpSqlSendEmail bool
	}
	Admin struct {
		Username string
		Password string
	}
	Mongo   string
	F       *Filter
	Prefix  []string
	U       *UploadRepo
	Current CurrentBing
	Mail    MailS
}

var (
	GormDB         *gorm.DB
	MCli           *mongo.Client
	ConfCollection *mongo.Collection
	Origin         *OriginConfig
)

func NewOriginConfig() *OriginConfig {
	cfg, err := ini.Load(server.TooPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return &OriginConfig{
		Server: struct {
			Host string
			Port string
		}{
			Host: cfg.Section("server").Key("host").String(),
			Port: cfg.Section("server").Key("port").String(),
		},
		U: &UploadRepo{
			cfg.Section("upload").Key("path").String(),
			cfg.Section("upload").Key("domain").String(),
			cfg.Section("upload").Key("maxsize").String(),
			cfg.Section("upload").Key("uri").String(),
		},
		Mail: NewMail(
			cfg.Section("mail").Key("username").String(),
			cfg.Section("mail").Key("host").String(),
			cfg.Section("mail").Key("port").MustInt(),
			cfg.Section("mail").Key("password").String(),
		),
		F: NewFilter(
			strings.Split(cfg.Section("filter").Key("domain").String(), ","),
			cfg.Section("filter").Key("open").MustBool(false),
			strings.Split(cfg.Section("filter").Key("whiteList").String(), ","),
			cfg.Section("filter").Key("key").String(),
		),
		BackUp: struct {
			Dns                string
			BackUpCycle        int
			BackUpSqlSendEmail bool
		}{
			Dns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
				cfg.Section("mysql").Key("username").String(),
				cfg.Section("mysql").Key("password").String(),
				cfg.Section("mysql").Key("host").String(),
				cfg.Section("mysql").Key("port").String(),
				cfg.Section("mysql").Key("database").String(),
				cfg.Section("mysql").Key("charset").String(),
			),
			BackUpCycle:        cfg.Section("mysql").Key("backup_cycle").MustInt(),
			BackUpSqlSendEmail: cfg.Section("mysql").Key("backup_email").MustBool(false),
		},
		Mongo: cfg.Section("mongo").Key("url").String(),
		Prefix: []string{
			"img", "sql", "md",
		},
		Admin: struct {
			Username string
			Password string
		}{
			Username: cfg.Section("admin").Key("username").String(),
			Password: cfg.Section("admin").Key("password").String(),
		},
	}
}

func main() {
	m_logs.InitGinLog("tool")
	r := gin.Default()
	r.StaticFS("/img", http.Dir(filepath.Join(Origin.U.UploadPath, Origin.Prefix[0])))
	InitRouter(r)
	err := r.Run(fmt.Sprintf("%s:%s", Origin.Server.Host, Origin.Server.Port))
	if err != nil {
		return
	}
}

func UpdatePhoto() {
	go func() {
		Origin.Current.mu.Lock()
		defer Origin.Current.mu.Unlock()
		Origin.Current.lastUpdateTime = time.Now()
		c := make(chan string, 1)
		err := BingPhoto(c)
		if err != nil {
			log.Println(err)
			return
		}
		Origin.Current.Url = <-c
		log.Println(Origin.Current.Url)
	}()
}

func init() {
	Origin = NewOriginConfig()
	GormDB = db.NewDB(Origin.BackUp.Dns)
	MCli = db.NewMongo(Origin.Mongo)
	ConfCollection = db.NewCollection(MCli)
	InitBaseConfig()
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {
		go UpdatePhoto()
	})
	_, err = c.AddFunc(fmt.Sprintf("0 0 */%v * *", Origin.BackUp.BackUpCycle), func() {
		m_logs.CleanOldFile(7, filepath.Join(Origin.U.UploadPath, Origin.Prefix[1]))
		InitBackUp(Origin.BackUp.Dns, filepath.Join(Origin.U.UploadPath, Origin.Prefix[0]), Origin.BackUp.BackUpSqlSendEmail)
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}
