package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//UploadAvatarUser sube una imagen a la base de datos
func UploadAvatarUser(usuario entidades.Usuario, ID string) (int64, error) {
	db, _ := db.GetConexion()

	row, err := db.Exec("UPDATE usuario SET Imagen = ? WHERE ID = ?", usuario.Imagen, ID)

	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
