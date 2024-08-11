package server

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

const (
	BlogConfPath    = "app/blog/configs"
	CommentConfPath = "app/comment/configs"
	UserConfPath    = "app/user/configs"
	GateWayConfPath = "app/gateway/configs"
	TooPath         = "/Users/ljn/Documents/HongDou-Go-Blog/kratos-blog/tool/config.ini"
)

// logPath

const (
	LogOutStreamPath = "log"
)
