package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	pb "kratos-blog/api/comment"
	"log"
)

var (
	CommentRequest = &pb.CreateCommentRequest{}
	RewardRequest  = &pb.CreateRewardRequest{}
)

func ReadComment(c *websocket.Conn, raw int) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error1", err)
			c.Close()
			return
		}
		switch raw {
		case 0:
			if err := json.Unmarshal(message, CommentRequest); err != nil {
				panic("json.Unmarshal error")
			}
			if CommentRequest == nil {
				panic("CommentRequest is nil")
			}
		case 1:
			err := json.Unmarshal(message, RewardRequest)
			if err != nil {
				panic("json.Unmarshal error")
			}
			if RewardRequest == nil {
				panic("RewardRequest is nil")
			}
		}
	}
}
