package vo

type CodeInfo struct {
	Code int           `json:"code"`
	Info interface{}   `json:"info"`
	Data interface{}   `json:"data"`
	List []interface{} `json:"list"`
}

type CodeBuilder interface {
	SetCode(code int) CodeBuilder
	SetInfo(interface{}) CodeBuilder
	SetData(interface{}) CodeBuilder
	SetList([]interface{}) CodeBuilder
	GetCodeInfo() CodeInfo
}

// Director
type Director struct {
	builder CodeBuilder
}

func (director *Director) Construct() {
	director.builder.SetCode(200).SetInfo(nil).SetData(nil).SetList(nil) //链式调用
}

func (director *Director) SetBuilder(builder CodeBuilder) {
	director.builder = builder
}

type Response struct {
	codeInfo CodeInfo
}

const (
	INSERT_ERROR        = "插入错误"
	INSERT_SUCCESS      = "插入成功"
	JSON_ERROR          = "JSON格式错误"
	DELETE_ERROR        = "删除失败"
	DELETE_SUCCESS      = "删除成功"
	OPERATE_FAIL        = "操作失败"
	UPDATE_FAIL         = "更新失败"
	UPDATE_SUCCESS      = "更新成功"
	QUERY_SUCCESS       = "查询成功"
	QUERY_FAIL          = "查询失败"
	QUERY_EMPTY         = "查询为空"
	FORBIDDEN_ACCESS    = "禁止访问"
	REQUEST_FAIL        = "请求失败"
	TALK_ERROR          = "请文明发言"
	TALK_FORBIDDEN      = "禁止评论"
	FREQUENCY_ANOMALIES = "频率异常"
	KEY_ERROR           = "验证码错误"
	NAME_EXIST          = "用户名已存在"
	EMAIL_EXIST         = "该邮箱已经绑定其他账户"
	REGISTER_SUCCESS    = "注册成功"
	LOGIN_SUCCESS       = "登录成功"
	LOGIN_FAIL          = "用户名或密码错误"
	EMAIL_SUCCESS       = "邮件发送成功"
	EMAIL_FAIL          = "邮件发送失败"
	EMAIL_NOT_MATCH     = "邮箱与注册邮箱不符"
	EMAIL_COOLDOWN      = "请稍后再试！"
	LIST_EMPTY          = "空列表"
	PERMISSION_ERROR    = "权限错误"
	CREATE_SUCCESS      = "创建成功"
	CREATE_FAIL         = "创建失败"
)

func (r *Response) SetCode(code int) CodeBuilder {
	r.codeInfo.Code = code
	return r
}

func (r *Response) SetInfo(info interface{}) CodeBuilder {
	r.codeInfo.Info = info
	return r
}

func (r *Response) SetData(data interface{}) CodeBuilder {
	r.codeInfo.Data = data
	return r
}

func (r *Response) SetList(list []interface{}) CodeBuilder {
	r.codeInfo.List = list
	return r
}

func (r *Response) GetCodeInfo() CodeInfo {
	return r.codeInfo
}
