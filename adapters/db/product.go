package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
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
func (p *ProductDb) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	p.db.QueryRow("SELECT COUNT(id) FROM products WHERE id = ?", product.GetID()).Scan(&rows)
	if rows == 0 {
		if _, err := p.create(product); err != nil {
			return nil, err
		}
	} else {
		if _, err := p.update(product); err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("INSERT INTO products(id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}

	return product, nil
}
