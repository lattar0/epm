package controllers

import (
	"epm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&employee)
	c.JSON(http.StatusOK, employee)
}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	models.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := models.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := models.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&employee)
	c.JSON(http.StatusOK, employee)
}

// Deleta um livro
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := models.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	models.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
