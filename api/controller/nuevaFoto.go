package controller

import (
	"apiFotos/api/config"
	"apiFotos/api/models"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func IngresaFotoTest(w http.ResponseWriter, r *http.Request) {
	nuevaFoto := models.Foto{
		Nombre:      "Foto 1",
		Descripcion: "Foto 1",
		Url:         "Foto 1",
		Usuario_id:  "1",
	}

	db := config.InitDatabase()

	//lamo a la funcion que ingresaFoto

	IngresaFoto(db, nuevaFoto)

	sqlDB, err := db.DB()

	if err != nil {
		log.Println("Error al conectar a la base de datos", err)
	} else {
		sqlDB.Close()
	}

}

func IngresaFoto(db *gorm.DB, nuevaFoto models.Foto) {
	//almacenar el neva foto en la base de datos

	err := db.Create(&nuevaFoto).Error
	if err != nil {
		log.Println("Error al ingresar la foto", err)
	} else {
		log.Println("Foto ingresada con éxito")
	}

}
