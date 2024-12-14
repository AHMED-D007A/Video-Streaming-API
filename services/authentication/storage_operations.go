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

func (s *AuthStorage) Register(user *UserPayload) error {
	query := "INSERT INTO users(username, email, password_hash) VALUES($1, $2, $3)"
	_, err := s.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthStorage) Verify() {
}
