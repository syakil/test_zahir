package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"test-zahir/contact"
	"test-zahir/handler"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/test_zahir?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	mig := db.AutoMigrate(&contact.ContactData{})
	if mig != nil {
		log.Fatal("Database Migration Failed")
	}
	if err != nil {
		log.Fatal("Database Connection Failed")
	}
	router := gin.Default()

	router.GET("/contact", handler.GetContact)
	router.GET("/contact/:id", handler.GetContactById)
	router.POST("/contact", handler.CreateContact)
	router.Run()

}
