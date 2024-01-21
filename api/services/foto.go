package services

import (
	"apiFotos/api/models"
	"log"

	"gorm.io/gorm"
)

//crud de fotos

func CreateFoto(db *gorm.DB, foto *models.Foto) error {
	//crear una foto en la base de datos

	err := db.Create(foto).Error
	if err != nil {
		log.Println("Error al crear la foto", err)
		return err
	}
	return nil
}

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

func DeleteFoto(db *gorm.DB, foto *models.Foto) error {
	//eliminar una foto de la base de datos

	err := db.Delete(foto).Error
	if err != nil {
		log.Println("Error al eliminar la foto", err)
		return err
	}
	return nil
}

func UpdateFoto(db *gorm.DB, foto *models.Foto) error {
	//actualizar una foto de la base de datos

	err := db.Save(foto).Error
	if err != nil {
		log.Println("Error al actualizar la foto", err)
		return err
	}
	return nil
}
