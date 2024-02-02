package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/tag"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/pkg/vo"
)

type Tag struct {
	ID      int    `json:"id" gorm:"primary_key;auto_increment"`
	TagName string `json:"tagName"`
}

func (t *Tag) TableName() string {
	return "tag"
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *tagRepo) CreateTag(ctx context.Context, request *pb.CreateTagRequest) *pb.CreateTagReply {
	permission := t.data.role.QueryUserMsg(ctx).GetRole().CheckPermission()
	if !permission {
		return &pb.CreateTagReply{
			Common: &pb.CommonReply{
				Code:   401,
				Result: vo.PERMISSION_ERROR,
			},
		}
	}
	var tag Tag
	t.data.pf.ParseJSONToStruct(request, &tag)
	if err := t.data.db.Create(&tag).Error; err != nil {
		return &pb.CreateTagReply{Common: &pb.CommonReply{Code: 500, Result: vo.INSERT_ERROR}}
	}
	return &pb.CreateTagReply{
		Common: &pb.CommonReply{Code: 200, Result: vo.INSERT_SUCCESS},
	}
}
