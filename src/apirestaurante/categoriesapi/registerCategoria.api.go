package categoriesapi

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model"
)

//RegisterCategories api para registrar categorias
func RegisterCategoriesApi(w http.ResponseWriter, r *http.Request) {
	var categorias entidades.Categories
	err := json.NewDecoder(r.Body).Decode(&categorias)
	if err != nil {
		http.Error(w, "Error con los datos enviados", http.StatusBadRequest)
		return
	}

	err2 := model.CreateCategories(&categorias)
	if err2 != nil {
		http.Error(w, "al crear el usuario", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
