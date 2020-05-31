package products

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//CreateProduct crea un nuevo producto
// func CreateProduct(producto *entidades.Producto) (err error) {
func CreateProduct(producto *entidades.Producto) (err error) {
	db, _ := db.GetConexion()

	user, err := db.Exec("INSERT INTO products(Name, Price, Discount_price, Description, Capacity, Deliverable, Image, Category_id, Market_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);",
		producto.Name,
		producto.Price,
		producto.Discount_price,
		producto.Description,
		producto.Capacity,
		producto.Deliverable,
		producto.Image,
		producto.Category_id,
		producto.Market_id)

	if err != nil {
		return err
	}

	producto.ID, _ = user.LastInsertId()
	return nil
}
