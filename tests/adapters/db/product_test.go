package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/simpletextbr/fullcycle-ports-and-adapters/adapters/db"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		price float,
		status string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products(id, name, price, status) VALUES (?, ?, ?, ?)`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec("abc", "Product Test", 10.30, "disabled")
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 10.30, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 10.30

	_, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, product.GetName())
	require.Equal(t, product.Price, product.GetPrice())
	require.Equal(t, product.Status, product.GetStatus())

	product.Status = "enabled"

	_, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Status, product.GetStatus())
}
