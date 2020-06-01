package productsapi

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model/products"
)

//CreateProductApi api para crear un producto
func CreateProductApi(w http.ResponseWriter, r *http.Request) {

	// idUsuario := apirestaurante.IDusuario
	// var id = strconv.Itoa(int(idUsuario))

	file, handler, err := r.FormFile("image")
	var extension = strings.Split(handler.Filename, ".")[1]

	nombre := r.PostFormValue("name")
	splitnombre := strings.Replace(nombre, " ", "-", -1)

	var archivo string = "uploads/products/" + splitnombre + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}

	precio, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	desc, _ := strconv.ParseFloat(r.PostFormValue("discount_price"), 64)
	capacity, _ := strconv.ParseFloat(r.PostFormValue("capacity"), 64)
	delivery, _ := strconv.Atoi(r.PostFormValue("deliverable"))
	category_id, _ := strconv.Atoi(r.PostFormValue("category_id"))
	market_id, _ := strconv.Atoi(r.PostFormValue("market_id"))

	producto := entidades.Producto{
		Name:           r.PostFormValue("name"),
		Price:          precio,
		Discount_price: desc,
		Description:    r.PostFormValue("description"),
		Capacity:       capacity,
		Deliverable:    int64(delivery),
		Image:          splitnombre + "." + extension,
		Category_id:    int64(category_id),
		Market_id:      int64(market_id),
	}

	if err != nil {
		http.Error(w, "Error con el envio de datos", http.StatusBadRequest)
		return
	}

	err = products.CreateProduct(&producto)
	if err != nil {
		http.Error(w, "Error al registrar producto", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
