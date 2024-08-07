package main

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Config 动态可调整配置写入MongoDB
type Config struct {
	Version string // 配置版本号
	Open    bool   // 启用配置
	Upload  struct {
		Path    string //上传路径
		Domain  string //代理域名
		MaxSize string //最大内存
		BingUrl string //每日随机图片获取接口
	}
	BackUp struct {
		Cycle     int  // 备份周期(天)
		OpenEmail bool // 备份是否开启邮件通知
	}
	QQEmail struct {
		Username string // 邮箱账号
		Password string // 邮箱密码
		Host     string
		Port     int
	}
}

func InitBaseConfig() {
	// defaultConfig 是你要插入或更新的默认配置
	var defaultConfig = Config{
		// ... 配置初始化
		Version: "1.0",
		Open:    true,
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
	}

	// 检查是否已有记录
	var existingConfig Config
	err := ConfCollection.FindOne(context.TODO(), bson.D{{}}).Decode(&existingConfig)
	_, e := GetCurrentConfig()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// 集合为空，插入新记录
			CreateConfig(defaultConfig)
			if err = LoadConfig(defaultConfig.Version); err != nil {
				log.Error(err)
				panic(err)
			}
		} else {
			// 其他错误
			panic(err)
		}
	} else if e != nil {
		if errors.Is(e, mongo.ErrNoDocuments) {
			if e = LoadConfig(existingConfig.Version); e != nil {
				log.Error(e)
				panic(e)
			}
		}
	}
	LoadCurrentConfig()
}

func CreateConfig(da Config) bool {
	da.Open = false
	if HasVersion(da.Version) {
		log.Info("Version already exists")
		return false
	}
	insertResult, err := ConfCollection.InsertOne(context.TODO(), da)
	if err != nil {
		log.Error(err)
		return false
	}
	log.Info("Inserted new config with id:", insertResult.InsertedID)
	return true
}

func FindAllConfig() []Config {
	cur, err := ConfCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Error(err)
		return nil
	}
	var result []Config
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var res Config
		if err := cur.Decode(&res); err != nil {
			log.Error(err)
			return nil
		}
		result = append(result, res)
	}
	return result
}

func FindConfig(version string) (Config, error) {
	var result Config
	err := ConfCollection.FindOne(context.TODO(), bson.D{{"version", version}}).Decode(&result)
	if err != nil {
		log.Error(err)
		return Config{}, err
	}
	return result, nil
}

func HasVersion(version string) bool {
	_, err := ConfCollection.FindOne(context.TODO(), bson.D{{"version", version}}).DecodeBytes()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func DeleteConfigByVersion(version string) bool {
	_, err := ConfCollection.DeleteOne(context.TODO(), bson.D{{"version", version}})
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func UpdateConfigByVersion(conf Config) bool {
	conf.Open = false
	if !HasVersion(conf.Version) {
		log.Error("Version does not exist")
		return false
	}
	_, err := ConfCollection.UpdateOne(context.TODO(),
		bson.D{{"version", conf.Version}}, bson.D{{"$set", conf}})
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func LoadConfig(version string) error {
	_, err := ConfCollection.UpdateMany(context.TODO(),
		bson.D{{"open", true}},
		bson.D{{"$set", bson.D{{"open", false}}}},
	)
	if err != nil {
		log.Error(err)
		return err
	}
	if !HasVersion(version) {
		log.Error("Version does not exist")
		return errors.New("version does not exist")
	}
	_, err = ConfCollection.UpdateOne(context.TODO(),
		bson.D{{"version", version}},
		bson.D{{"$set", bson.D{{"open", true}}}},
	)
	if err != nil {
		log.Error(err)
		return err
	}
	LoadCurrentConfig()
	return nil
}

func GetCurrentConfig() (Config, error) {
	var result Config
	err := ConfCollection.FindOne(context.TODO(),
		bson.D{{"open", true}}).Decode(&result)
	if err != nil {
		log.Error(err)
		return Config{}, err
	}
	return result, nil
}

func LoadCurrentConfig() {
	conf, err := GetCurrentConfig()

	if err != nil {
		log.Error(err)
		panic(err)
	}
	Origin.BackUp.BackUpCycle = conf.BackUp.Cycle
	Origin.BackUp.BackUpSqlSendEmail = conf.BackUp.OpenEmail
	Origin.U = NewUploadRepo(conf.Upload.Path, conf.Upload.Domain, conf.Upload.MaxSize, conf.Upload.BingUrl)
	Origin.Mail = NewMail(conf.QQEmail.Username, conf.QQEmail.Host, conf.QQEmail.Port, conf.QQEmail.Password)
	InitPath(Origin.Prefix...)
	UpdatePhoto()
}
