package main

// Definirea structurii pentru Employee
type Employee struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}
