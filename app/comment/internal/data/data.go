package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"kratos-blog/app/comment/internal/conf"
	"kratos-blog/pkg/model"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

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
