package router

import (
	api "github.com/dekklabs/restaurantes/src/apirestaurante"
	apimart "github.com/dekklabs/restaurantes/src/apirestaurante/marketsapi"
	"github.com/dekklabs/restaurantes/src/middlew"
	"github.com/gorilla/mux"
)

func Routers(router *mux.Router) {
	// Usuarios
	router.HandleFunc("/api/v1/registro", middlew.VerifyDB(api.RegistroUsuarioApi)).Methods("POST")
	router.HandleFunc("/api/v1/update-user", middlew.VerifyDB(middlew.VarifyJwt(api.UpdateUserApi))).Methods("PUT")
	router.HandleFunc("/api/v1/upload-avatar", middlew.VerifyDB(middlew.VarifyJwt(api.UploadAvatarUserApi))).Methods("PUT")
	router.HandleFunc("/api/v1/get-avatar", middlew.VerifyDB(api.GetAvatarProfileApi)).Methods("GET")

	// Categorias
	router.HandleFunc("/api/v1/get-categorias", middlew.VerifyDB(api.GetCategoriasApi)).Methods("GET")

	// Markets
	router.HandleFunc("/api/v1/get-markets", middlew.VerifyDB(apimart.GetMarkets)).Methods("GET")

	// Products
	router.HandleFunc("/api/v1/get-products", middlew.VerifyDB(api.GetProductsApi)).Methods("GET")
	router.HandleFunc("/api/v1/get-imagen-producto", middlew.VerifyDB(api.GetImageProductApi)).Methods("GET")

	/* Login */
	router.HandleFunc("/api/v1/login", middlew.VerifyDB(api.LoginApi)).Methods("POST")

	router.HandleFunc("/api/v1/ver-perfil", middlew.VerifyDB(middlew.VarifyJwt(api.GetProfileApi))).Methods("GET")
}
