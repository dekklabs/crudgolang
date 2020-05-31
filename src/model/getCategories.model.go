package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//GetCategoriesModel deuvel la consulta de categorías a la base de datos
func GetCategoriesModel() (categoria []entidades.Categories, err error) {
	db, dbErr := db.GetConexion()

	if dbErr != nil {
		return nil, dbErr
	}

	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	var categorias []entidades.Categories
	for rows.Next() {
		var id int64
		var name string
		var description string
		var created_at string
		var updated_at string

		err2 := rows.Scan(&id, &name, &description, &created_at, &updated_at)
		if err2 != nil {
			return nil, err2
		} else {
			categoria := entidades.Categories{
				ID:          id,
				Name:        name,
				Description: description,
				Created_at:  created_at,
				Updated_at:  updated_at,
			}
			categorias = append(categorias, categoria)
		}
	}

	return categorias, nil
}
