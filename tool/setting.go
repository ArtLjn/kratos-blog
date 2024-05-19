package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormDB 基础设置

var (
	GormDB  *gorm.DB
	Dns     string
	OutPath string
)

func NewDB(option ...string) {
	Dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		option[0], option[1], option[2], option[3], option[4], option[5])
	g, err := gorm.Open(mysql.Open(Dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	GormDB = g
}

type Setting struct {
	key string
	val string
}

func GetVal(key string) Setting {
	var setting Setting
	if err := GormDB.Where("key = ?", key).First(&setting).Error; err != nil {
		return Setting{}
	}
	return setting
}

func SetVal(setting Setting) {
	if err := GormDB.Where("key = ?", setting.key).First(&setting).Error; err != nil {
		GormDB.Create(setting)
	} else {
		GormDB.Model(&setting).Update("val", setting.val)
	}
}
