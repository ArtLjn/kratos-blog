package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/friend"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/pkg/vo"
	"time"
)

type Friend struct {
	ID      int    `json:"id" gorm:"primary_key;auto_increment"`
	Title   string `json:"title"`
	Preface string `json:"preface"`
	Url     string `json:"url"`
	Photo   string `json:"photo"`
	Date    string `json:"date"`
}

func (f *Friend) TableName() string {
	return "friend_table"
}

type friendRepo struct {
	data *Data
	log  *log.Helper
}

func NewFriendRepo(data *Data, logger log.Logger) biz.FriendRepo {
	return &friendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *friendRepo) CreateFriend(ctx context.Context, request *pb.CreateFriendRequest) *pb.CreateFriendReply {
	var friend Friend
	t.data.pf.ParseJSONToStruct(request.Data, &friend)
	friend.Date = time.Now().Format("2006-01-02")
	if err := t.data.db.Create(&friend).Error; err != nil {
		return &pb.CreateFriendReply{Common: &pb.CommonReply{Code: 500, Result: vo.INSERT_ERROR}}
	}
	return &pb.CreateFriendReply{
		Common: &pb.CommonReply{Code: 200, Result: vo.INSERT_SUCCESS},
	}
}

func (t *friendRepo) DeleteFriend(ctx context.Context, request *pb.DeleteFriendRequest) *pb.DeleteFriendReply {
	if err := t.data.pf.DeleteFunc(Friend{}, map[string]interface{}{"id": request.Id}); err != nil {
		return &pb.DeleteFriendReply{
			Common: &pb.CommonReply{
				Code:   500,
				Result: vo.DELETE_ERROR,
			},
		}
	}
	return &pb.DeleteFriendReply{
		Common: &pb.CommonReply{
			Code:   200,
			Result: vo.DELETE_SUCCESS,
		},
	}
}

func (t *friendRepo) SearchAllFriend(ctx context.Context, request *pb.ListFriendRequest) *pb.ListFriendReply {
	data, err := t.data.pf.QueryFunc(Friend{}, nil, true)
	if err != nil {
		t.log.Log(log.LevelError, err)
		return &pb.ListFriendReply{
			Common: &pb.CommonReply{
				Code:   400,
				Result: vo.QUERY_FAIL,
			},
		}
	}
	var friendData []*pb.FriendData
	e := t.data.pf.ParseJSONToStruct(data, &friendData)
	if e != nil {
		t.log.Log(log.LevelError, e)
	}
	if len(friendData) == 0 {
		return &pb.ListFriendReply{
			Common: &pb.CommonReply{
				Code:   400,
				Result: vo.QUERY_EMPTY,
			},
		}
	}
	return &pb.ListFriendReply{
		Common: &pb.CommonReply{
			Code:   200,
			Result: vo.QUERY_SUCCESS,
		},
		Data: friendData,
	}
}

func (t *friendRepo) UpdateFriend(ctx context.Context, request *pb.UpdateFriendRequest) *pb.UpdateFriendReply {
	var val map[string]interface{}
	t.data.pf.ParseJSONToStruct(request.Data, &val)
	err := t.data.pf.UpdateFunc(Friend{}, map[string]interface{}{"id": request.Data.Id}, val, false)
	if err != nil {
		t.log.Log(log.LevelError, err)
		return &pb.UpdateFriendReply{
			Common: &pb.CommonReply{
				Code:   vo.BAD_REQUEST,
				Result: vo.UPDATE_FAIL,
			},
		}
	}
	return &pb.UpdateFriendReply{Common: &pb.CommonReply{Code: vo.SUCCESS_REQUEST, Result: vo.UPDATE_SUCCESS}}
}

func (t *friendRepo) GetFriendByCond(ctx context.Context, request *pb.GetFriendRequest) *pb.GetFriendReply {
	body, err := t.data.pf.QueryFunc(Friend{}, map[string]interface{}{"id": request.Id}, false)
	if err != nil {
		t.log.Log(log.LevelError, err)
	}
	var friend pb.FriendData
	t.data.pf.ParseJSONToStruct(body, &friend)
	return &pb.GetFriendReply{
		Common: &pb.CommonReply{
			Code:   vo.SUCCESS_REQUEST,
			Result: vo.QUERY_SUCCESS,
		},
		Data: &friend,
	}
}
