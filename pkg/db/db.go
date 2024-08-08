/*
@Time : 2024/8/8 下午1:37
@Author : ljn
@File : db
@Software: GoLand
*/

package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
