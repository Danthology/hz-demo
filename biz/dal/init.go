package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

const MySQLDSN = "root:root@(localhost:3306)/hz_demo?charset=utf8mb4&parseTime=True&loc=Local"

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
