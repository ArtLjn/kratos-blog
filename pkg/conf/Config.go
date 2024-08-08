/*
@Time : 2024/8/8 下午1:50
@Author : ljn
@File : Config
@Software: GoLand
*/

package conf

// Config 动态可调整配置写入MongoDB
type Config struct {
	Version string // 配置版本号
	Open    bool   // 启用配置
	Upload  struct {
		Path    string //上传路径
		Domain  string //代理域名
		MaxSize string //最大内存
		BingUrl string //每日随机图片获取接口
	}
	BackUp struct {
		Cycle     int  // 备份周期(天)
		OpenEmail bool // 备份是否开启邮件通知
	}
	QQEmail struct {
		Username string // 邮箱账号
		Password string // 邮箱密码
		Host     string
		Port     int
	}
	Admin struct {
		Username string
		Password string
	}
}
