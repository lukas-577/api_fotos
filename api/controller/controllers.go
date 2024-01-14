package controller

import (
	"fmt"
	"net/http"
)

func ControllerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
