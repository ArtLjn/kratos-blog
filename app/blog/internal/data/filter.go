package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type FilterRepo struct {
	data *Data
	log  *log.Helper
}

func NewFilterRepo(data *Data, logger log.Logger) *FilterRepo {
	return &FilterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f *FilterRepo) NewFilter(whiteList []string) http.FilterFunc {
	return f.data.role.FilterPermission(whiteList)
}
