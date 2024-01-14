package models

type Factura struct {
	ID          int    `gorm:"id"`
	Fecha_pago  string `gorm:"fecha_pago"`
	Total       string `gorm:"total"`
	Estado      string `gorm:"estado"`
	Descripcion string `gorm:"descripcion"`
	Usuario_id  string `gorm:"usuario_id"`
}
