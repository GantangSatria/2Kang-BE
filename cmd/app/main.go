package main

import (
	"log"
	"2Kang/internal/bootstrap"
)

func main() {
	app := bootstrap.InitializeApp()
	log.Println("Server running on http://localhost:8080")
	app.Listen(":8080")
}