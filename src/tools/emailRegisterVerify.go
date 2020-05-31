package tools

import (
	"github.com/dekklabs/restaurantes/src/db"
)

//VerifyEmail función que revisa si un correo ya esta registrado en la base de datos
func VerifyEmail(email string) (encontrado bool, err error) {
	db, _ := db.GetConexion()

	row := db.QueryRow("SELECT * FROM usuario WHERE Email = ? LIMIT 1", email)

	var id string
	var nombre string
	var apellidos string
	var username string
	var pass string
	// var email string
	var created_at string
	var updated_at string

	err = row.Scan(&id, &nombre, &apellidos, &username, &pass, &email, &created_at, &updated_at)
	if err != nil {
		return true, err
	} else {
		return false, err
	}
}
