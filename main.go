package main

import (
	"apiFotos/api/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	handler.Init()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("servidor escucha en el puerto:" + PORT)

	http.ListenAndServe(":"+PORT, nil)
}
