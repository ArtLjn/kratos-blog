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
)

// RoleSwitch
const (
	Admin   = "admin"
	User    = "user"
	Visitor = "visitor"
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
