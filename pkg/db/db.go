/*
@Time : 2024/8/8 下午1:37
@Author : ljn
@File : db
@Software: GoLand
*/

package db

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dns string) *gorm.DB {
	g, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	return g
}

func NewMongo(url string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(url)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewCollection(cli *mongo.Client) *mongo.Collection {
	return cli.Database("hongDouBlog").Collection("config")
}

func NewRegistrar(address, scheme string) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = address
	c.Scheme = scheme
	c.Namespace = ""
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(address, scheme string) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = address
	c.Scheme = scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRDB(addr, password string, db int64) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     addr,
			DB:       int(db),
			Password: password,
		},
	)
}
