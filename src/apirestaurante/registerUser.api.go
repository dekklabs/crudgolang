package apirestaurante

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model"
	"github.com/dekklabs/restaurantes/src/tools"
)

//RegistroUsuarioApi registro api
func RegistroUsuarioApi(w http.ResponseWriter, r *http.Request) {
	var usuario entidades.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(usuario.Pass) < 6 {
		http.Error(w, "La cantidad de caracteres del password es menor a 6", 400)
		return
	}

	_, encontrado, _ := tools.VerifyExitsuser(usuario.Email, usuario.Username)
	if encontrado == true {
		http.Error(w, "Email o username ya existe en la base de datos", 400)
		return
	}

	err2 := model.RegistroUsuario(&usuario)
	if err2 != nil {
		http.Error(w, "error al registrar", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
