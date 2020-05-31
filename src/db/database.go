package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetConexion devuelve la conexion al a base de datos
func GetConexion() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "resturanteapi"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}

//CheckConexionDb revisa si la conexion a la base de datos es estable
func CheckConexionDb() int {
	_, err := GetConexion()
	if err != nil {
		return 0
	}
	return 1
}
