package dao

import (
	"marketplace-soa/db"
	"marketplace-soa/model"
)

// POST on /clients
func InsertClient(client model.Client) (int, error) {
	query := `INSERT INTO clients (name, email, document) VALUES ($1, $2, $3) returning id;`
	err := db.GetDB().QueryRow(query, client.Name, client.Email, client.Document).Scan(&client.ID)
	if err != nil {
		return 0, err
	}
	return client.ID, nil
}

// PUT on /clients/:id
func UpdateClient(client model.Client) error {
	query := `UPDATE clients SET name = $1, email = $2, document = $3 WHERE id = $4;`
	_, err := db.GetDB().Exec(query, client.Name, client.Email, client.Document, client.ID)
	if err != nil {
		return err
	}
	return nil
}

// DELETE on /clients/:id
func DeleteClient(id int) error {
	query := `DELETE FROM clients WHERE id = $1;`
	_, err := db.GetDB().Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GET on /clients/:id
func GetClientByID(id int) (model.Client, error) {
	query := `SELECT id, name, email, document FROM clients WHERE id = $1;`
	row := db.GetDB().QueryRow(query, id)
	var client model.Client
	err := row.Scan(&client.ID, &client.Name, &client.Email, &client.Document)
	if err != nil {
		return model.Client{}, err
	}
	return client, nil
}

// GET on /clients
func GetAllClients() ([]model.Client, error) {
	query := `SELECT id, name, email, document FROM clients;`
	rows, err := db.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clients []model.Client
	for rows.Next() {
		var client model.Client
		err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Document)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
