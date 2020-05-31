package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dekklabs/restaurantes/src/draw"
	"github.com/dekklabs/restaurantes/src/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Running inicia el servidor de restaurantes
func StartServer() {
	r := mux.NewRouter()

	//Rutas
	//Frontend
	router.Routers(r)
	//Backend
	router.RouterBack(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	handler := cors.Default().Handler(r)

	draw.Drawing()

	//Server encendido
	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		fmt.Println("Error :,c ", err)
	}
}
