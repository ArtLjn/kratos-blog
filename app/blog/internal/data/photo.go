package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/photo"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/pkg/vo"
)

type Photo struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Photo    string `json:"photo"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Position string `json:"position"`
}

func (t *Photo) TableName() string {
	return "blogphoto"
}

type photoRepo struct {
	data *Data
	log  *log.Helper
}

func NewPhotoRepo(data *Data, logger log.Logger) biz.PhotoRepo {
	return &photoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *photoRepo) CreatePhoto(ctx context.Context, request *pb.CreatePhotoRequest) *pb.CreatePhotoReply {
	var photo Photo
	t.data.pf.ParseJSONToStruct(request.Data, &photo)
	if err := t.data.db.Create(&photo).Error; err != nil {
		return &pb.CreatePhotoReply{Common: &pb.CommonReply{Code: 500, Result: vo.INSERT_ERROR}}
	}
	return &pb.CreatePhotoReply{
		Common: &pb.CommonReply{Code: 200, Result: vo.INSERT_SUCCESS},
	}
}

func (t *photoRepo) DeletePhoto(ctx context.Context, request *pb.DeletePhotoRequest) *pb.DeletePhotoReply {
	if err := t.data.pf.DeleteFunc(Photo{}, map[string]interface{}{"id": request.Id}); err != nil {
		return &pb.DeletePhotoReply{
			Common: &pb.CommonReply{
				Code:   500,
				Result: vo.DELETE_ERROR,
			},
		}
	}
	return &pb.DeletePhotoReply{
		Common: &pb.CommonReply{
			Code:   200,
			Result: vo.DELETE_SUCCESS,
		},
	}
}

func (t *photoRepo) SearchAllPhoto(ctx context.Context, request *pb.ListPhotoRequest) *pb.ListPhotoReply {
	data, err := t.data.pf.QueryFunc(Photo{}, nil, true)
	if err != nil {
		t.log.Log(log.LevelError, err)
		return &pb.ListPhotoReply{
			Common: &pb.CommonReply{
				Code:   400,
				Result: vo.QUERY_FAIL,
			},
		}
	}
	var photoData []*pb.PhotoData
	e := t.data.pf.ParseJSONToStruct(data, &photoData)
	if e != nil {
		t.log.Log(log.LevelError, e)
	}
	return &pb.ListPhotoReply{
		Common: &pb.CommonReply{
			Code:   200,
			Result: vo.QUERY_SUCCESS,
		},
		Data: photoData,
	}
}
