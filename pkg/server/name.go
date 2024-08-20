package server

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	BlogService    = "blog.service"
	CommentService = "comment.service"
	UserService    = "user.service"
	GateWay        = "gateway"
)

const (
	AdminNotes     string = "adminNotes"
	Notes          string = "notes"
	TableName      string = "BlogVisits"
	Comment        string = "comment"
	Appear         string = "appear"
	CacheBlog      string = "cacheBlog"
	SuggestBlog    string = "suggestBlog"
	Token                 = "token"
	TokenMangerKey        = "token_manager"
	AdminToken            = "admin_token"
)

// RoleSwitch
const (
	VIP         = "vip"
	User        = "user"
	Visitor     = "visitor"
	SystemAdmin = "admin"
)

func GetRole(role string) int32 {
	switch role {
	case SystemAdmin:
		return SystemManager
	case VIP:
		return Vip
	case User:
	case Visitor:
		return VisitorOrUser
	}
	return VisitorOrUser
}

const (
	SystemManager = iota
	Vip
	VisitorOrUser
)

// ConfigPath

var (
	BlogConfPath     string
	CommentConfPath  string
	UserConfPath     string
	GateWayConfPath  string
	TooPath          string
	LogOutStreamPath string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	BlogConfPath = os.Getenv("BLOG_CONF_PATH")
	CommentConfPath = os.Getenv("COMMENT_CONF_PATH")
	UserConfPath = os.Getenv("USER_CONF_PATH")
	GateWayConfPath = os.Getenv("GATEWAY_CONF_PATH")
	TooPath = os.Getenv("TOOL_CONF_PATH")
	LogOutStreamPath = os.Getenv("LOG_OUT_STREAM_PATH")
}
