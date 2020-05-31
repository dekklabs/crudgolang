package middlew

import (
	"net/http"

	"github.com/dekklabs/restaurantes/src/db"
)

//VerifyDB middleware que verifica la conexion a la base de datos
func VerifyDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConexionDb() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
