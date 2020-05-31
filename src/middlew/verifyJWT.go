package middlew

import (
	"net/http"

	"github.com/dekklabs/restaurantes/src/apirestaurante"
)

//VarifyJWT permite validar el JWT que nos viene en la petición
func VarifyJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := apirestaurante.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
