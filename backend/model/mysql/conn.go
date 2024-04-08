package mysqlConn

import (
	"backend/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConnMySql(conf conf.ServerConfig) {
	var err error
	dbParams := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MYSQLDB.USER,
		conf.MYSQLDB.PASSWORD,
		conf.MYSQLDB.HOST,
		conf.MYSQLDB.NAME,
	)
	DB, err = gorm.Open(mysql.Open(dbParams))

	//DB, err = gorm.Open(mysql.Open(dbParams), &gorm.Config{})
	if err != nil {

	}
	sqlDB, err2 := DB.DB()
	if err2 != nil {

	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
