package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test-zahir/contact"
)

func GetContact(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama": "syakil",
	})
}

func GetContactById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func CreateContact(c *gin.Context) {
	var contacts contact.Contact

	err := c.ShouldBindJSON(&contacts)
	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition:%s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         contacts.Id,
		"name":       contacts.Name,
		"gender":     contacts.Gender,
		"phone":      contacts.Phone,
		"email":      contacts.Email,
		"created_at": contacts.CreatedAt,
		"updated_at": contacts.UpdatedAt,
	})
}
