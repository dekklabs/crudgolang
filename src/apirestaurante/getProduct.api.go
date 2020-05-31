package apirestaurante

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/model"
)

//GetProductsApi api que deuevle todos los productos
func GetProductsApi(w http.ResponseWriter, r *http.Request) {
	productos, err := model.GetProducts()

	if err != nil {
		http.Error(w, "Error al traer los productos", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productos)
}
