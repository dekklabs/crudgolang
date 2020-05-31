package categories

import (
	"github.com/dekklabs/restaurantes/src/db"
)

//DeleteCategoria elimina una categoría
func DeleteCategoria(id int) (int64, error) {
	db, _ := db.GetConexion()

	row, err := db.Exec("DELETE FROM categories WHERE ID = ?", id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
