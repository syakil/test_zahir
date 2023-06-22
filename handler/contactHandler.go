package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test-zahir/contact"
)

type contactHandler struct {
	contactService contact.Service
}

func NewContactHandler(contactService contact.Service) *contactHandler {
	return &contactHandler{contactService}
}

func (handler *contactHandler) GetContact(c *gin.Context) {
	contactResponse, err := handler.contactService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contactResponse,
	})
}

func (handler *contactHandler) GetContactById(c *gin.Context) {
	id := c.Param("id")
	contactResponse, err := handler.contactService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contactResponse,
	})
}

func (handler *contactHandler) CreateContact(c *gin.Context) {
	var contactRequest contact.ContactRequest

	err := c.ShouldBindJSON(&contactRequest)
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
	contactResponse, err := handler.contactService.Create(contactRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contactResponse,
	})
}

func (handler *contactHandler) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	_, err := handler.contactService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Data Success",
	})
}

func (handler *contactHandler) UpdateContact(c *gin.Context) {
	var contactUpdate contact.ContactUpdate

	err := c.ShouldBindJSON(&contactUpdate)
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
	id := c.Param("id")
	contactResponse, err := handler.contactService.Update(id, contactUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contactResponse,
	})
}
