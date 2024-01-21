package services

import (
	"apiFotos/api/models"
	"log"

	"gorm.io/gorm"
)

//crud de fotos

func GetFotos(db *gorm.DB) ([]models.Foto, error) {
	//obtener todas las fotos de la base de datos

	var fotos []models.Foto
	err := db.Find(&fotos).Error
	if err != nil {
		log.Println("Error al obtener las fotos", err)
		return nil, err
	}
	return fotos, nil
}

func CreateFoto(db *gorm.DB, foto *models.Foto) error {
	//crear una foto en la base de datos

	err := db.Create(foto).Error
	if err != nil {
		log.Println("Error al crear la foto", err)
		return err
	}
	return nil
}
