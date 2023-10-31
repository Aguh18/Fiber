package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit()   {
	var err error
	const Postgres = "host=localhost user=postgres password=postgres dbname=fiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := Postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}