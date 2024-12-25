package videoManagement

import (
	"database/sql"
	"errors"
)

type VideoStorage struct {
	db *sql.DB
}

func NewVideoStorage(db *sql.DB) *VideoStorage {
	return &VideoStorage{db: db}
}

func (s *VideoStorage) GetVideos(userID int) ([]Video, []string, error) {
	query := `SELECT id, title, description, thumbnail_path, created_at, updated_at FROM videos WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var videos []Video
	var thumbnail_paths []string
	for rows.Next() {
		var video Video
		var thumbnail_path string
		err := rows.Scan(&video.ID, &video.Title, &video.Description, &thumbnail_path, &video.CreatedAt, &video.UpdatedAt)
		if err != nil {
			return nil, nil, err
		}
		videos = append(videos, video)
		thumbnail_paths = append(thumbnail_paths, thumbnail_path)
	}

	return videos, thumbnail_paths, nil
}

func (s *VideoStorage) GetVideoThumbnail(videoID int) (string, error) {
	query := `SELECT thumbnail_path FROM videos WHERE id = $1`
	row := s.db.QueryRow(query, videoID)

	var thumbnailPath string
	err := row.Scan(&thumbnailPath)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("video not found")
		}
		return "", err
	}

	return thumbnailPath, nil
}

func (s *VideoStorage) UpdateVideo(video *VideoPayload) error {
	query := `UPDATE videos SET title = $1, description = $2, thumbnail_path = $3, updated_at = NOW() WHERE id = $4`
	_, err := s.db.Exec(query, video.Title, video.Description, video.Thumbnail_path, video.ID)
	return err
}

func (s *VideoStorage) DeleteVideo(id int) error {
	query := `DELETE FROM videos WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *VideoStorage) GetUserIDByEmail(email string) (int, error) {
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
