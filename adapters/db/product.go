package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.IProduct, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return &product, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return &product, err
	}

	return &product, nil
}
