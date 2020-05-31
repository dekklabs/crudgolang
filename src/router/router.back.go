package router

import (
	apires "github.com/dekklabs/restaurantes/src/apirestaurante/categoriesapi"
	apimart "github.com/dekklabs/restaurantes/src/apirestaurante/marketsapi"
	apiprodu "github.com/dekklabs/restaurantes/src/apirestaurante/productsapi"
	"github.com/dekklabs/restaurantes/src/middlew"
	"github.com/gorilla/mux"
)

//RouterBack rutas para la administración
func RouterBack(router *mux.Router) {
	//Categorias
	router.HandleFunc("/api/v1/register-categories", middlew.VerifyDB(middlew.VarifyJwt(apires.RegisterCategoriesApi))).Methods("POST")
	router.HandleFunc("/api/v1/update-categorias", middlew.VerifyDB(middlew.VarifyJwt(apires.UpdateCategoriesApi))).Methods("PUT")
	router.HandleFunc("/api/v1/delete-categorias", middlew.VerifyDB(middlew.VarifyJwt(apires.DeleteCategoriesApi))).Methods("DELETE")

	//Markets
	router.HandleFunc("/api/v1/register-markets", middlew.VerifyDB(apimart.RegisterMerketsApi)).Methods("POST")

	//Products
	router.HandleFunc("/api/v1/register-product", middlew.VerifyDB(middlew.VarifyJwt(apiprodu.CreateProductApi))).Methods("POST")
	router.HandleFunc("/api/v1/update-product", middlew.VerifyDB(middlew.VarifyJwt(apiprodu.UpdateProductApi))).Methods("PUT")
	router.HandleFunc("/api/v1/delete-product", middlew.VerifyDB(middlew.VarifyJwt(apiprodu.DeleteProductApi))).Methods("DELETE")
}
