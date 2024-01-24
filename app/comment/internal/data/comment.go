package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/app/comment/internal/biz"
)

type commRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommRepo(data *Data, logger log.Logger) biz.CommRepo {
	return &commRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
