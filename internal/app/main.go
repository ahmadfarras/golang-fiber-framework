package main

import (
	"ahmadfarras/fiberframework/internal/route"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(:3306)/belajar_golang?parseTime=true"
	gormConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := gormConn.DB()
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
	})
	route.InitHttpRoute(app, gormConn)
	app.Listen(":8080")
}
