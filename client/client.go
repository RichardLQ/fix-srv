package client

import (
	"github.com/jinzhu/gorm"
)

var Global = struct {
	UserConf        CostVar  `confs:"read=rainbow;format=json"`
	Mini            gorm.DB  `confs:"name=Mini;read=rainbow"`
	FixServiceConf UserConf `confs:"read=rainbow;format=json"`
}{}

//CostVar 加密字段
type CostVar struct {
	JwtKey string
}

//UserConf 运行端口
type UserConf struct {
	JwtKey   string
	HttpPort string
	HttpIp   string
	RpcPort  string
	RpcIp    string
}
