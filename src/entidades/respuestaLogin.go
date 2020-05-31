package entidades

//RespuestaLogin tiene el token que devuelve el login
type RespuestaLogin struct {
	Token string `json:"token"`
}