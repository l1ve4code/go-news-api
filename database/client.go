package database

import (
	"go-news-api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connString string){
	Instance, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Connected to Database...")
}

func Migrate()  {
	Instance.AutoMigrate(&entities.News{}, &entities.Comment{})
	log.Println("Database Migration Completed...")
}