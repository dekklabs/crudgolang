package model

import (
	"fmt"

	"github.com/dekklabs/restaurantes/src/db"
	"github.com/dekklabs/restaurantes/src/entidades"
)

//CreateCategories Crea registro de categorias
func CreateCategories(categoria *entidades.Categories) (err error) {
	db, _ := db.GetConexion()

	fmt.Println(db)

	results, err := db.Exec("INSERT INTO categories(Name, Description) VALUES(?, ?)", categoria.Name, categoria.Description)
	if err != nil {
		return err
	} else {
		categoria.ID, _ = results.LastInsertId()
		return nil
	}
}
