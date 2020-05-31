package model

import (
	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/tools"
)

//RegistroUsuario registro de usuarios
func RegistroUsuario(usuario *entidades.Usuario) (err error) {
	db, _ := db.GetConexion()
	usuario.Pass, _ = tools.CifrarPassword(usuario.Pass)
	usuario.Imagen = "none"
	results, err := db.Exec("INSERT INTO usuario(Nombre, Apellidos, Username, Pass, Email, Imagen) VALUES(?, ?, ?, ?, ?, ?)",
		usuario.Nombre, usuario.Apellidos, usuario.Username, usuario.Pass, usuario.Email, usuario.Imagen)
	if err != nil {
		return err
	} else {
		usuario.ID, _ = results.LastInsertId()
		return nil
	}
}
