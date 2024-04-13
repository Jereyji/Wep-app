package sqlstore

import (
	"database/sql"

	"github.com/Jereyji/Wep-app.git/internal/app/model"
	"github.com/Jereyji/Wep-app.git/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, name, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Email, u.Name, u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, name, encrypted_password FROM users WHERE id = $1", 
		id,
		).Scan(
			&u.ID, 
			&u.Email, 
			&u.Name, 
			&u.EncryptedPassword,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, name, encrypted_password FROM users WHERE email = $1", 
		email,
		).Scan(
			&u.ID, 
			&u.Email, 
			&u.Name, 
			&u.EncryptedPassword,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}

	return u, nil
}