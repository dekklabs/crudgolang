package products

import (
	"github.com/dekklabs/restaurantes/src/db"
)

//DeleteProduct borrar producto con el id
func DeleteProduct(id int64) (int64, error) {
	db, _ := db.GetConexion()
	row, err := db.Exec("DELETE FROM products WHERE ID = ?", id)

	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
