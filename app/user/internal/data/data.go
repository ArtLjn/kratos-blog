package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"kratos-blog/app/user/internal/conf"
	db2 "kratos-blog/pkg/db"
	"kratos-blog/pkg/model"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	log     *log.Helper
	c       *conf.Bootstrap
	db      *gorm.DB
	rdb     *redis.Client
	pf      model.PublicFunc
	mgo     *mongo.Client
	collect *mongo.Collection
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	pf := model.NewOFunc(l, db)
	cli := db2.NewMongo(c.Mongo.GetUrl())
	return &Data{log: l, c: c, db: db, rdb: rdb, pf: pf, mgo: cli,
		collect: db2.NewCollection(cli)}, nil
}
