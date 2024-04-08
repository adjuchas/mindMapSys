package conf

import (
	"github.com/go-ini/ini"
)

type mysqlDB struct {
	USER        string
	PASSWORD    string
	HOST        string
	NAME        string
	TABLEPREFIX string
}

type redisDb struct {
	HOST     string
	PASSWORD string
	INDEX    string
}

type ServerConfig struct {
	MYSQLDB mysqlDB
	REDISDB redisDb
}

func LoadServerConfig() ServerConfig {

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	database, err := cfg.GetSection("database")
	if err != nil {
		panic(err)
	}

	redis, err := cfg.GetSection("redis")
	if err != nil {
		panic(err)
	}

	Config := ServerConfig{
		MYSQLDB: mysqlDB{
			USER:        database.Key("USER").MustString(""),
			PASSWORD:    database.Key("PASSWORD").MustString(""),
			HOST:        database.Key("HOST").MustString(""),
			NAME:        database.Key("NAME").MustString(""),
			TABLEPREFIX: database.Key("TABLEPREFIX").MustString(""),
		},
		REDISDB: redisDb{
			HOST:     redis.Key("HOST").MustString(""),
			PASSWORD: redis.Key("PASSWORD").MustString(""),
			INDEX:    redis.Key("INDEX").MustString(""),
		},
	}
	return Config
}
