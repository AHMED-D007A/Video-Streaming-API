package videoUploading

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type UploadStorage struct {
	db *sql.DB
}

func NewUploadStorage(db *sql.DB) *UploadStorage {
	return &UploadStorage{
		db: db,
	}
}

func (s *UploadStorage) StoreVideo(video *Video) (int, error) {
	query := `INSERT INTO videos (user_id, title, description, file_path, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var videoID int
	err := s.db.QueryRow(query, video.UserID, video.Title, video.Description, video.FilePath, video.CreatedAt, video.UpdatedAt).Scan(&videoID)
	if err != nil {
		return 0, err
	}
	return videoID, nil
}

func (s *UploadStorage) UpdateVideoFilePath(videoID int, filePath string, thumbnailPath string) error {
	query := `UPDATE videos SET file_path = $1, thumbnail_path = $2 WHERE id = $3`
	_, err := s.db.Exec(query, filePath, thumbnailPath, videoID)
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

// TranscodeToMultipleResolutions transcodes a raw video file to multiple resolutions using FFmpeg
func TranscodeToMultipleResolutions(inputPath, outputDir string) error {
	resolutions := []string{"720", "480", "360"}
	for _, res := range resolutions {
		outputPath := fmt.Sprintf("%s/output_%sp.mp4", outputDir, res)
		cmd := exec.Command("ffmpeg", "-i", inputPath, "-vf", fmt.Sprintf("scale=-2:%s", res), "-c:a", "copy", outputPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("FFmpeg output: %s", string(output))
			return fmt.Errorf("failed to transcode video to %sp: %w", res, err)
		}
	}
	return nil
}

// ChunkVideoToHLS chunks a video file into HLS segments using FFmpeg
func ChunkVideoToHLS(outputDir string) error {
	resolutions := []string{"720", "480", "360"}
	for _, res := range resolutions {
		inputPath := fmt.Sprintf("%s/output_%sp.mp4", outputDir, res)
		outputPath := fmt.Sprintf("%s/output_%sp.m3u8", outputDir, res)
		cmd := exec.Command("ffmpeg", "-i", inputPath, "-c", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", outputPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("FFmpeg output: %s", string(output))
			return fmt.Errorf("failed to chunk video to HLS: %w", err)
		}
	}
	return nil
}

// CreateMasterPlaylist creates a master playlist for adaptive bitrate streaming
func CreateMasterPlaylist(outputDir string) error {
	resolutions := []string{"720", "480", "360"}
	masterPlaylistPath := fmt.Sprintf("%s/master.m3u8", outputDir)
	file, err := os.Create(masterPlaylistPath)
	if err != nil {
		return fmt.Errorf("failed to create master playlist: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("#EXTM3U\n")
	if err != nil {
		return fmt.Errorf("failed to write to master playlist: %w", err)
	}

	for _, res := range resolutions {
		_, err = file.WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%s000,RESOLUTION=%sx%s\n", res, res, res))
		if err != nil {
			return fmt.Errorf("failed to write to master playlist: %w", err)
		}
		_, err = file.WriteString(fmt.Sprintf("output_%sp.m3u8\n", res))
		if err != nil {
			return fmt.Errorf("failed to write to master playlist: %w", err)
		}
	}

	return nil
}
