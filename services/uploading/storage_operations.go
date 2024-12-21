package uploading

import (
	"database/sql"
	"errors"
)

type UploadStorage struct {
	db *sql.DB
}

func NewUploadStorage(db *sql.DB) *UploadStorage {
	return &UploadStorage{
		db: db,
	}
}

func (s *UploadStorage) StoreVideo(video *Video) error {
	query := `INSERT INTO videos (user_id, title, description, file_path, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(query, video.UserID, video.Title, video.Description, video.FilePath, video.CreatedAt, video.UpdatedAt)
	return err
}

func (s *UploadStorage) GetUserIDByEmail(email string) (int, error) {
	query := `SELECT id FROM users WHERE email = $1`
	row := s.db.QueryRow(query, email)

	var userID int
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return userID, nil
}
