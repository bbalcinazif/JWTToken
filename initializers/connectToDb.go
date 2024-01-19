package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectedToDB() {
	var err error
	dsn := "host=berry.db.elephantsql.com user=vjtqjeqo password=2M1bfUw5YBPubQEwXIoRkRybOt1vJsqq dbname=vjtqjeqo port=5432 sslmode=disable "
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
}
