package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//GetImagenProduct trae la imagen de un producto
func GetImagenProduct(id int) (products entidades.Producto, err error) {
	db, _ := db.GetConexion()

	row, err := db.Query("SELECT Image FROM products WHERE ID = ?", id)

	if err != nil {
		return products, err
	}

	for row.Next() {
		var image string
		err = row.Scan(&image)

		products = entidades.Producto{
			Image: image,
		}
	}

	return products, nil
}
