package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Global DB object
var db *gorm.DB
var err error

// Funcție pentru conectarea la baza de date PostgreSQL
func connectToDatabase() {
	dsn := "user=postgres password=123456789 dbname=employee_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Nu s-a putut conecta la baza de date:", err)
	}
	fmt.Println("Conectat cu succes la PostgreSQL")
	// Migrare pentru a crea tabela daca nu exista
	db.AutoMigrate(&Employee{})
}
