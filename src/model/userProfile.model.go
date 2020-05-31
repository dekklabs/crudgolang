package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//UserProfile muestra el perfil de un usuario
func UserProfile(id int) (user entidades.Usuario, err error) {
	db, err := db.GetConexion()

	if err != nil {
		return user, err
	}

	profile, err2 := db.Query("SELECT * FROM usuario WHERE ID = ? LIMIT 1;", id)

	if err2 != nil {
		return user, err2
	}

	defer profile.Close()

	if profile.Next() {
		var id2 int64
		var nombre string
		var apellidos string
		var username string
		var pass string
		var email string
		var imagen string
		var created_at string
		var updated_at string

		err3 := profile.Scan(&id2, &nombre, &apellidos, &username, &pass, &email, &imagen, &created_at, &updated_at)
		if err3 != nil {
			return user, err
		} else {
			user = entidades.Usuario{
				ID:         id2,
				Nombre:     nombre,
				Apellidos:  apellidos,
				Username:   username,
				Pass:       pass,
				Email:      email,
				Imagen:     imagen,
				Created_at: created_at,
				Updated_at: updated_at,
			}
		}
	} else {
		return user, err
	}

	return user, nil
}
