package apirestaurante

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model"
	"github.com/dekklabs/restaurantes/src/tokenjwt"
)

//LoginApi api para inicio de sessión de un usuario
func LoginApi(w http.ResponseWriter, r *http.Request) {
	var usuario entidades.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos", http.StatusBadRequest)
		return
	}

	user, existe := model.Login(usuario.Username, usuario.Pass)
	if existe == false {
		http.Error(w, "Usurio no existe", http.StatusBadRequest)
		return
	}

	jwtKey, err := tokenjwt.GeneroJWT(user)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generr el token correspondiente", http.StatusBadRequest)
		return
	}

	resp := entidades.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
