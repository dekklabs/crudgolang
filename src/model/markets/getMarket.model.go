package markets

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//GetMarket trae todos los markets
func GetMarket() (markets []entidades.Markets, err error) {
	db, _ := db.GetConexion()

	rows, err := db.Query("SELECT * FROM markets;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var name string
		var description string
		var address string
		var latitude string
		var longitude string
		var phone int64
		var mobile int64
		var information string
		var closed int64
		var image string
		var created_at string
		var updated_at string

		err2 := rows.Scan(
			&id,
			&name,
			&description,
			&address,
			&latitude,
			&longitude,
			&phone,
			&mobile,
			&information,
			&closed,
			&image,
			&created_at,
			&updated_at)

		if err2 != nil {
			return nil, err2
		} else {
			market := entidades.Markets{
				ID:          id,
				Name:        name,
				Description: description,
				Address:     address,
				Latitude:    latitude,
				Longitude:   longitude,
				Phone:       phone,
				Mobile:      mobile,
				Information: information,
				Closed:      closed,
				Image:       image,
				Created_at:  created_at,
				Updated_at:  updated_at,
			}
			markets = append(markets, market)
		}
	}

	return markets, nil
}
