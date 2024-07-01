package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reviewhub/database"
	"reviewhub/model"

	"github.com/gorilla/mux"
)

// Quando roda essa requisição aparece uma mensagem
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListarSeries(w http.ResponseWriter, r *http.Request) {
	var series []model.Serie
	database.DB.Find(&series)
	json.NewEncoder(w).Encode(series)
}

func ListarPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var serie model.Serie
	database.DB.First(&serie, id)
	json.NewEncoder(w).Encode(serie)
}

func InserirSerie(w http.ResponseWriter, r *http.Request) {
	var serie model.Serie
	json.NewDecoder(r.Body).Decode(&serie)
	database.DB.Create(&serie)
	json.NewEncoder(w).Encode(serie)
}

func DeletarSerie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var serie model.Serie
	database.DB.Delete(&serie, id)
	json.NewEncoder(w).Encode(serie)
}

func EditarSerie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var serie model.Serie
	database.DB.First(&serie, id)
	json.NewDecoder(r.Body).Decode(&serie)
	database.DB.Save(&serie)
	json.NewEncoder(w).Encode(serie)
}
