package service

import (
	"marketplace-soa/dao"
	"marketplace-soa/model"
)

type ClientService interface {
	Create(client model.Client) (int, error)
	Update(client model.Client) error
	Delete(id int) error
	GetByID(id int) (model.Client, error)
	GetAll() ([]model.Client, error)
}

type clientService struct{}

func NewClientService() ClientService {
	return &clientService{}
}

func (s *clientService) Create(client model.Client) (int, error) {
	return dao.InsertClient(client)
}

func (s *clientService) Update(client model.Client) error {
	return dao.UpdateClient(client)
}

func (s *clientService) Delete(id int) error {
	return dao.DeleteClient(id)
}

func (s *clientService) GetByID(id int) (model.Client, error) {
	return dao.GetClientByID(id)
}

func (s *clientService) GetAll() ([]model.Client, error) {
	return dao.GetAllClients()
}
