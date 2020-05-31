package apirestaurante

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/dekklabs/restaurantes/src/model"
)

//GetAvatarProfileApi api para traer el avatar de un usuario
func GetAvatarProfileApi(w http.ResponseWriter, r *http.Request) {
	idTemp := r.URL.Query().Get("id")

	if len(idTemp) < 1 {
		http.Error(w, "Es necesario un id", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idTemp)

	user, err := model.UserProfile(id)
	if err != nil {
		http.Error(w, "Error al traer la imagen", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + user.Imagen)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}
}
