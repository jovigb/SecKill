package setup

import (
	"SecKill/sk_admin/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

func InitMysql(hostMysql, portMysql, userMysql, pwdMysql, dbMysql string) {
	// DbConfig := map[string]interface{}{
	// 	// Default database configuration
	// 	"Default": "mysql_dev",
	// 	// (Connection pool) Max open connections, default value 0 means unlimit.
	// 	"SetMaxOpenConns": 300,
	// 	// (Connection pool) Max idle connections, default value is 1.
	// 	"SetMaxIdleConns": 10,
	// 	// Define the database configuration character "mysql_dev".
	// 	"Connections": map[string]map[string]string{
	// 		"mysql_dev": map[string]string{
	// 			"host":     hostMysql,
	// 			"username": userMysql,
	// 			"password": pwdMysql,
	// 			"port":     portMysql,
	// 			"database": dbMysql,
	// 			"charset":  "utf8",
	// 			"protocol": "tcp",
	// 			"prefix":   "",      // Table prefix
	// 			"driver":   "mysql", // Database driver(mysql,sqlite,postgres,oracle,mssql)
	// 		},
	// 	},
	// }
	DbConfig := &gorose.DbConfigSingle{
		Driver:          "mysql",                                                                                             // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,                                                                                                // 是否开启sql日志
		SetMaxOpenConns: 300,                                                                                                 // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 10,                                                                                                  // (连接池)闲置的连接数
		Prefix:          "",                                                                                                  // 表前缀
		Dsn:             fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userMysql, pwdMysql, hostMysql, portMysql, dbMysql), // 数据库链接
	}

	connection, err := gorose.Open(DbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	config.SecAdminConfCtx.DbConf = &config.DbConf{DbConn: connection}
}
