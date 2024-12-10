package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Funcție pentru a obține toți angajații
func getEmployees(c *gin.Context) {
	var employees []Employee
	if err := db.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nu s-au găsit angajați"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// Funcție pentru a obține un angajat după ID
func getEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	if err := db.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Angajatul nu a fost găsit"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// Funcție pentru a crea un nou angajat
func createEmployee(c *gin.Context) {
	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date invalid"})
		return
	}

	if err := db.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nu s-a putut salva angajatul"})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

// Funcție pentru a actualiza un angajat existent
func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	if err := db.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Angajatul nu a fost găsit"})
		return
	}

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date invalid"})
		return
	}

	if err := db.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nu s-a putut actualiza angajatul"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// Funcție pentru a șterge un angajat
func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	if err := db.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Angajatul nu a fost găsit"})
		return
	}

	if err := db.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nu s-a putut șterge angajatul"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Angajatul a fost șters"})
}
