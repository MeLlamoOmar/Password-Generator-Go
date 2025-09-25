package service

import (
	"goPasswordGenerator/model"
	"goPasswordGenerator/store"
)

type service struct {
	store store.Store
}

func New(s store.Store) *service {
	return &service{store: s}
}

func (s *service) GetAllPasswords() ([]*model.Password, error) {
	return s.store.GetAll()
}

func (s *service) GetPasswordById(id int) (*model.Password, error) {
	return s.store.GetById(id)
}

func (s *service) CreatePassword(p *model.Password) (*model.Password, error) {
	return s.store.Create(p)
}

func (s *service) UpdatePassword(id int, password *model.Password) (*model.Password, error) {
	return s.store.Update(id, password)
}

func (s *service) DeletePassword(id int) error {
	return s.store.Delete(id)
}