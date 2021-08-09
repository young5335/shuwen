package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shuwen/config"
)

func InitDb() (*gorm.DB,error) {
	host := config.Host
	dbname := config.Database
	username := config.Username
	password := config.Password
	charset := config.Charset
	port := config.Port

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", username, password, host, port, dbname, charset)
	db, err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil {
		return nil,err
	}

	return db,nil
}
