package authentication

import "database/sql"

type AuthStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) Register(*UserPayload) error {
	return nil
}

func (s *AuthStorage) Verify() {
}
