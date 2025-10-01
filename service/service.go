package service

import (
	"fmt"
	"goPasswordGenerator/model"
	"goPasswordGenerator/store"
	"goPasswordGenerator/util"
	"os"
)

type service struct {
	store store.Store
}

func New(s store.Store) *service {
	return &service{store: s}
}

func (s *service) GetAllPasswords() ([]*model.Password, error) {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		return nil, fmt.Errorf("the ENCRYPT_KEY doesnt exist on .env")
	}

	passwords, err := s.store.GetAll()
	for _, password := range passwords {
		password.Password, err = util.Decrypt(password.Password, []byte(key))
		if err != nil {
			return nil, err
		}
	}
	return passwords, err
}

func (s *service) GetPasswordById(id int) (*model.Password, error) {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		return nil, fmt.Errorf("the ENCRYPT_KEY doesnt exist on .env")
	}

	p, err := s.store.GetById(id)
	if err != nil {
		return nil, err
	}

	p.Password, err = util.Decrypt(p.Password, []byte(key))
	if err != nil {
		return nil, err
	}
	
	return p, nil
}

func (s *service) CreatePassword(p *model.Password) (*model.Password, error) {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		return nil, fmt.Errorf("the ENCRYPT_KEY doesnt exist on .env")
	}
	encryptPass, err := util.Encrypt([]byte(p.Password), []byte(key))
	if err != nil {
		return nil, err
	}

	p.Password = encryptPass
	return s.store.Create(p)
}

func (s *service) UpdatePassword(id int, password *model.Password) (*model.Password, error) {
	return s.store.Update(id, password)
}

func (s *service) DeletePassword(id int) error {
	return s.store.Delete(id)
}