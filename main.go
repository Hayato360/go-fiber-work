package main

import (
	"fmt"
	"go-fiber-work/database"
	m "go-fiber-work/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)

	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database successfully")
	database.DBConn.AutoMigrate(&m.Dogs{})
	database.DBConn.AutoMigrate(&m.Company{})
}


func main() {
	app := fiber.New()
	initDatabase()
	routes.Ine
}