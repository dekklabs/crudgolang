package productsapi

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model/products"
)

//DeleteProductApi api para borrar un producto
func DeleteProductApi(w http.ResponseWriter, r *http.Request) {
	var producto entidades.Producto

	err := json.NewDecoder(r.Body).Decode(&producto)

	if err != nil {
		http.Error(w, "Error con el envio de datos", http.StatusBadRequest)
		return
	}

	_, err = products.DeleteProduct(producto.ID)

	if err != nil {
		http.Error(w, "Error al eliminar el producto", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
