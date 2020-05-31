package tools

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//VerifyExitsuser verifica si un usuario existe en la base de datos
func VerifyExitsuser(email, username string) (user entidades.Usuario, encontrado bool, err error) {
	db, _ := db.GetConexion()

	// rows, err := db.Query("SELECT * FROM usuario WHERE Email = ? LIMIT 1", email)
	rows, err := db.Query("SELECT * FROM usuario WHERE Email = ? OR Username = ? LIMIT 1", email, username)

	if err != nil {
		encontrado = false
		return user, encontrado, err
	}

	if rows.Next() {
		var id int64
		var nombre string
		var apellidos string
		var username string
		var pass string
		var email2 string
		var imagen string
		var created_at string
		var updated_at string

		err = rows.Scan(&id, &nombre, &apellidos, &username, &pass, &email2, &imagen, &created_at, &updated_at)

		if err != nil {
			encontrado = false
			return user, encontrado, err
		}

		user = entidades.Usuario{
			ID:         id,
			Nombre:     nombre,
			Apellidos:  apellidos,
			Username:   username,
			Pass:       pass,
			Email:      email2,
			Imagen:     imagen,
			Created_at: created_at,
			Updated_at: updated_at,
		}
	} else {
		encontrado = false
		return user, encontrado, nil
	}
	encontrado = true
	return user, encontrado, nil
}

//VerifyUserName verifica si un username existe en la base de datos
func VerifyUserName(username string) (user entidades.Usuario, encontrado bool, err error) {
	db, _ := db.GetConexion()

	// rows, err := db.Query("SELECT * FROM usuario WHERE Email = ? LIMIT 1", email)
	rows, err := db.Query("SELECT * FROM usuario WHERE Username = ? LIMIT 1", username)

	if err != nil {
		encontrado = false
		return user, encontrado, err
	}

	if rows.Next() {
		var id int64
		var nombre string
		var apellidos string
		var username string
		var pass string
		var email2 string
		var imagen string
		var created_at string
		var updated_at string

		err = rows.Scan(&id, &nombre, &apellidos, &username, &pass, &email2, &imagen, &created_at, &updated_at)

		if err != nil {
			encontrado = false
			return user, encontrado, err
		}

		user = entidades.Usuario{
			ID:         id,
			Nombre:     nombre,
			Apellidos:  apellidos,
			Username:   username,
			Pass:       pass,
			Email:      email2,
			Imagen:     imagen,
			Created_at: created_at,
			Updated_at: updated_at,
		}

	} else {
		encontrado = false
		return user, encontrado, nil
	}
	encontrado = true
	return user, encontrado, nil
}
