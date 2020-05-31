package categoriesapi

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model/categories"
)

//UpdateCategoriesApi api para actualizar las categorías
func UpdateCategoriesApi(w http.ResponseWriter, r *http.Request) {
	var categorias entidades.Categories

	err := json.NewDecoder(r.Body).Decode(&categorias)

	if err != nil {
		http.Error(w, "Error con los datos enviados", http.StatusBadRequest)
		return
	}

	_, err = categories.UpdateCategories(&categorias)
	if err != nil {
		http.Error(w, "Error con los datos enviados", http.StatusBadRequest)
		return
	}
}
