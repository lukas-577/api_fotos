package routes

import (
	"apiFotos/api/models"
	"apiFotos/api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigureFotoRouter(router *gin.Engine, db *gorm.DB) {
	fotoGroup := router.Group("/foto")

	fotoGroup.GET("/get", func(ctx *gin.Context) {

		fotos, err := services.GetFotos(db)
		if err != nil {
			// Manejar el error de alguna manera, por ejemplo, devolver un error JSON
			ctx.JSON(500, gin.H{"error": "Error al obtener fotos"})
			return
		}

		ctx.JSON(200, fotos)
	})

	fotoGroup.POST("/create", func(ctx *gin.Context) {

		var foto models.Foto

		err := ctx.BindJSON(&foto)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Error al crear la foto"})
			return
		}

		err = services.CreateFoto(db, &foto)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error al crear la foto"})
			return
		}

		ctx.JSON(200, foto)

	})

	fotoGroup.DELETE("/delete/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		var foto models.Foto

		err := db.First(&foto, id).Error

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error al obtener la foto"})
			return
		}

		err = services.DeleteFoto(db, &foto)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error al eliminar la foto"})
			return
		}

		ctx.JSON(200, foto)

	})

	fotoGroup.PUT("/update/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		var foto models.Foto

		err := db.First(&foto, id).Error

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error al obtener la foto"})
			return
		}

		err = ctx.BindJSON(&foto)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Error al actualizar la foto"})
			return
		}

		err = services.UpdateFoto(db, &foto)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error al actualizar la foto"})
			return
		}

		ctx.JSON(200, foto)

	})
}
