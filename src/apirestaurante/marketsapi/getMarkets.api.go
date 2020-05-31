package marketsapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dekklabs/restaurantes/src/model/markets"
)

//GetMarkets api que devuelve todos los markets
func GetMarkets(w http.ResponseWriter, r *http.Request) {
	markets, err := markets.GetMarket()

	count := len(markets)

	fmt.Println("Total registros:", count)

	if err != nil {
		http.Error(w, "Error al traer los datos", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(markets)
}
