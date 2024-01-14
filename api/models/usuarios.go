package models

type Usuario struct {
	ID          int    `gorm:"id"`
	Nombre      string `gorm:"nombre"`
	Apellido    string `gorm:"apellido"`
	Email       string `gorm:"email"`
	Password    string `gorm:"password"`
	Img_prefile string `gorm:"img_prefile"`
}
