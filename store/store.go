package store

import (
	"database/sql"
	"goWeb/model"
)

// import "database/sql"

type Store interface {
	GetAll() ([]*model.Password, error)
	GetById(id int) (*model.Password, error)
	Create(p *model.Password) (*model.Password, error)
	Update(id int, password *model.Password) (*model.Password, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

// Create implements Store.
func (s *store) Create(p *model.Password) (*model.Password, error) {
	q := `INSERT INTO passwords (label, password, created_at) VALUES (?, ?, ?)`

	res, err := s.db.Exec(q, p.Label, p.Password, p.CreatedAt)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	p.ID = int(id)

	return p, nil
}

// Delete implements Store.
func (s *store) Delete(id int) error {
	q := `DELETE FROM passwords WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements Store.
func (s *store) GetAll() ([]*model.Password, error) {
	q := `SELECT id, label, password FROM passwords`

	rows, err := s.db.Query(q, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []*model.Password
	for rows.Next() {
		var p model.Password
		if err := rows.Scan(&p.ID, &p.Label, &p.Password); err != nil {
			return nil, err
		}

		passwords = append(passwords, &p)
	}

	return passwords, nil
}

// GetById implements Store.
func (s *store) GetById(id int) (*model.Password, error) {
	q := `SELECT id, label, password FROM passwords WHERE id = ?`

	var p model.Password

	err := s.db.QueryRow(q, id).Scan(&p.ID, &p.Label, &p.Password)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Update implements Store.
func (s *store) Update(id int, password *model.Password) (*model.Password, error) {
	q := `UPDATE passwords SET label = ?, password = ? WHERE id = ?`

	_, err := s.db.Exec(q, password.Label, password.Password, id)
	if err != nil {
		return nil, err
	}

	password.ID = id

	return password, nil

}

func New(db *sql.DB) Store {
	return &store{db: db}
}
