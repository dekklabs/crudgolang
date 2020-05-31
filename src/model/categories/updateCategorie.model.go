package categories

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//UpdateCategories actiañoza una categoria
func UpdateCategories(categoria *entidades.Categories) (int64, error) {
	db, _ := db.GetConexion()

	row, err := db.Exec("UPDATE categories SET Name = ?, Description = ? WHERE ID = ?;", categoria.Name, categoria.Description, categoria.ID)

	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
