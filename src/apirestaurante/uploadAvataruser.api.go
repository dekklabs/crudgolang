package apirestaurante

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model"
)

//UploadAvatarUserApi api para subir avatar del usuario
func UploadAvatarUserApi(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]

	var idUser = strconv.Itoa(int(IDusuario))
	var archivo string = "uploads/avatars/" + idUser + "." + extension

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
	var usuario entidades.Usuario
	var status int64
	usuario.Imagen = idUser + "." + extension
	status, err = model.UploadAvatarUser(usuario, idUser)

	if err != nil || status == 0 {
		http.Error(w, "Error al grabar el avatar en al DB"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
