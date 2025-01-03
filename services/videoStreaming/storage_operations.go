package videoStreaming

import (
	"database/sql"
	"errors"
)

type StreamingStorage struct {
	db *sql.DB
}

func NewStreamingStorage(db *sql.DB) *StreamingStorage {
	return &StreamingStorage{
		db: db,
	}
}

func (s *StreamingStorage) GetVideoFilePath(videoID int) (string, error) {
	query := `SELECT file_path FROM videos WHERE id = $1`
	row := s.db.QueryRow(query, videoID)

	var filePath string
	err := row.Scan(&filePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("video not found")
		}
		return "", err
	}

	return filePath, nil
}
