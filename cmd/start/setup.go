package main

import (
	"fmt"
	"marketplace-soa/db"
	"marketplace-soa/model"
)

var (
	products = []model.Product{
		{Name: "Laptop", Description: "High performance laptop", Price: 1200.00, Category: "Electronics", Active: true},
		{Name: "Smartphone", Description: "Latest model smartphone", Price: 800.00, Category: "Electronics", Active: true},
		{Name: "Headphones", Description: "Noise-cancelling headphones", Price: 150.00, Category: "Accessories", Active: true},
		{Name: "Coffee Maker", Description: "Automatic coffee maker", Price: 100.00, Category: "Home Appliances", Active: true},
		{Name: "Gaming Console", Description: "Next-gen gaming console", Price: 500.00, Category: "Entertainment", Active: true},
		{Name: "Smartwatch", Description: "Feature-rich smartwatch", Price: 250.00, Category: "Wearables", Active: true},
		{Name: "Tablet", Description: "Lightweight tablet", Price: 300.00, Category: "Electronics", Active: true},
		{Name: "Camera", Description: "Digital SLR camera", Price: 700.00, Category: "Photography", Active: true},
		{Name: "Bluetooth Speaker", Description: "Portable Bluetooth speaker", Price: 80.00, Category: "Audio", Active: true},
		{Name: "E-reader", Description: "Compact e-reader device", Price: 120.00, Category: "Books", Active: true},
	}

	clients = []model.Client{
		{Name: "Alice Johnson", Email: "alice@example.com", Document: "12345678901"},
		{Name: "Bob Smith", Email: "bob@example.com", Document: "98765432100"},
		{Name: "Charlie Brown", Email: "charlie@example.com", Document: "45678912300"},
		{Name: "Diana Prince", Email: "diana@example.com", Document: "32165498700"},
		{Name: "Ethan Hunt", Email: "ethan@example.com", Document: "65432178900"},
		{Name: "Fiona Glenanne", Email: "fiona@example.com", Document: "78901234500"},
		{Name: "George Clooney", Email: "george@example.com", Document: "32145678900"},
		{Name: "Hannah Montana", Email: "hannah@example.com", Document: "98765432102"},
		{Name: "Ian Fleming", Email: "ian@example.com", Document: "45612378900"},
		{Name: "Jane Doe", Email: "jane@example.com", Document: "12312312300"},
	}
)

func main() {
	fmt.Println("Starting setup...")
	createTables()
	seedTables()
	fmt.Println("Setup completed.")
}

func createTables() {
	createClientsTable()
	createProductsTable()
}

func seedTables() {
	validateClientsTable()
	validateProductsTable()
}

func validateClientsTable() {
	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM clients").Scan(&count)
	if err != nil {
		fmt.Println("Error validating clients table:", err)
		return
	}
	if count == 0 {
		seedClients()
		fmt.Println("Clients table was empty, seeded data.")
		return
	}
	fmt.Println("Clients table already has data, skipping seeding.")
}

func validateProductsTable() {
	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		fmt.Println("Error validating products table:", err)
		return
	}
	if count == 0 {
		seedProducts()
		fmt.Println("Products table was empty, seeded data.")
		return
	}
	fmt.Println("Products table already has data, skipping seeding.")
}

func createClientsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS clients (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		document VARCHAR(14) NOT NULL
	);`
	_, err := db.GetDB().Exec(query)
	if err != nil {
		fmt.Println("Error creating clients table:", err)
	}
}

func createProductsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		price NUMERIC(10, 2) NOT NULL,
		category VARCHAR(50),
		active BOOLEAN DEFAULT TRUE
	);`
	_, err := db.GetDB().Exec(query)
	if err != nil {
		fmt.Println("Error creating products table:", err)
	}
}

func seedClients() {
	for _, client := range clients {
		query := `INSERT INTO clients (name, email, document) VALUES ($1, $2, $3);`
		_, err := db.GetDB().Exec(query, client.Name, client.Email, client.Document)
		if err != nil {
			fmt.Println("Error seeding client:", err)
		}
	}
	fmt.Println("Clients seeded successfully.")
}

func seedProducts() {
	for _, product := range products {
		query := `INSERT INTO products (name, description, price, category, active) VALUES ($1, $2, $3, $4, $5);`
		_, err := db.GetDB().Exec(query, product.Name, product.Description, product.Price, product.Category, product.Active)
		if err != nil {
			fmt.Println("Error seeding product:", err)
		}
	}
	fmt.Println("Products seeded successfully.")
}
