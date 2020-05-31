package products

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//UpdateProduct actualiza un producto
func UpdateProduct(producto *entidades.Producto, id int64) (int64, error) {
	db, _ := db.GetConexion()

	row, err := db.Exec(`UPDATE products SET 
		Name=?, 
		Price=?, 
		Discount_price=?, 
		Description=?, 
		Capacity=?, 
		Deliverable=?, 
		Category_id=?, 
		Market_id=?
		WHERE ID = ?`,
		producto.Name,
		producto.Price,
		producto.Discount_price,
		producto.Description,
		producto.Capacity,
		producto.Deliverable,
		producto.Category_id,
		producto.Market_id,
		producto.ID)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
