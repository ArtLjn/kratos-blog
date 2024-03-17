package server

const (
	BlogService    = "blog.service"
	CommentService = "comment.service"
	UserService    = "user.service"
	GateWay        = "gateway"
)

const (
	AdminNotes  string = "adminNotes"
	Notes       string = "notes"
	TableName   string = "BlogVisits"
	Comment     string = "gateway"
	Appear      string = "appear"
	CacheBlog   string = "cacheBlog"
	SuggestBlog string = "suggestBlog"
	Token              = "token"
)

// RoleSwitch
const (
	Admin   = "admin"
	User    = "user"
	Visitor = "visitor"
)

// ConfigPath

var (
	//BlogConfPath    = "/usr/local/hongDou/blog/configs"
	//CommentConfPath = "/usr/local/hongDou/comment/configs"
	//UserConfPath    = "/usr/local/hongDuo/user/configs"
	//GateWayConfPath = "/usr/local/hongDou/gateway/configs"

	BlogConfPath    = "configs"
	CommentConfPath = "/usr/local/hongDou/comment/configs"
	UserConfPath    = "/usr/local/hongDuo/user/configs"
	GateWayConfPath = "/usr/local/hongDou/gateway/configs"

	//BlogConfPath    = "app/blog/configs"
	//CommentConfPath = "app/comment/configs"
	//UserConfPath    = "app/user/configs"
	//GateWayConfPath = "app/gateway/configs"
)
