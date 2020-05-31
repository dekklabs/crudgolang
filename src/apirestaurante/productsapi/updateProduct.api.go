package productsapi

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/apirestaurante"
	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model/products"
)

//UpdateProductApi api para actualizar un producto
func UpdateProductApi(w http.ResponseWriter, r *http.Request) {
	var producto entidades.Producto

	id := apirestaurante.IDusuario

	err := json.NewDecoder(r.Body).Decode(&producto)

	if err != nil {
		http.Error(w, "Error con los datos que ingresó", http.StatusBadRequest)
		return
	}

	_, err = products.UpdateProduct(&producto, id)

	if err != nil {
		http.Error(w, "Error al actualizar el producto", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
