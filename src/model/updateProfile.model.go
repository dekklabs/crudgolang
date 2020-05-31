package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//UpdateProfile actualiza un usuario
func UpdateProfile(usuario *entidades.Usuario) (int64, error) {
	db, _ := db.GetConexion()

	row, err := db.Exec("UPDATE usuario SET Nombre = ?, Apellidos = ? WHERE ID = ?", usuario.Nombre, usuario.Apellidos, usuario.ID)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
