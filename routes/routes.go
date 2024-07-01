package routes

import (
	"log"
	"net/http"
	"reviewhub/controller"
	"reviewhub/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Função para cuidar das requisições
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/series", controller.ListarSeries).Methods("Get")
	r.HandleFunc("/api/series/{id}", controller.ListarPorId).Methods("Get")
	r.HandleFunc("/api/series", controller.InserirSerie).Methods("Post")
	r.HandleFunc("/api/series/{id}", controller.DeletarSerie).Methods("Delete")
	r.HandleFunc("/api/series/{id}", controller.EditarSerie).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
