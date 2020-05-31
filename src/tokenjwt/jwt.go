package tokenjwt

import (
	"time"

	"github.com/dekklabs/restaurantes/src/entidades"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera un token para el login
func GeneroJWT(usuario entidades.Usuario) (string, error) {
	miClave := []byte("restaurantesDesarrolloDekk")

	payload := jwt.MapClaims{
		"ID":         usuario.ID,
		"Nombre":     usuario.Nombre,
		"Apellidos":  usuario.Apellidos,
		"Username":   usuario.Username,
		"Email":      usuario.Email,
		"Imagen":     usuario.Imagen,
		"Created_at": usuario.Created_at,
		"Updated_at": usuario.Updated_at,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
