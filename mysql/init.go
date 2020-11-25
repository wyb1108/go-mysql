package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Db *sql.DB
)

func Init(file string) {
	err := initConfig(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if mysqlConfig.Host == "" {
		fmt.Println("mysql服务地址未配置, 请在mysqlConfig.json中配置")
	}
	if mysqlConfig.Schema == "" {
		fmt.Println("访问数据库未配置, 请在mysqlConfig.json中配置")
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Schema)
	fmt.Println(connString)
	Db, err = sql.Open("mysql", connString)
	Db.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	Db.SetMaxIdleConns(mysqlConfig.MaxIdleConn)

}
