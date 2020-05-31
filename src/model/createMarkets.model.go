package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//CreateMarketsModel modelo que crea una tienda
func CreateMarketsModel(markets *entidades.Markets) (err error) {
	db, err := db.GetConexion()

	if err != nil {
		return err
	}

	results, err2 := db.Exec("INSERT INTO markets(Name, Description, Address, Latitude, Longitude, Phone, Mobile, Information, Closed, Image)  VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", markets.Name, markets.Description, markets.Address, markets.Latitude, markets.Longitude, markets.Phone, markets.Mobile, markets.Information, markets.Closed, markets.Image)

	if err2 != nil {
		return err2
	}

	markets.ID, _ = results.LastInsertId()
	return nil
}
