package mysql

import (
	"github.com/wyb1108/go-utils/util"
)

type MysqlConf struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Sentinel    bool   `json:"sentinel"`
	Schema      string `json:"schema"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	MaxOpenConn int    `json:"maxOpenConn"`
	MaxIdleConn int    `json:"maxIdleConn"`
}

var (
	mysqlConfig MysqlConf
)

func initConfig(configFile string) error {
	return util.ParseJsonFile(configFile, &mysqlConfig)
}
