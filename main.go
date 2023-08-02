package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	Db "github.com/simpletextbr/fullcycle-ports-and-adapters/adapters/db"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite3")
	productDbAdapter := Db.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Test", 10.30)
	productService.Enable(product)
}
