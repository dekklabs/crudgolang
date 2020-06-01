package apirestaurante

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/dekklabs/restaurantes/src/model"
)

//getImageProductApi api que devuelve la imagen de un producto
func GetImageProductApi(w http.ResponseWriter, r *http.Request) {
	idTemp := r.URL.Query().Get("id")

	if len(idTemp) < 1 {
		http.Error(w, "Debe ingresar un id", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idTemp)

	product, err := model.GetImagenProduct(id)

	fmt.Println(product)

	if err != nil {
		http.Error(w, "Error al traer el producto", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/products/" + product.Image)

	if err != nil {
		http.Error(w, "Error al abrir el archivo", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar el archivo", http.StatusBadRequest)
		return
	}
}
