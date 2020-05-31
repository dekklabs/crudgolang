package model

import (
	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/tools"
	"golang.org/x/crypto/bcrypt"
)

//Login inicio de sesión
func Login(username, password string) (entidades.Usuario, bool) {

	user, encontrado, _ := tools.VerifyUserName(username)

	usuario := entidades.Usuario{}
	if encontrado == false {
		return usuario, false
	}

	userPassword := []byte(user.Pass)
	passwordBody := []byte(password)

	user.Pass = ""

	err := bcrypt.CompareHashAndPassword(userPassword, passwordBody)
	if err != nil {
		return usuario, false
	}

	return user, true
}
