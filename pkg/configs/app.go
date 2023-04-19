package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load env")
	} else {
		fmt.Println("Loaded all env")
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")

	dsn := username + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)

	d, err := gorm.Open(driver, dsn)
	// d, err := gorm.Open("mysql", "root:Asdf#1234@(localhost)/bookstore?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
