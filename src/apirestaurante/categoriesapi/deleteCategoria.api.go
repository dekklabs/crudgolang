package categoriesapi

import (
	"net/http"
	"strconv"

	"github.com/dekklabs/restaurantes/src/model/categories"
)

//DeleteCategoriesApi api para borrar una categoría
func DeleteCategoriesApi(w http.ResponseWriter, r *http.Request) {
	idTemp := r.URL.Query().Get("id")

	if len(idTemp) < 1 {
		http.Error(w, "Es necesario el id", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idTemp)

	_, err := categories.DeleteCategoria(id)
	if err != nil {
		http.Error(w, "Error al eliminar la categoría", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
