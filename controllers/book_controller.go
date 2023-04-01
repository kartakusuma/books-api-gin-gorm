package controllers

import (
	"books-api-gin-gorm/database"
	"books-api-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBooks(c *gin.Context) {
	db := database.GetDB()

	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": books,
	})
}

func GetBookByID(c *gin.Context) {
	db := database.GetDB()
	bookID := c.Param("id")

	var book models.Book

	if err := db.First(&book, bookID).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Book data is not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func CreateBook(c *gin.Context) {
	db := database.GetDB()

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := db.Create(&book).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {
	db := database.GetDB()
	bookID := c.Param("id")

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if db.Model(&book).Where("id = ?", bookID).Updates(&book).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Can not update book data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update book data",
		"book":    book,
	})
}

func DeleteBook(c *gin.Context) {
	db := database.GetDB()
	bookID := c.Param("id")

	var book models.Book

	if db.Delete(&book, bookID).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Can not delete book data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book data has been successfully deleted",
	})
}
