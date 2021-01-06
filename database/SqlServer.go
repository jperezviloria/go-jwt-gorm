package database

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectSqlServer() *gorm.DB {

	dsn := "sqlserver://sa:Password01.@localhost:1433?database=JWTAPP"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
