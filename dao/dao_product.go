package dao

import (
	"marketplace-soa/db"
	"marketplace-soa/model"
)

func InsertProduct(product model.Product) (int, error) {
	query := `INSERT INTO products (name, description, price, category, active) VALUES ($1, $2, $3, $4, $5) returning id;`
	err := db.GetDB().QueryRow(query, product.Name, product.Description, product.Price, product.Category, product.Active).Scan(&product.ID)
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func UpdateProduct(product model.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, category = $4, active = $5 WHERE id = $6;`
	_, err := db.GetDB().Exec(query, product.Name, product.Description, product.Price, product.Category, product.Active, product.ID)
	return err
}

func DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1;`
	_, err := db.GetDB().Exec(query, id)
	return err
}

func GetProductByID(id int) (model.Product, error) {
	var product model.Product
	query := `SELECT id, name, description, price, category, active FROM products WHERE id = $1;`
	err := db.GetDB().QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Active)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	query := `SELECT id, name, description, price, category, active FROM products;`
	rows, err := db.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Active); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
