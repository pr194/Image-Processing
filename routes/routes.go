package routes

import (
	"net/http"

	"github.com/pr194/Collaborative-tool/controllers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/upload", controllers.UploadFile)
	mux.HandleFunc("/process", controllers.Processfile)

}
