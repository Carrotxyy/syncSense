package setting

type Config struct {
	WxAddr      string `json:"wxaddr"`       // 微信服务器域名
	DbType      string `json:"dbtype"`       // 数据库类型
	DbUser      string `json:"dbuser"`       // 数据库账号
	DbPassword  string `json:"dbpassword"`   // 数据库密码
	DbIP        string `json:"dbip"`         // 数据库ip
	DbName      string `json:"dbname"`       // 数据库名
	TablePrefix string `json:"table-prefix"` // 数据库表扩展
}

