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

var (
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

var (
	LogOutStreamPath = "log"
)
