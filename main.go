package main

import (
	"apiFotos/api/config"
	"apiFotos/api/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// inicializo la base de datos

	db := config.InitDatabase()

	// inicializo gin
	router := gin.Default()

	//configuro cors para conectarme con el front

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
	}

	router.Use(cors.New(config))

	//configuro las rutas

	routes.ConfigureFotoRouter(router, db)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("servidor escucha en el puerto:" + PORT)

	http.ListenAndServe(":"+PORT, router)
}
