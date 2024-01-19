package handler

import (
	"apiFotos/api/controller"
	"net/http"
)

func Init() {
	http.HandleFunc("/", controller.ControllerTest)
	http.HandleFunc("/ingresaFoto", controller.IngresaFotoTest)
}
