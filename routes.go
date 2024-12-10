package main

import "github.com/gin-gonic/gin"

// Funcție pentru configurarea rutelor API
func setupRoutes(r *gin.Engine) {
	r.GET("/employees", getEmployees)
	r.GET("/employee/:id", getEmployee)
	r.POST("/employee", createEmployee)
	r.PUT("/employee/:id", updateEmployee)
	r.DELETE("/employee/:id", deleteEmployee)
}
