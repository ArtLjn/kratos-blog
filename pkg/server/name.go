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
	//BlogConfPath    = "blog/configs"
	//CommentConfPath = "comment/configs"
	//UserConfPath    = "user/configs"
	//GateWayConfPath = "gateway/configs"

	BlogConfPath    = "app/blog/configs"
	CommentConfPath = "app/comment/configs"
	UserConfPath    = "app/user/configs"
	GateWayConfPath = "app/gateway/configs"

	//BlogConfPath    = "../../configs"
	//CommentConfPath = "../../configs"
	//UserConfPath    = "../../configs"
	//GateWayConfPath = "../../configs"
)

// logPath

const (
	LogOutStreamPath = "log"
)

// 服务连接密钥

const (
	//Kratos_BlogKey 博客服务请求头
	Kratos_BlogKey = "X-HongDou-Key-Gateway"
	//ToolKey 资源服务器请求头
	ToolKey = "X-Tool-Key"
)
