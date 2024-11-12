package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"kratos-blog/app/blog/internal/conf"
	"kratos-blog/pkg/model"
	"sync"
)

var (
	CTX = context.Background()
	// ProviderSet is data providers.
	ProviderSet = wire.NewSet(NewData)
	mu          sync.Mutex
)

// Data .
type Data struct {
	log *log.Helper
	db  *gorm.DB
	rdb *redis.Client
	pf  model.PublicFunc
	c   *conf.Bootstrap
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	pf := model.NewOFunc(l, db)
	return &Data{log: l, db: db, rdb: rdb, pf: pf, c: c}, nil
}

//func InitDB(db *gorm.DB) {
//	if err := db.AutoMigrate(
//		&Blog{},
//		&Tag{},
//		&Friend{},
//		&Photo{},
//	); err != nil {
//		panic(err)
//	}
//}
