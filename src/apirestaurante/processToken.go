package apirestaurante

import (
	"errors"
	"strings"

	"github.com/dekklabs/restaurantes/src/entidades"
	"github.com/dekklabs/restaurantes/src/tools"
	"github.com/dgrijalva/jwt-go"
)

//UserName el valor del username en todos los endpoints
var UserName string

//IDusuario es el id devuelto del modelo, que se usará en todos los endpoints
var IDusuario int64

//ProcessToken procesa el token
func ProcessToken(tk string) (*entidades.Claim, bool, int64, error) {
	miClave := []byte("restaurantesDesarrolloDekk")
	claims := &entidades.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, 0, errors.New("Formato de token no válido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := tools.VerifyUserName(claims.Username)
		if encontrado == true {
			UserName = claims.Username
			IDusuario = claims.ID
		}
		return claims, encontrado, IDusuario, nil
	}

	if !tkn.Valid {
		return claims, false, 0, errors.New("Token inválido")
	}

	return claims, false, 0, err
}
