package models

type Foto struct {
	ID          int    `gorm:"id"`
	Nombre      string `gorm:"nombre"`
	Descripcion string `gorm:"descripcion"`
	Url         string `gorm:"url"`
	Usuario_id  string `gorm:"usuario_id"`
}
