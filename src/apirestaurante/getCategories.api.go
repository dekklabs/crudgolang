package apirestaurante

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/model"
)

//GetCategoriasApi Api para la obtención de las categorías
func GetCategoriasApi(w http.ResponseWriter, r *http.Request) {
	categorias, err := model.GetCategoriesModel()
	if err != nil {
		http.Error(w, "Error con la consulta", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categorias)
}
