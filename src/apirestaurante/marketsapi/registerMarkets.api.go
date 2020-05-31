package marketsapi

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/model"
)

//RegisterMerketsApi api para registrar markets
func RegisterMerketsApi(w http.ResponseWriter, r *http.Request) {
	var markets entidades.Markets

	err := json.NewDecoder(r.Body).Decode(&markets)

	if err != nil {
		http.Error(w, "Error al enviar los datos", http.StatusBadRequest)
		return
	}

	err = model.CreateMarketsModel(&markets)
	if err != nil {
		http.Error(w, "Error al guardar los datos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
