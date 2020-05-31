package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//GetProducts obtiene la lista de productos
func GetProducts() (products []entidades.Producto, err error) {
	db, _ := db.GetConexion()

	rows, err := db.Query("SELECT * FROM products;")

	defer db.Close()

	for rows.Next() {
		var id int64
		var name string
		var price float64
		var discount_price float64
		var description string
		var capacity float64
		var deliverable int64
		var image string
		var category_id int64
		var market_id int64
		var created_at string
		var updated_at string

		err = rows.Scan(
			&id,
			&name,
			&price,
			&discount_price,
			&description,
			&capacity,
			&deliverable,
			&image,
			&category_id,
			&market_id,
			&created_at,
			&updated_at)
		if err != nil {
			return nil, err
		}
		product := entidades.Producto{
			ID:             id,
			Name:           name,
			Price:          price,
			Discount_price: discount_price,
			Description:    description,
			Capacity:       capacity,
			Deliverable:    deliverable,
			Image:          image,
			Category_id:    category_id,
			Market_id:      market_id,
			Created_at:     created_at,
			Updated_at:     updated_at,
		}

		products = append(products, product)
	}

	return products, nil
}
