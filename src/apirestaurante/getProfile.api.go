package apirestaurante

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dekklabs/restaurantes/src/model"
)

//GetProfileApi api que devuelve el usuario buscado por su identificador
func GetProfileApi(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "Debe ingresar un id", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)

	user, err := model.UserProfile(idInt)
	user.Pass = ""
	if err != nil {
		http.Error(w, "Usuario no existe", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
