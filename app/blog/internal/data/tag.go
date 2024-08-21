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
	TagName string `gorm:"column:tagName" json:"tagName"`
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
	var tag Tag
	t.data.pf.ParseJSONToStruct(request.Data, &tag)
	if err := t.data.db.Create(&tag).Error; err != nil {
		return &pb.CreateTagReply{Common: &pb.CommonReply{Code: 500, Result: vo.INSERT_ERROR}}
	}
	return &pb.CreateTagReply{
		Common: &pb.CommonReply{Code: 200, Result: vo.INSERT_SUCCESS},
	}
}

func (t *tagRepo) DeleteTag(ctx context.Context, request *pb.DeleteTagRequest) *pb.DeleteTagReply {
	if err := t.data.pf.DeleteFunc(Tag{}, map[string]interface{}{"id": request.Id}); err != nil {
		return &pb.DeleteTagReply{
			Common: &pb.CommonReply{
				Code:   500,
				Result: vo.DELETE_ERROR,
			},
		}
	}
	return &pb.DeleteTagReply{
		Common: &pb.CommonReply{
			Code:   200,
			Result: vo.DELETE_SUCCESS,
		},
	}
}

func (t *tagRepo) SearchAllTag(ctx context.Context, request *pb.ListTagRequest) *pb.ListTagReply {
	data, err := t.data.pf.QueryFunc(Tag{}, nil, true)
	if err != nil {
		t.log.Log(log.LevelError, err)
		return &pb.ListTagReply{
			Common: &pb.CommonReply{
				Code:   400,
				Result: vo.QUERY_FAIL,
			},
		}
	}
	var tagData []*pb.TagData
	e := t.data.pf.ParseJSONToStruct(data, &tagData)
	if e != nil {
		t.log.Log(log.LevelError, e)
	}
	return &pb.ListTagReply{
		Common: &pb.CommonReply{
			Code:   200,
			Result: vo.QUERY_SUCCESS,
		},
		Data: tagData,
	}
}
