package apirestaurante

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model"
)

//UpdateUserApi api para actualizar un usuario
func UpdateUserApi(w http.ResponseWriter, r *http.Request) {
	var usuario entidades.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Error con los datos enviados", http.StatusBadRequest)
		return
	}

	_, err = model.UpdateProfile(&usuario)
	if err != nil {
		http.Error(w, "Error al actualizar los datos", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
