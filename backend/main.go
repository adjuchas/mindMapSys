package main

import (
	"backend/conf"
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"backend/router"
)

var ServerConfig conf.ServerConfig

func main() {
	ServerConfig = conf.LoadServerConfig()
	mysqlConn.InitConnMySql(ServerConfig)
	redisConn.InitConnRedis(ServerConfig)

	router.Start()
}
