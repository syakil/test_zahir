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
	contactRepository := contact.NewRepository(db)
	contactService := contact.NewService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)
	router := gin.Default()

	router.GET("/contact", contactHandler.GetContact)
	router.GET("/contact/:id", contactHandler.GetContactById)
	router.POST("/contact", contactHandler.CreateContact)
	router.Run()

}
