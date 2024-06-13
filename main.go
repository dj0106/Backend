package main

import (
	Config "InvitKaro/Configuration"
	"InvitKaro/router"
	"fmt"
)

var ServerConfig Config.ServerConfiguration

func init() {

	//mysqldb.InitializeMySqlConn(Config.ServerConfig.MYSQL_HOST,
	//	Config.ServerConfig.MYSQL_PORT, Config.ServerConfig.MYSQL_USERNAME,
	//	Config.ServerConfig.MYSQL_DB_NAME, Config.ServerConfig.MYSQL_PASSWORD)

}

func StartService() {
	router.RouteDispatcher()
}

func main() {
	fmt.Print("Let's start")
	StartService()
}
