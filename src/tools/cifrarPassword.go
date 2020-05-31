package tools

import "golang.org/x/crypto/bcrypt"

//CifrarPassword cifra la contraseña del registro
func CifrarPassword(password string) (string, error) {
	costo := 8
	hash, err := bcrypt.GenerateFromPassword([]byte(password), costo)

	return string(hash), err
}
