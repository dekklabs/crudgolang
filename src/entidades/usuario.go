package entidades

//Usuario entidad de usuario
type Usuario struct {
	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Apellidos  string `json:"apellidos"`
	Username   string `json:"username"`
	Pass       string `json:"pass,omitempty"`
	Email      string `json:"email"`
	Imagen     string `json:"imagen,omitempty"`
	Created_at string `json:"created_at"`
	Updated_at string `sjon:"updated_at"`
}
