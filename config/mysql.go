package config

import (
	"fmt"

	"github.com/cecep31/fiberapi/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(config Config) *gorm.DB {
	dbuser := config.Get("DB_USER")
	dbpass := config.Get("DB_PASS")
	dbhost := config.Get("DB_HOST")
	dbport := config.Get("DB_PORT")
	dbname := config.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	return db
}
