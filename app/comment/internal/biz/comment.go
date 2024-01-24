package biz

import "github.com/go-kratos/kratos/v2/log"

type CommRepo interface {
}

type CommUseCase struct {
	repo CommRepo
	log  *log.Helper
}

func NewUCommUseCase(repo CommRepo, logger log.Logger) *CommUseCase {
	return &CommUseCase{repo: repo, log: log.NewHelper(logger)}
}
